package chrome

import "github.com/gopherjs/gopherjs/js"

type DesktopCapture struct {
	o *js.Object
}

/*
* Methods:
 */

// ChooseDesktopMedia shows desktop media picker UI with the specified set of sources.
func (d *DesktopCapture) ChooseDesktopMedia(sources []string, targetTab Tab, callback func(streamId string)) int {
	return d.o.Call("chooseDesktopMedia", sources, targetTab, callback).Int()
}

// CancelChooseDesktopMedia hides desktop media picker dialog shown by chooseDesktopMedia().
func (d *DesktopCapture) CancelChooseDesktopMedia(desktopMediaRequestId int) {
	d.o.Call("cancelChooseDesktopMedia", desktopMediaRequestId)
}
