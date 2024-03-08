package models

import (
	"conference-booking-rest-api/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Conference struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (c Conference) Save() error {
	query := `INSERT INTO conferences(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	result, err := stmt.Exec(c.Name, c.Description, c.Location, c.DateTime, c.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	c.ID = id
	return err
}

func GetAllConferences() ([]Conference, error) {
	query := "SELECT * FROM conferences"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var conferences []Conference

	for rows.Next() {
		var conference Conference
		err := rows.Scan(&conference.ID, &conference.Name, &conference.Description, &conference.Location, &conference.DateTime, &conference.UserID)

		if err != nil {
			return nil, err
		}
		conferences = append(conferences, conference)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return conferences, nil

}

func GetConferenceByID(id int64) (*Conference, error) {
	query := "SELECT * FROM conferences WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	
	var conference Conference
	err := row.Scan(&conference.ID, &conference.Name, &conference.Description, &conference.Location, &conference.DateTime, &conference.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("conference with ID %d not found", id)
		} else {
			log.Fatal(err)
		}
	}
	return &conference, nil
}

func (c Conference) Update() error {
	query := `
	UPDATE conferences
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(c.Name, c.Description, c.Location, c.DateTime, c.ID)
	return err
}

func (c Conference) Delete() error {
	query := `DELETE FROM conferences WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(c.ID)
	return err
}
