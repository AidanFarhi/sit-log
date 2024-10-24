package repository

import (
	"database/sql"

	"github.com/AidanFarhi/sitlog/model"
)

type Repository interface {
	GetEventsForChild(childID int, adultID int) ([]model.Event, error)
}

type SimpleRepository struct {
	DB *sql.DB
}

func NewSimpleRepository(db *sql.DB) SimpleRepository {
	return SimpleRepository{DB: db}
}

func (sr SimpleRepository) GetEventsForChild(childID int, adultID int) ([]model.Event, error) {
	events := []model.Event{}
	query := `
		SELECT
			e.timestamp,
			e.type,
			e.description,
			e.start_time,
			e.end_time,
			e.duration
		FROM
			event e
		JOIN
			adult_child_relation acr
		ON
			e.child_id = acr.child_id
		WHERE
			acr.child_id = ?
			AND acr.adult_id = ?
	`
	rows, err := sr.DB.Query(query, childID, adultID)
	defer rows.Close()
	if err != nil {
		return events, err
	}
	for rows.Next() {
		event := model.Event{}
		err := rows.Scan(
			&event.TimeStamp,
			&event.Type,
			&event.Description,
			&event.StartTime,
			&event.EndTime,
			&event.Duration,
		)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	return events, nil
}
