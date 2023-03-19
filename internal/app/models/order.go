package models

import (
	"encoding/json"
	"time"
)

type status int

func (s status) String() string {
	switch s {
	case 0:
		return "NEW"
	case 1:
		return "PROCESSING"
	case 2:
		return "INVALID"
	case 3:
		return "PROCESSED"
	}

	return ""
}

const (
	New status = iota
	Processing
	Invalid
	Processed
)

type Order struct {
	Id         string    `json:"number"`
	UserId     string    `json:"-"`
	Accrual    int       `json:"accrual,omitempty"`
	Status     status    `json:"status"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type OrderAlias Order

	aliasValue := struct {
		OrderAlias
		Accrual    float32 `json:"accrual"`
		Status     string  `json:"status"`
		UploadedAt string  `json:"uploaded_at"`
	}{
		OrderAlias: OrderAlias(o),
		Accrual:    float32(o.Accrual) / 100,
		Status:     o.Status.String(),
		UploadedAt: o.UploadedAt.Format(time.RFC3339),
	}

	return json.Marshal(aliasValue)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	type OrderAlias Order

	aliasValue := &struct {
		*OrderAlias
		Status  string  `json:"status"`
		Accrual float32 `json:"accrual"`
	}{
		OrderAlias: (*OrderAlias)(o),
	}

	if err := json.Unmarshal(data, aliasValue); err != nil {
		return err
	}
	switch aliasValue.Status {
	case "REGISTERED":
		o.Status = 0
	case "PROCESSING":
		o.Status = 1
	case "INVALID":
		o.Status = 2
	case "PROCESSED":
		o.Status = 3
	}
	o.Accrual = int(aliasValue.Accrual * 100)

	return nil
}
