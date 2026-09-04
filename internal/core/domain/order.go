package core_domain

import "time"

type Order struct {
	Number     string
	Status     string
	Accrual    *float64
	UploadedAt time.Time
}
