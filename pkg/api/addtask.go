package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go1f/pkg/db"
)

var (
	errPost  = errors.New("expected POST method")
	errTitle = errors.New("task title is empty")
)

type Resp struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func errorJson(w http.ResponseWriter, msg any) {
	writeJson(w, Resp{
		Error: fmt.Sprint(msg),
	})
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}

func parseJson(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = r.Body.Close(); err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func parseTask(r *http.Request) (*db.Task, error) {
	var task db.Task

	err := parseJson(r, &task)
	return &task, err
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(formatDate)
	}
	t, err := time.Parse(formatDate, task.Date)
	if err != nil {
		return err
	}
	task.Date = t.Format(formatDate)
	var next string
	if len(task.Repeat) > 0 {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(formatDate)
		} else {
			task.Date = next
		}
	}
	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := parseTask(r)
	if err != nil {
		errorJson(w, err)
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
	var id int64
	if id, err = db.AddTask(task); err != nil {
		errorJson(w, err)
		return
	}

	writeJson(w, Resp{ID: strconv.FormatInt(id, 10)})
}
