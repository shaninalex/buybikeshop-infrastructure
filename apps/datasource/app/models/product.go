package models

import "time"

type Product struct {
	ID               uint64
	Title            string
	Description      string
	ShortDescription string
	CreatedAt        time.Time
}
