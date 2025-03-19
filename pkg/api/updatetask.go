package api

import (
	"net/http"

	"go1f/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := parseTask(r)
	if err != nil {
		errorJson(w, err)
		return
	}
	if len(task.ID) == 0 {
		errorJson(w, `undefined task id`)
		return
	}
	if len(task.Title) == 0 {
		errorJson(w, errTitle)
		return
	}
	if err = checkDate(task); err != nil {
		errorJson(w, err)
		return
	}
	if err = db.UpdateTask(task); err != nil {
		errorJson(w, err)
		return
	}
	writeJson(w, Resp{})
}
