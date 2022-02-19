package entity

import "time"

type LogMail struct {
	Id        int64     `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type MailInput struct {
	To      string `json:"to"`
	Message string `json:"message"`
}
