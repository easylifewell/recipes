package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"

	"github.com/Sirupsen/logrus"
)

const (
	dbname string = "./recipes.db"
)

//CreateDB create the database for recipes recommend
func CreateDB() {
	os.Remove(dbname)

	db, err := sql.Open("sqlite3", dbname)
	defer db.Close()

	sqlStmt := `
	create table clocks (
		id		INTEGER PRIMARY KEY AUTOINCREMENT,
		mtype		TEXT,
		clock_id	TEXT,
		begin_time	INTEGER,
		end_time	INTEGER,
		organ		TEXT,
		UNIQUE(clock_id)
	);
	delete from clocks;
	`
	clocksData := []struct {
		mytype    string
		clockID   string
		beginTime int
		endTime   int
		organ     string
	}{
		{"早餐", "丑", 1, 3, "肝"},
		{"早餐", "寅", 3, 5, "肺"},
		{"早餐", "卯", 5, 7, "大肠"},
		{"早餐", "辰", 7, 9, "胃"},
		{"午餐", "已", 9, 11, "脾"},
		{"午餐", "午", 11, 13, "心"},
		{"午餐", "未", 13, 15, "小肠"},
		{"午餐", "申", 15, 17, "膀胱"},
		{"晚餐", "酉", 17, 19, "肾"},
		{"晚餐", "戌", 19, 21, "心包"},
		{"晚餐", "亥", 21, 23, "三焦"},
		{"晚餐", "子", 23, 1, "胆"},
	}
	_, err = db.Exec(sqlStmt)
	if err != nil {
		logrus.Errorf("%q: %s\n", err, sqlStmt)
	}

	tx, err := db.Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into clocks(mtype, clock_id, begin_time, end_time, organ) values (?, ?, ?, ?, ?)")
	if err != nil {
		logrus.Fatal(err)
	}
	defer stmt.Close()

	for _, d := range clocksData {
		_, err = stmt.Exec(d.mytype, d.clockID, d.beginTime, d.endTime, d.organ)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, clock_id, mtype, begin_time, end_time, organ from clocks")
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var clockID string
		var mType string
		var beginTime int
		var endTime int
		var organ string
		err = rows.Scan(&id, &clockID, &mType, &beginTime, &endTime, &organ)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(id, mType, clockID, beginTime, endTime, organ)
	}
	err = rows.Err()
	if err != nil {
		logrus.Fatal(err)
	}
}
