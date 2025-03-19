package api

import (
	"go1f/pkg/db"
	"net/http"
)

func delTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if len(id) == 0 {
		errorJson(w, `unknown id to delete`)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		errorJson(w, err)
		return
	}

	writeJson(w, Resp{})
}
