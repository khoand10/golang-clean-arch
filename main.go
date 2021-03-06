package main

import (
	"golang-clean-arch/api"
	"golang-clean-arch/database"
	"log"
	"net/http"
)

func main() {
	//root := api.Init()
	handler := api.NewHandler()

	_, closeFunc, err := database.InitSqlite()
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":8080", handler))

	defer func() {
		closeFunc()
	}()
}
