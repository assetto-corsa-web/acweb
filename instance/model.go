package instance

import (
	"time"
)

type Log struct {
	File string    `json:"file"`
	Date time.Time `json:"date"`
	Size int64     `json:"size"`
}
