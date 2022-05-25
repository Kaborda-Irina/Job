package controllers

import (
	"Job/server/repo"
	"fmt"
	"net/http"
)

func PutData(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, repo.PutTable())
}
