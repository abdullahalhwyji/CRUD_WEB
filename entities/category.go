package entities

import "time"

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}