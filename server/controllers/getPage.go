package controllers

import (
	"Job/server/repo"
	"fmt"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	allDataFromDB := repo.GetData()
	for _, data := range allDataFromDB {
		fmt.Fprint(w, data)
	}
}
