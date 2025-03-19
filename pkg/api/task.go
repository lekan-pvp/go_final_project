package api

import (
	"go1f/pkg/db"
	"net/http"
)

type TaskResp struct {
	*db.Task
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if len(id) == 0 {
		errorJson(w, "не указан идентификатор")
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		errorJson(w, err)
		return
	}
	writeJson(w, TaskResp{
		Task: task,
	})
}
