package main

import (
	_ "embed"
	"errors"
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

	isStart := true
	timerStop := make(chan int)
	currentRow := 0

	controllerGUI.Common.StartButton.Connect("clicked", func() {
		text, err := controllerGUI.Common.TaskNameField.GetText()

		helpers.CheckGUIError(controllerGUI, "Error receiving text from field", err)

		if len(text) < 3 {
			helpers.CheckGUIError(controllerGUI, "Validation error", errors.New("Task name field must contains more then 2 letters"))
			return
		}

		if isStart {
			t.Start()
			go handleTimer(&t, controllerGUI, timerStop)
			controllerGUI.Common.StartButtonImage.SetFromIconName("media-playback-stop", gtk.ICON_SIZE_BUTTON)
		} else {
			t.Stop()
			timerStop <- -1
			controllerGUI.Common.StartButtonImage.SetFromIconName("media-playback-start", gtk.ICON_SIZE_BUTTON)

			taskName, _ := gtk.LabelNew(text)

			taskDuration, _ := gtk.LabelNew(helpers.ConvertSecondsToHumanTime(t.Duration))
			taskDuration.SetHAlign(gtk.ALIGN_CENTER)

			button, _ := gtk.ButtonNewFromIconName("gtk-delete", gtk.ICON_SIZE_BUTTON)

			controllerGUI.Common.TasksGrid.Attach(taskName.ToWidget(), 0, currentRow, 1, 1)
			controllerGUI.Common.TasksGrid.Attach(taskDuration.ToWidget(), 1, currentRow, 1, 1)
			controllerGUI.Common.TasksGrid.Attach(button.ToWidget(), 3, currentRow, 1, 1)

			controllerGUI.Common.MainWindow.ShowAll()

			t.Clear()
			currentRow++
		}
		isStart = !isStart
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
