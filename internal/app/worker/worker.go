package worker

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/cucumberjaye/gophermart/configs"
	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

const workers = 5

type AccrualRepository interface {
	GetWaitingOrders() ([]string, error)
	UpdateOrder(models.Order) error
}

type Worker struct {
	client     *http.Client
	wg         *sync.WaitGroup
	cancelFunc context.CancelFunc
	repo       AccrualRepository
	ch         chan string
}

func New(repo AccrualRepository) *Worker {
	return &Worker{
		wg:     new(sync.WaitGroup),
		repo:   repo,
		ch:     make(chan string),
		client: &http.Client{},
	}
}

func (w *Worker) Start(pctx context.Context) {
	log.Info().Msg("worker running")
	ctx, cancelFunc := context.WithCancel(pctx)
	w.cancelFunc = cancelFunc
	w.wg.Add(1)
	go w.ordersGetter(ctx)
	for i := 0; i < workers; i++ {
		w.wg.Add(1)
		go w.spawnWorkers(ctx)
	}
}

func (w *Worker) Stop() {
	w.cancelFunc()
	w.wg.Wait()
}

func (w *Worker) ordersGetter(ctx context.Context) {
	defer w.wg.Done()
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			orders, err := w.repo.GetWaitingOrders()
			if err != nil {
				log.Err(err).Stack().Send()
				break
			}
			fmt.Println("***", orders)
			for i := range orders {
				w.ch <- orders[i]
			}
		}
	}
}

func (w *Worker) spawnWorkers(ctx context.Context) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case orderID := <-w.ch:
			response, err := w.client.Get(configs.AccrualSystemAddress + "/api/orders/" + orderID)
			if err != nil {
				log.Error().Err(err).Send()
				break
			}
			if response.StatusCode == 429 {
				time.Sleep(time.Second)
				log.Info().Msg("to much")
				break
			}
			var input models.Order
			err = render.DecodeJSON(response.Body, &input)
			response.Body.Close()
			if err != nil {
				log.Error().Err(err).Stack().Send()
				break
			}
			fmt.Println("+++", input)
			err = w.repo.UpdateOrder(input)
			if err != nil {
				log.Error().Err(err).Stack().Send()
			}
		}
	}
}
