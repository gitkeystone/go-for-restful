package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, err := dbDriver.Prepare(train)
	if err != nil {
		log.Println(err)
	}

	// Create train table
	_, err = statement.Exec()
	if err != nil {
		log.Println("Table already exists!")
	}

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()

	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()

	log.Println("All tables created/initialized successfully!")
}
