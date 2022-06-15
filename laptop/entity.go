package laptop

import (
	"time"
)

type Gaming struct {
	ID         int       `json:"id"`
	Brand      string    `json:"brand"`
	Price      int       `json:price"`
	Location   string    `json:"location"`
	Processor  string    `json:"processor"`
	Generation int       `json:"gen"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at`
}
