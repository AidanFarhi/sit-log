package main

import (
	"database/sql"
	"fmt"
	"log"

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

	// run a test query
	event, err := repo.GetEvent(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(event)
}
