package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Chrome struct {
	o                  js.Object
	Tabs               *Tabs
	Windows            *Windows
	Runtime            *Runtime
	Alarms             *Alarms
	Bookmarks          *Bookmarks
	BrowserAction      *BrowserAction
	BrowsingData       *BrowsingData
	Commands           *Commands
	ContextMenus       *ContextMenus
	Cookies            *Cookies
	Debugger           *Debugger
	DeclarativeContent *DeclarativeContent
	DesktopCapture     *DesktopCapture
	Downloads          *Downloads
	Enterprise         *Enterprise
	Extension          *Extension
	FileBrowserHandler *FileBrowserHandler
	FileSystemProvider *FileSystemProvider
	FontSettings       *FontSettings
	Gcm                *Gcm
	History            *History
	I18n               *I18n
	Identity           *Identity
	Idle               *Idle
	Input              *Input
	Notification       *Notification
	Omnibox            *Omnibox
	PageAction         *PageAction
	PageCapture        *PageCapture
	Permissions        *Permissions
	Power              *Power
	Privacy            *Privacy
}

func NewChrome() *Chrome {
	c := &Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = &Tabs{o: c.o.Get("tabs")}
	c.Windows = &Windows{o: c.o.Get("windows")}
	c.Runtime = &Runtime{o: c.o.Get("runtime")}
	c.Alarms = &Alarms{o: c.o.Get("alarms")}
	c.Bookmarks = &Bookmarks{o: c.o.Get("bookmarks")}
	c.BrowserAction = &BrowserAction{o: c.o.Get("browserAction")}
	c.BrowsingData = &BrowsingData{o: c.o.Get("browsingData")}
	c.Commands = &Commands{o: c.o.Get("commands")}
	c.ContextMenus = &ContextMenus{o: c.o.Get("contextMenus")}
	c.Cookies = &Cookies{o: c.o.Get("cookies")}
	c.Debugger = &Debugger{o: c.o.Get("debugger")}
	c.DeclarativeContent = NewDeclarativeContent(c.o.Get("declarativeContent"))
	c.DesktopCapture = &DesktopCapture{o: c.o.Get("desktopCapture")}
	c.Downloads = &Downloads{o: c.o.Get("downloads")}
	c.Enterprise = NewEnterprise(c.o.Get("enterprise"))
	c.Extension = &Extension{o: c.o.Get("extension")}
	c.FileBrowserHandler = &FileBrowserHandler{o: c.o.Get("fileBrowserHandler")}
	c.FileSystemProvider = &FileSystemProvider{o: c.o.Get("fileSystemProvider")}
	c.FontSettings = &FontSettings{o: c.o.Get("fontSettings")}
	c.Gcm = &Gcm{o: c.o.Get("gcm")}
	c.History = &History{o: c.o.Get("history")}
	c.I18n = &I18n{o: c.o.Get("i18n")}
	c.Identity = &Identity{o: c.o.Get("identity")}
	c.Idle = &Idle{o: c.o.Get("idle")}
	c.Input = NewInput(c.o.Get("input"))
	c.Notification = &Notification{o: c.o.Get("notification")}
	c.Omnibox = &Omnibox{o: c.o.Get("omnibox")}
	c.PageAction = &PageAction{o: c.o.Get("pageAction")}
	c.PageCapture = &PageCapture{o: c.o.Get("pageCapture")}
	c.Permissions = &Permissions{o: c.o.Get("permissions")}
	c.Power = &Power{o: c.o.Get("power")}
	c.Privacy = NewPrivacy(c.o.Get("privacy"))
	return c
}
