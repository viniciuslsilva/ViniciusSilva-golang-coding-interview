package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	ID   string `json: "id"`
	Code string `json: "code"`
	Name string `json: "name"`
}

// select * from state where code =?/ like %%

type Report struct {
	Num         int     `json:"num"`
	Header      string  `json:"header"`
	Description *string `json:"description"`
	Terms       string  `json:"terms"`
}
