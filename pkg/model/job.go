package model

import (
	"time"
)

//https://github.com/asaskevich/govalidator TODO add in a validator

type Job struct {
	UUID                  string    `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name                  string    `json:"name"`
	Description           string    `json:"description,omitempty"`
	Frequency             string    `json:"frequency,omitempty" sql:"frequency"`
	StartDate             time.Time `json:"start_date,omitempty" sql:"start_date"`
	EndDate               time.Time `json:"end_date,omitempty" sql:"end_date"`
	Enable                *bool     `json:"enable"`
	DestroyAfterCompleted bool      `json:"destroy_after_completed,omitempty" sql:"destroy_after_completed"`
	CreatedAt             time.Time `json:"created_on,omitempty"`
	UpdatedAt             time.Time `json:"updated_on,omitempty"`
}
