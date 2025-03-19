package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date,title,comment,repeat) VALUES
	(:date,:title,:comment,:repeat)`
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title),
		sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int, search string) ([]*Task, error) {
	var ret []*Task
	var err error
	if len(search) == 0 {
		query := `SELECT * FROM scheduler ORDER BY date LIMIT :limit`
		err = db.Select(&ret, query, sql.Named("limit", limit))
	} else {
		t, err := time.Parse(`02.01.2006`, search)
		if err == nil {
			query := `SELECT * FROM scheduler WHERE date = :date LIMIT :limit`
			err = db.Select(&ret, query, sql.Named("date", t.Format(`20060102`)), sql.Named("limit", limit))
		} else {
			query := `SELECT * FROM scheduler WHERE title LIKE :search OR 
		comment LIKE :search ORDER BY date LIMIT :limit`
			search = fmt.Sprintf("%%%s%%", search)
			err = db.Select(&ret, query, sql.Named("search", search), sql.Named("limit", limit))
		}
	}
	if ret == nil {
		ret = make([]*Task, 0)
	}
	return ret, err
}

func GetTask(id string) (*Task, error) {
	var ret Task
	query := `SELECT * FROM scheduler WHERE id = :id`
	err := db.Get(&ret, query, sql.Named("id", id))
	return &ret, err
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = :id`
	res, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for deleting task`)
	}

	return err
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date=:date,title=:title,comment=:comment,repeat=:repeat WHERE id = :id`
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title),
		sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat), sql.Named("id", task.ID))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func UpdateDate(next, id string) error {
	res, err := db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
		sql.Named("date", next), sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for todo task`)
	}
	return nil
}
