// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type History struct {
	ID        int64     `json:"id"`
	IDAccount int64     `json:"id_account"`
	Amount    int64     `json:"amount"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID           int64     `json:"id"`
	IDAccount    int64     `json:"id_account"`
	IDService    int64     `json:"id_service"`
	PriceService int64     `json:"price_service"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type Service struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}