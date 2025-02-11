package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
)

func GetEventsForChild(db *sql.DB, childID int, adultID int) ([]model.Event, error) {
	events := []model.Event{}
	query := `
		SELECT e.timestamp, e.type, e.description, e.start_time, e.end_time, e.duration
		FROM event e
		JOIN adult_child_relation acr
		ON e.child_id = acr.child_id
		WHERE acr.child_id = ? AND acr.adult_id = ?
	`
	rows, err := db.Query(query, childID, adultID)
	defer rows.Close()
	if err != nil {
		return events, err
	}
	for rows.Next() {
		event := model.Event{}
		err := rows.Scan(
			&event.TimeStamp, &event.Type, &event.Description,
			&event.StartTime, &event.EndTime, &event.Duration,
		)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	return events, nil
}

func CreateEvent(db *sql.DB, r *http.Request) error {
	// TODO: actually add this
	// duration, err := calculateDuration(newEvent.StartTime, newEvent.EndTime)
	// newEvent.Duration = duration

	// this is so SQLite actually uses foreign keys
	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	childID, err := GetChildID(db, r.FormValue("child_name"))
	if err != nil {
		fmt.Println("error getting childId")
		return err
	}
	query := `
		INSERT INTO event(child_id, type, description, start_time, end_time, duration)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(
		query,
		childID, r.FormValue("type"), r.FormValue("description"),
		r.FormValue("start_time"), r.FormValue("end_time"), "0",
	)
	if err != nil {
		fmt.Println("error inserting new event into table")
		return err
	}
	return nil
}

// func calculateDuration(startTime string, endTime string) (string, error) {
// 	layout := "15:04:05"
// 	start, err := time.Parse(layout, startTime)
// 	if err != nil {
// 		return "", fmt.Errorf("invalid start time format: %v", err)
// 	}
// 	end, err := time.Parse(layout, endTime)
// 	if err != nil {
// 		return "", fmt.Errorf("invalid end time format: %v", err)
// 	}
// 	duration := end.Sub(start)
// 	if duration < 0 {
// 		duration += 24 * time.Hour
// 	}
// 	hours := int(duration.Hours())
// 	minutes := int(duration.Minutes()) % 60
// 	seconds := int(duration.Seconds()) % 60
// 	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
// }
