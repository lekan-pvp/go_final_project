package api

import (
	"net/http"

	"go1f/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50, r.FormValue("search"))
	if err != nil {
		errorJson(w, err)
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
