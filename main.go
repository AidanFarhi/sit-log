package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// events := []model.Event{}
	reader, _ := os.Open("local_data/data.csv")
	csvReader := csv.NewReader(reader)
	rows, _ := csvReader.ReadAll()
	for _, row := range rows {
		event, description, duration := row[1], row[2], row[5]
		timeStamp, startTime, endTime := row[0], row[3], row[4]
		date := strings.Split(timeStamp, " ")[0]
		startTimeStamp := date + " " + startTime
		endTimeStamp := date + " " + endTime
		fmt.Println(timeStamp, event, description, startTimeStamp, endTimeStamp, duration)
	}

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure the database is available
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
