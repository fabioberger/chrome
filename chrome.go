package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Chrome struct {
	o       js.Object
	Tabs    Tabs
	Windows Windows
	Runtime Runtime
}

func NewChrome() Chrome {
	c := Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = Tabs{o: js.Global.Get(CHROME).Get("tabs")}
	c.Windows = Windows{o: js.Global.Get(CHROME).Get("windows")}
	c.Runtime = Runtime{o: js.Global.Get(CHROME).Get("runtime")}
	return c
}
