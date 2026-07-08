package main

import (
	"github.com/viniciuslsilva/ViniciusSilva-golang-coding-interview/internal/app"
	"github.com/viniciuslsilva/ViniciusSilva-golang-coding-interview/internal/config"
	"github.com/viniciuslsilva/ViniciusSilva-golang-coding-interview/internal/models"
	"github.com/viniciuslsilva/ViniciusSilva-golang-coding-interview/pkg/common"
)

func main() {
	config.LoadConfig([]string{"./internal/config"}, "config")
	db, err := common.ConnectDBWithConfig()
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&models.State{})

	app.Start(db)
}
