package models

import (
	"database/sql"
	// "errors"
	"time"
)

type Book struct{
	bookID		int
	Title		string
	Author		string
	About		string
	Year		string 		// publish year
	Status 		bool 		// completed - 1, reading - 0
	Created		time.Time
}

type BookModel struct{
	DB * sql.DB
}

func (m *BookModel) Insert(task string) (int, error){
	stmt := `INSERT INTO tasks (task, created)
    VALUES(?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, task)	
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func(m* BookModel) Get(id int) (*Book, error){
	return nil, nil
}

func (m* BookModel) Latest() ([]*Book, error){
	return nil, nil
}

func (m* BookModel) Delete(id int) error{
	return nil
}