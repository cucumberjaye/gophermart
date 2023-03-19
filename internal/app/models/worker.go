package models

import "encoding/json"

type OrderAccrual struct {
	ID      string `json:"number"`
	Accrual int    `json:"accrual,omitempty"`
	Status  status `json:"status"`
}

func (o *OrderAccrual) UnmarshalJSON(data []byte) error {
	type OrderAlias OrderAccrual

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
