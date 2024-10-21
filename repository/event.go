package repository

import (
	"database/sql"

	"github.com/AidanFarhi/sitlog/model"
	_ "github.com/mattn/go-sqlite3"
)

type EventRepository interface {
	GetEvent(ID int) (model.Event, error)
	GetEventsForAdult(adultID int) ([]model.Event, error)
	CreateEvent(event model.Event) error
}

type SQLiteEventRepository struct {
	DB *sql.DB
}

func NewSQLiteEventRepository(db *sql.DB) SQLiteEventRepository {
	return SQLiteEventRepository{DB: db}
}

func (ser SQLiteEventRepository) GetEvent(ID int) (model.Event, error) {
	event := model.Event{}
	query := `
		SELECT
			id,
			adult_id,
			child_id,
			timestamp,
			type,
			description,
			start_time,
			end_time,
			duration
		FROM
			event
		WHERE
			id = ?
	`
	row := ser.DB.QueryRow(query, ID)
	err := row.Scan(
		&event.ID,
		&event.AdultID,
		&event.ChildID,
		&event.TimeStamp,
		&event.Type,
		&event.Description,
		&event.StartTime,
		&event.EndTime,
		&event.EventDuration,
	)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (ser SQLiteEventRepository) GetEventsForAdult(adultID int) ([]model.Event, error) {
	events := []model.Event{}
	query := `
		SELECT
			id,
			adult_id,
			child_id,
			timestamp,
			type,
			description,
			start_time,
			end_time,
			duration
		FROM
			event
		WHERE
			adult_id = ?
	`
	rows, err := ser.DB.Query(query, adultID)
	defer rows.Close()
	if err != nil {
		return events, err
	}
	for rows.Next() {
		event := model.Event{}
		err := rows.Scan(
			&event.ID,
			&event.AdultID,
			&event.ChildID,
			&event.TimeStamp,
			&event.Type,
			&event.Description,
			&event.StartTime,
			&event.EndTime,
			&event.EventDuration,
		)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (ser SQLiteEventRepository) CreateEvent(event model.Event) error {
	query := `
		INSERT INTO 
			event 
				(adult_id, child_id, type, description, start_time, end_time, duration)
			values
				(?, ?, ?, ?, ?, ?, ?)
	`
	_, err := ser.DB.Exec(
		query,
		event.AdultID,
		event.ChildID,
		event.Type,
		event.Description,
		event.StartTime,
		event.EndTime,
		event.EventDuration,
	)
	if err != nil {
		return err
	}
	return nil
}
