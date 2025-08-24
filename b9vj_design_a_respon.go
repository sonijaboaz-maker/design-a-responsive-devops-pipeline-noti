package main

import (
	"time"
)

type Notifier struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Environment string    `json:"environment"`
	PipelineID  string    `json:"pipeline_id"`
	Stage       string    `json:"stage"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Notification struct {
	ID        string `json:"id"`
	NotifierID string `json:"notifier_id"`
	Type      string `json:"type"` // email, slack, etc.
	Message   string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type DevOpsPipeline struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Environment string    `json:"environment"`
	Stages     []Stage    `json:"stages"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Stage struct {
	ID        string `json:"id"`
	PipelineID string `json:"pipeline_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}