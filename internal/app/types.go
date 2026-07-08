package app

import (
	"github.com/viniciuslsilva/ViniciusSilva-golang-coding-interview/internal/models"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}

type StatesResponse struct {
	BaseResponse
	States []models.State `json:"data"`
}

type ReportResponse struct {
	BaseResponse
	Report []models.Report `json:"data"`
}
