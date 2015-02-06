package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Chrome struct {
	o         js.Object
	Tabs      Tabs
	Windows   Windows
	Runtime   Runtime
	Alarms    Alarms
	Bookmarks Bookmarks
}

func NewChrome() Chrome {
	c := Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = Tabs{o: c.o.Get("tabs")}
	c.Windows = Windows{o: c.o.Get("windows")}
	c.Runtime = Runtime{o: c.o.Get("runtime")}
	c.Alarms = Alarms{o: c.o.Get("alarms")}
	c.Bookmarks = Bookmarks{o: c.o.Get("bookmarks")}
	return c
}
