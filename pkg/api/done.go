package api

import (
	"net/http"
	"time"

	"go1f/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if len(id) == 0 {
		errorJson(w, `unknown id to done`)
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		errorJson(w, err)
		return
	}
	if len(task.Repeat) == 0 {
		if err = db.DeleteTask(id); err != nil {
			errorJson(w, err)
			return
		}
		writeJson(w, Resp{})
		return
	}
	t, err := time.Parse(formatDate, task.Date)
	if err != nil {
		errorJson(w, err)
		return
	}
	now := time.Now()
	if now.Format(formatDate) < t.Format(formatDate) {
		now = t
	}
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		errorJson(w, err)
		return
	}
	if err = db.UpdateDate(next, id); err != nil {
		errorJson(w, err)
		return
	}
	writeJson(w, Resp{})
}
