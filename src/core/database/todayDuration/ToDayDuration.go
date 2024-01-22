package todayduration

import (
	"encoding/json"
	"gui-mini-ttracker/core/database"
	"gui-mini-ttracker/helpers"
	"time"
)

type model struct {
	Duration        int
	LatestTimeStamp int64
}

const today_duration_key = "today-duration"

func IncrementDuration(duration int) {
	rawData := database.GetServiceData(today_duration_key)

	var model model

	if rawData.Data != "" {
		json.Unmarshal([]byte(rawData.Data), &model)

		if isToday(model.LatestTimeStamp) {
			model.Duration += duration
		} else {
			model.Duration = duration
		}
	}

	model.LatestTimeStamp = time.Now().Unix()

	json, err := json.Marshal(model)
	helpers.CheckError("Error convert data to Json", err)

	database.SaveServiceData(today_duration_key, string(json))
}

func GetToDayDuration() int {
	rawData := database.GetServiceData(today_duration_key)

	var model model

	if rawData.Data != "" {
		json.Unmarshal([]byte(rawData.Data), &model)
	}

	if model.LatestTimeStamp != 0 && isToday(model.LatestTimeStamp) {
		return model.Duration
	} else {
		json, err := json.Marshal(model)
		helpers.CheckError("Error convert data to Json", err)

		database.SaveServiceData(today_duration_key, string(json))
		return 0
	}
}

func isToday(timestamp int64) bool {
	currentTime := time.Now()
	time := time.Unix(timestamp, 0)

	return time.Day() == currentTime.Day() && time.Month() == currentTime.Month() && time.Year() == currentTime.Year()
}
