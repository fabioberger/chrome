package chrome

import "github.com/gopherjs/gopherjs/js"

type DesktopCapture struct {
	o js.Object
}

/*
* Methods:
 */

func (d *DesktopCapture) ChooseDesktopMedia(sources []map[string]interface{}, targetTab map[string]interface{}, callback func(streamId string)) int {
	return d.o.Call("chooseDesktopMedia", sources, targetTab, callback).Int()
}

func (d *DesktopCapture) CancelChooseDesktopMedia(desktopMediaRequestId int) {
	d.o.Call("cancelChooseDesktopMedia", desktopMediaRequestId)
}
