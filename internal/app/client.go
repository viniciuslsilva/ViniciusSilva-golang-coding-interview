package app

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

const (
	baseUrl = "https://api.ers.usda.gov/data"

	statesUri = "/arms/state"
	reportUri = "arms/report"
)

func FetchStates() (StatesResponse, error) {
	resp, err := getBaseRequest().Get(statesUri)
	if err != nil {
		return StatesResponse{}, err
	}
	if resp.IsError() {
		return StatesResponse{}, errors.New(resp.String())
	}

	var statesData StatesResponse
	err = json.Unmarshal(resp.Body(), &statesData)
	if err != nil {
		return StatesResponse{}, err
	}

	return statesData, nil
}

func FetchReport() (ReportResponse, error) {
	resp, err := getBaseRequest().Get(reportUri)
	if err != nil {
		return ReportResponse{}, err
	}
	if resp.IsError() {
		return ReportResponse{}, errors.New(resp.String())
	}

	var reportsData ReportResponse
	err = json.Unmarshal(resp.Body(), &reportsData)
	if err != nil {
		return ReportResponse{}, err
	}

	return reportsData, nil
}

func getBaseRequest() *resty.Request {
	req := resty.New().SetHostURL(baseUrl).R()

	req.SetHeader("Accept", "application/json")
	req.SetError(DefaultError{"An error occurred"})

	req.SetQueryParams(map[string]string{
		"api_key": viper.GetString("api-key"),
	})
	return req
}

type DefaultError struct {
	Message string
}
