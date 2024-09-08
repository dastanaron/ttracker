package main

import (
	"context"
	"ttracker/core/database"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Log(value string) {
	runtime.LogInfo(a.ctx, value)
}

func (a *App) SaveTask(model database.TaskModel) {
	database.Save(model)
	database.IncrementDuration(model.Duration)
}

func (a *App) GetTodayTask() []database.TaskModel {
	return database.GetToDay()
}

func (a *App) GetTodayDuration() int {
	return database.GetToDayDuration()
}

func (a *App) GetLatest(count int) []database.TaskModel {
	return database.GetLatest(count)
}
