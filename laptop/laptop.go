package laptop

import "encoding/json"

type Laptop struct {
	Id         json.Number `binding:"required"`
	Price      json.Number `binding:"required,numeric"`
	Brand      string
	Stock      int32
	LocationWh string `json:"location_wh"`
}
