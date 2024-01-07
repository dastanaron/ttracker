package GUIController

import (
	"github.com/gotk3/gotk3/gtk"
)

type GUICommonWindow struct {
	MainWindow *gtk.Window
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
