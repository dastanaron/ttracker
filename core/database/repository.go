package database

import (
	"time"
	e "ttracker/core/errors"
	"ttracker/core/utils"
)

type TaskModel struct {
	Id        int
	Name      string
	StartTime time.Time
	Duration  int
	Project   string
}

type ServiceData struct {
	Name string
	Data string
}

func AddRow(model TaskModel) {
	db := GetConnection()

	_, err := db.Exec("INSERT INTO tasks (name, startTime, endTime, duration, project) VALUES (?, ?, ?, ?, ?)", model.Name, model.StartTime.Unix(), time.Now().Unix(), model.Duration, model.Project)
	e.CheckError("Error record data to database", err)
}

func Save(model TaskModel) {
	db := GetConnection()

	var foundModel TaskModel

	row := db.QueryRow("select id, duration from tasks where name = ?", model.Name)

	row.Scan(&foundModel.Id, &foundModel.Duration)

	if foundModel.Id != 0 {
		model.Duration = foundModel.Duration + model.Duration
		_, err := db.Exec("UPDATE tasks SET duration = ?, endTime = ? WHERE id = ?", model.Duration, time.Now().Unix(), foundModel.Id)
		e.CheckError("Error update data in database", err)
	} else {
		AddRow(model)
	}
}

func GetLatest(count int) []TaskModel {
	db := GetConnection()

	result := []TaskModel{}

	rows, err := db.Query("select id, name, startTime, duration, project from tasks order by startTime desc limit ?", count)
	e.CheckError("Error receive data from database", err)

	for rows.Next() {
		var row TaskModel
		var startTimeUnix int
		err = rows.Scan(&row.Id, &row.Name, &startTimeUnix, &row.Duration, &row.Project)
		if err != nil {
			e.CheckError("Error transform data from database to TaskModel", err)
		}

		row.StartTime = time.Unix(int64(startTimeUnix), 0)

		result = append(result, row)
	}

	rows.Close()

	return result
}

func GetToDay() []TaskModel {
	db := GetConnection()

	bod := utils.Bod(time.Now())

	result := []TaskModel{}

	rows, err := db.Query("select id, name, startTime, duration, project from tasks where endTime >= ?", bod.Unix())
	e.CheckError("Error receive data from database", err)

	for rows.Next() {
		var row TaskModel
		var startTimeUnix int
		err = rows.Scan(&row.Id, &row.Name, &startTimeUnix, &row.Duration, &row.Project)
		if err != nil {
			e.CheckError("Error transform data from database to TaskModel", err)
		}

		row.StartTime = time.Unix(int64(startTimeUnix), 0)

		result = append(result, row)
	}

	rows.Close()

	return result
}

func SaveServiceData(name, data string) {
	db := GetConnection()

	var foundModel ServiceData

	row := db.QueryRow("select name from service_data where name = ?", name)
	row.Scan(&foundModel.Name)

	if foundModel.Name != "" {
		_, err := db.Exec("UPDATE service_data SET data = ? WHERE name = ?", data, foundModel.Name)
		e.CheckError("Error update data in database", err)
	} else {
		_, err := db.Exec("INSERT INTO service_data (name, data) VALUES (?, ?)", name, data)
		e.CheckError("Error record data to database", err)
	}
}

func GetServiceData(key string) ServiceData {
	db := GetConnection()

	var foundModel ServiceData

	row := db.QueryRow("select name, data from service_data where name = ?", key)
	row.Scan(&foundModel.Name, &foundModel.Data)

	return foundModel
}
