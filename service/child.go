package service

import "database/sql"

func GetChildID(db *sql.DB, childName string) (int, error) {
	// add a check for adult id
	var childId int
	err := db.QueryRow("SELECT id FROM child WHERE name = ?", childName).Scan(&childId)
	if err != nil {
		return -1, err
	}
	return childId, nil
}
