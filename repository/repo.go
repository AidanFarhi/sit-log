package repository

import (
	"database/sql"

	"github.com/AidanFarhi/sitlog/model"
	_ "github.com/mattn/go-sqlite3"
)

type EventRepository interface {
	GetEvent(ID int) (model.Event, error)
	GetEventsForChildID(childID int) ([]model.Event, error)
	CreateEvent(event model.Event) error
}

type SQLiteEventRepository struct {
	DB *sql.DB
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
