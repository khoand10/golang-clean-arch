package main

import (
	"golang-clean-arch/api"
	"golang-clean-arch/database"
	"log"
	"net/http"
)

func main() {
	root := api.Init()

	_, closeFunc, err := database.InitSqlite()
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":8080", root))
	defer func() {
		closeFunc()
	}()
}
