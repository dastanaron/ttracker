package main

import (
	_ "embed"
	"errors"
	"gui-mini-ttracker/core/database"
	guiController "gui-mini-ttracker/gui-controller"
	"gui-mini-ttracker/helpers"
	timer "gui-mini-ttracker/library/timer"
	"log"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

//go:embed main.glade
var gladeInterface string

//go:embed icon.png
var icon []byte

var isRunnedTimer bool

type Task struct {
	StartTime time.Time
	Duration  int
	Name      string
}

func main() {
	gtk.Init(nil)

	pixBuf, err := gdk.PixbufNewFromBytesOnly(icon)
	helpers.CheckError("Error", err)

	gui, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = gui.AddFromString(gladeInterface)
	helpers.CheckError("Error", err)

	controllerGUI, err := guiController.NewGUIController(gui)
	helpers.CheckError("Error", err)

	controllerGUI.Common.MainWindow.SetTitle("MINI tTracker")
	controllerGUI.Common.MainWindow.SetIcon(pixBuf)
	controllerGUI.Common.MainWindow.ToWindow().SetIcon(pixBuf)
	controllerGUI.Common.MainWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})

	controllerGUI.Common.Timer.SetText("00:00:00")

	controllerGUI.ErrorDialog.CloseButton.Connect("clicked", func() {
		controllerGUI.ErrorDialog.ErrorDialog.Hide()
	})

	controllerGUI.Common.MainWindow.ShowAll()

	t := timer.New()

	isRunnedTimer = false
	timerStop := make(chan int)
	currentRow := checkRecentData(*controllerGUI)

	controllerGUI.Common.StartButton.Connect("clicked", func() {
		text, err := controllerGUI.Common.TaskNameField.GetText()

		helpers.CheckGUIError(controllerGUI, "Error receiving text from field", err)

		if len(text) < 3 {
			helpers.CheckGUIError(controllerGUI, "Validation error", errors.New("Task name field must contains more then 2 letters"))
			return
		}

		if !isRunnedTimer {
			t.Start()
			go handleTimer(&t, controllerGUI, timerStop)
			controllerGUI.Common.StartButtonImage.SetFromIconName("media-playback-stop", gtk.ICON_SIZE_BUTTON)
		} else {
			t.Stop()
			timerStop <- -1
			controllerGUI.Common.StartButtonImage.SetFromIconName("media-playback-start", gtk.ICON_SIZE_BUTTON)

			addTaskToGrid(*controllerGUI, text, t.Duration, currentRow)

			database.Save(database.TaskModel{
				Name:      text,
				StartTime: t.StartTime,
				Duration:  t.Duration,
				Project:   "",
			})

			t.Clear()
			currentRow++
		}
		isRunnedTimer = !isRunnedTimer
	})

	gtk.Main()
}

func handleTimer(t *timer.Timer, gui *guiController.GUIInterface, done chan int) {
	for {
		select {
		case <-done:
			return
		default:
			duration := <-t.Tick
			t.Duration = duration
			gui.Common.Timer.SetText(helpers.ConvertSecondsToHumanTime(duration))
		}
	}
}

func addTaskToGrid(gui guiController.GUIInterface, name string, duration int, row int) {
	taskName, _ := gtk.LabelNew(name)

	taskDuration, _ := gtk.LabelNew(helpers.ConvertSecondsToHumanTime(duration))
	taskDuration.SetHAlign(gtk.ALIGN_CENTER)

	button, _ := gtk.ButtonNewFromIconName("gtk-media-play", gtk.ICON_SIZE_BUTTON)

	gui.Common.TasksGrid.Attach(taskName.ToWidget(), 0, row, 1, 1)
	gui.Common.TasksGrid.Attach(taskDuration.ToWidget(), 1, row, 1, 1)
	gui.Common.TasksGrid.Attach(button.ToWidget(), 3, row, 1, 1)
	gui.Common.MainWindow.ShowAll()

	button.Connect("clicked", func() {
		if isRunnedTimer {
			gui.Common.StartButton.Clicked()
		}
		gui.Common.TaskNameField.SetText(name)
		gui.Common.StartButton.Clicked()
	})
}

func checkRecentData(gui guiController.GUIInterface) int {
	currentRow := 0

	models := database.GetToDay()

	for _, model := range models {
		addTaskToGrid(gui, model.Name, model.Duration, currentRow)
		currentRow++
	}

	return currentRow
}
