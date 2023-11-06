package main

import (
	"time"

	"github.com/google/uuid"
)

type Alert struct {
	Id          string
	AlertType   string
	Time        int64
	Description string
	CreatedAt   time.Time
}

type CreateAlertReq struct {
	AlertType   string
	Time        int64
	Description string
}

func NewAlert(alertType string, alertTime int64, description string) (*Alert, error) {

	return &Alert{
		Id:          uuid.New().String(),
		AlertType:   alertType,
		Time:        alertTime,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil

}
