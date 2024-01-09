package GUIController

import (
	"github.com/gotk3/gotk3/gtk"
)

type GUICommonWindow struct {
	MainWindow       *gtk.Window
	StartButton      *gtk.Button
	StartButtonImage *gtk.Image
	TaskNameField    *gtk.Entry
	Timer            *gtk.Label
	TasksGrid        *gtk.Grid
	Total            *gtk.Label
}

type GUIErrorDialog struct {
	ErrorDialog *gtk.MessageDialog
	CloseButton *gtk.Button
}

type GUIInterface struct {
	Common      GUICommonWindow
	ErrorDialog GUIErrorDialog
}

func NewGUIController(gtkBuilder *gtk.Builder) (*GUIInterface, error) {

	commonWindow := GUICommonWindow{}
	errorDialog := GUIErrorDialog{}

	obj, err := gtkBuilder.GetObject("window_main")
	if err != nil {
		return nil, err
	}

	commonWindow.MainWindow = obj.(*gtk.Window)

	obj, err = gtkBuilder.GetObject("start_button")
	if err != nil {
		return nil, err
	}

	commonWindow.StartButton = obj.(*gtk.Button)

	obj, err = gtkBuilder.GetObject("start_button_image")
	if err != nil {
		return nil, err
	}

	commonWindow.StartButtonImage = obj.(*gtk.Image)

	obj, err = gtkBuilder.GetObject("task_name_field")
	if err != nil {
		return nil, err
	}

	commonWindow.TaskNameField = obj.(*gtk.Entry)

	obj, err = gtkBuilder.GetObject("current_time_label")
	if err != nil {
		return nil, err
	}
	commonWindow.Timer = obj.(*gtk.Label)

	obj, err = gtkBuilder.GetObject("tasks_grid")
	if err != nil {
		return nil, err
	}
	commonWindow.TasksGrid = obj.(*gtk.Grid)

	obj, err = gtkBuilder.GetObject("total_time")
	if err != nil {
		return nil, err
	}
	commonWindow.Total = obj.(*gtk.Label)

	obj, err = gtkBuilder.GetObject("error_dialog")
	if err != nil {
		return nil, err
	}
	errorDialog.ErrorDialog = obj.(*gtk.MessageDialog)

	obj, err = gtkBuilder.GetObject("error_dialog_close")
	if err != nil {
		return nil, err
	}
	errorDialog.CloseButton = obj.(*gtk.Button)

	guiInterface := &GUIInterface{
		Common:      commonWindow,
		ErrorDialog: errorDialog,
	}

	return guiInterface, nil
}
