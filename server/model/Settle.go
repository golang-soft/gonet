package model

import (
	"time"
)

type (
	Settle struct {
		table string `sql:"table;name:settle"`

		Id    int64     `sql:"primary;name:id" json:"id"`
		Round int64     `sql:"name:round" json:"round"`
		Data  string    `sql:"name:data" json:"data"`
		Ts    time.Time `sql:"datetime;name:ts" json:"ts"`
	}
)
