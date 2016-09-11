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
	if err != nil {
		logrus.Fatalf("open sqlite3 db error: %q", err)
	}
	defer db.Close()

	createClocks(db)
	createFoodMatch(db)
	createFoodList(db)
}
func createFoodList(db *sql.DB) {
	sqlStmt := `
	create table food_list(
		id              INTEGER PRIMARY KEY AUTOINCREMENT,
		food_id		TEXT,
		food_name	TEXT,
		kind_id		TEXT,
		kind_name	TEXT,
		UNIQUE(food_id)
	);
	delete from food_list;
	`

	foodListData := []struct {
		foodID   string
		foodName string
		kindID   string
		kindName string
	}{
		{"FF1", "苹果", "C2", "水果"},
		{"FF2", "香蕉", "C2", "水果"},
		{"FF3", "梨", "C2", "水果"},
		{"FF4", "芒果", "C2", "水果"},
		{"FF5", "菠萝", "C2", "水果"},
		{"FF6", "草莓", "C2", "水果"},
		{"FF7", "杏", "C2", "水果"},
	}

	_, err := db.Exec(sqlStmt)
	if err != nil {
		logrus.Errorf("%q: %s\n", err, sqlStmt)
	}

	tx, err := db.Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into food_list(food_id, food_name, kind_id, kind_name) values (?, ?, ?, ?)")
	if err != nil {
		logrus.Fatal(err)
	}
	defer stmt.Close()

	for _, d := range foodListData {
		_, err = stmt.Exec(d.foodID, d.foodName, d.kindID, d.kindName)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, food_id, food_name, kind_id, kind_name from food_list")
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var foodID string
		var foodName string
		var kindID string
		var kindName string
		err = rows.Scan(&id, &foodID, &foodName, &kindID, &kindName)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(id, foodID, foodName, kindID, kindName)
	}
	err = rows.Err()
	if err != nil {
		logrus.Fatal(err)
	}
}

func createFoodMatch(db *sql.DB) {
	sqlStmt := `
	create table food_matches (
		id		INTEGER PRIMARY KEY AUTOINCREMENT,
		mtype		TEXT,
		food_id		TEXT,
		food_des	TEXT,
		priority	INTEGER
	);
	delete from food_matches;
	`

	foodMatchesData := []struct {
		mtype    string
		foodID   string
		foodDes  string
		priority int
	}{
		{"早餐", "C2", "水果", 2},
		{"早餐", "C3", "坚果干果", 3},
		{"早餐", "C5", "蔬菜", 1},
		{"早餐", "C8", "奶制品", 1},
		{"午餐", "C1", "肉类", 1},
		{"午餐", "C2", "水果", 2},
		{"午餐", "C5", "蔬菜", 1},
		{"晚餐", "C8", "奶制品", 1},
		{"晚餐", "C1", "肉类", 3},
		{"晚餐", "C2", "水果", 1},
		{"晚餐", "C3", "坚果干果", 2},
		{"晚餐", "C5", "蔬菜", 1},
	}
	_, err := db.Exec(sqlStmt)
	if err != nil {
		logrus.Errorf("%q: %s\n", err, sqlStmt)
	}

	tx, err := db.Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into food_matches(mtype, food_id, food_des, priority) values (?, ?, ?, ?)")
	if err != nil {
		logrus.Fatal(err)
	}
	defer stmt.Close()

	for _, d := range foodMatchesData {
		_, err = stmt.Exec(d.mtype, d.foodID, d.foodDes, d.priority)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, mtype, food_id, food_des, priority from food_matches")
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var mType string
		var foodID string
		var foodDes string
		var priority int
		err = rows.Scan(&id, &mType, &foodID, &foodDes, &priority)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(id, mType, foodID, foodDes, priority)
	}
	err = rows.Err()
	if err != nil {
		logrus.Fatal(err)
	}
}

func createClocks(db *sql.DB) {
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
	_, err := db.Exec(sqlStmt)
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
