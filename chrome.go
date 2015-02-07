package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Chrome struct {
	o                  js.Object
	Tabs               Tabs
	Windows            Windows
	Runtime            Runtime
	Alarms             Alarms
	Bookmarks          Bookmarks
	BrowserAction      BrowserAction
	BrowsingData       BrowsingData
	Commands           Commands
	ContextMenus       ContextMenus
	Cookies            Cookies
	Debugger           Debugger
	DeclarativeContent DeclarativeContent
	DesktopCapture     DesktopCapture
	Downloads          Downloads
	Enterprise         Enterprise
}

func NewChrome() Chrome {
	c := Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = Tabs{o: c.o.Get("tabs")}
	c.Windows = Windows{o: c.o.Get("windows")}
	c.Runtime = Runtime{o: c.o.Get("runtime")}
	c.Alarms = Alarms{o: c.o.Get("alarms")}
	c.Bookmarks = Bookmarks{o: c.o.Get("bookmarks")}
	c.BrowserAction = BrowserAction{o: c.o.Get("browserAction")}
	c.BrowsingData = BrowsingData{o: c.o.Get("browsingData")}
	c.Commands = Commands{o: c.o.Get("commands")}
	c.ContextMenus = ContextMenus{o: c.o.Get("contextMenus")}
	c.Cookies = Cookies{o: c.o.Get("cookies")}
	c.Debugger = Debugger{o: c.o.Get("debugger")}
	c.DeclarativeContent = DeclarativeContent{
		o:             c.o.Get("declarativeContent"),
		OnPageChanged: OnPageChanged{o: c.o.Get("declarativeContent").Get("onPageChanged")},
	}
	c.DesktopCapture = DesktopCapture{o: c.o.Get("desktopCapture")}
	c.Downloads = Downloads{o: c.o.Get("downloads")}
	c.Enterprise = Enterprise{
		o:            c.o.Get("enterprise"),
		PlatformKeys: PlatformKeys{o: c.o.Get("enterprise").Get("platformKeys")},
	}
	return c
}
