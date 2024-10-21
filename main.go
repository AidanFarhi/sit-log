package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/repository"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// events := []model.Event{}
	// reader, _ := os.Open("local_data/data.csv")
	// csvReader := csv.NewReader(reader)
	// rows, _ := csvReader.ReadAll()
	// for _, row := range rows {
	// 	event, description, duration := row[1], row[2], row[5]
	// 	timeStamp, startTime, endTime := row[0], row[3], row[4]
	// 	date := strings.Split(timeStamp, " ")[0]
	// 	startTimeStamp := date + " " + startTime
	// 	endTimeStamp := date + " " + endTime
	// 	fmt.Println(timeStamp, event, description, startTimeStamp, endTimeStamp, duration)
	// }

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initalize the repository
	repo := repository.SQLiteEventRepository{
		DB: db,
	}

	// Define some initial parameters
	eventTypes := []string{"Babysitting", "Diaper Change", "Snack", "Playdate", "Daycare"}
	descriptions := []string{
		"Watched the child for a while.",
		"Changed the child's diaper.",
		"Gave the child a snack.",
		"Playdate with friends.",
		"Full day daycare service.",
	}
	numEvents := 10 // Number of events to generate

	// Slice to store the generated events
	var eventsToInsert []model.Event

	for i := 1; i <= numEvents; i++ {
		// Generate random data for each event (simple example)
		eventTypeIndex := i % len(eventTypes)
		adID := (i % 3) + 1    // Cycle through adult IDs (1-3)
		childID := (i % 4) + 1 // Cycle through child IDs (1-4)
		startTime := time.Now().Add(time.Duration(i) * time.Hour).Format("2006-01-02 15:04")
		endTime := time.Now().Add(time.Duration(i+1) * time.Hour).Format("2006-01-02 17:04")
		eventDuration := fmt.Sprintf("%d hours", i)

		event := model.Event{
			ID:            0,
			AdultID:       adID,
			ChildID:       childID,
			TimeStamp:     "",
			Type:          eventTypes[eventTypeIndex],
			Description:   descriptions[eventTypeIndex],
			StartTime:     startTime,
			EndTime:       endTime,
			EventDuration: eventDuration,
		}

		// Append the event to the slice
		eventsToInsert = append(eventsToInsert, event)
	}

	for _, e := range eventsToInsert {
		repo.CreateEvent(e)
	}

	eventsFromDb, _ := repo.GetEventsForChildID(2)
	for _, e := range eventsFromDb {
		fmt.Println(e)
	}
}
