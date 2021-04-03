package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db, _ = sql.Open("mysql", "root:password@tcp(mysql:3306)/health")
)

func RecordStepsData(steps int, name, dt string) {
	fmt.Printf("steps=%d, name=%s, dt=%s\n", steps, name, dt)
	// perform a db.Query insert
	insert, err := db.Query(fmt.Sprintf("INSERT INTO steps (steps, name, activity_date) VALUES ( %d, '%s', '%s' )", steps, name, dt))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func RecordIDData(name, dt string) {
	// perform a db.Query insert
	insert, err := db.Query(fmt.Sprintf("INSERT INTO id (name, activity_date) VALUES ( '%s', '%s' )", name, dt))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func RecordActivityData(name, dt, activityName string, activityDuration int) {
	// perform a db.Query insert
	insert, err := db.Query(fmt.Sprintf("INSERT INTO activity (name, activity_date, activity_name, activity_duration) VALUES ( '%s', '%s', '%s', %d )", name, dt, activityName, activityDuration))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
