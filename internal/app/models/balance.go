package models

import (
	"encoding/json"
	"time"
)

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

func (b Balance) MarshalJSON() ([]byte, error) {
	aliasValue := struct {
		Current   float32 `json:"current"`
		Withdrawn float32 `json:"withdrawn"`
	}{
		Current:   float32(b.Current) / 100,
		Withdrawn: float32(b.Withdrawn) / 100,
	}

	return json.Marshal(aliasValue)
}

func (b *Balance) UnmarshalJSON(data []byte) error {
	aliasValue := &struct {
		Current   float32 `json:"current"`
		Withdrawn float32 `json:"withdrawn"`
	}{}

	if err := json.Unmarshal(data, aliasValue); err != nil {
		return err
	}

	b.Current = int(aliasValue.Current * 100)
	b.Withdrawn = int(aliasValue.Withdrawn * 100)

	return nil
}

type Withdraw struct {
	Order       string    `json:"order"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"-"`
}

func (w Withdraw) MarshalJSON() ([]byte, error) {
	type WithdrawAlias Withdraw

	aliasValue := struct {
		WithdrawAlias
		Sum         float32 `json:"sum"`
		ProcessedAt string  `json:"processedAt"`
	}{
		WithdrawAlias: WithdrawAlias(w),
		Sum:           float32(w.Sum) / 100,
		ProcessedAt:   w.ProcessedAt.Format(time.RFC3339),
	}

	return json.Marshal(aliasValue)
}

func (w *Withdraw) UnmarshalJSON(data []byte) error {
	type WithdrawAlias Withdraw

	aliasValue := &struct {
		*WithdrawAlias
		Sum float32 `json:"sum"`
	}{
		WithdrawAlias: (*WithdrawAlias)(w),
	}
	if err := json.Unmarshal(data, aliasValue); err != nil {
		return err
	}

	w.Sum = int(aliasValue.Sum * 100)

	return nil
}
