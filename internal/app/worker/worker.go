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
	fmt.Println("worker running")
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
	for {
		select {
		case <-ctx.Done():
			return
		default:
			orders, err := w.repo.GetWaitingOrders()
			if err != nil {
				fmt.Println(err)
			}
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
		case orderId := <-w.ch:
			response, err := w.client.Get(configs.AccrualSystemAddress + "/api/orders/" + orderId)
			if err != nil {
				log.Err(err).Send()
				break
			}
			if response.StatusCode == 429 {
				time.Sleep(time.Second)
				break
			}
			var input models.Order
			err = render.DecodeJSON(response.Body, &input)
			if err != nil {
				log.Err(err).Send()
				break
			}
			err = w.repo.UpdateOrder(input)
			if err != nil {
				log.Err(err).Send()
			}
		}
	}
}
