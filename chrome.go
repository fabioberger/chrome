package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Object map[string]interface{}

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
	Proxy              *Proxy
	Sessions           *Sessions
	Storage            *Storage
	System             *System
	TabCapture         *TabCapture
	TopSites           *TopSites
	Tts                *Tts
	TtsEngine          *TtsEngine
	WebNavigation      *WebNavigation
	WebRequest         *WebRequest
	WebStore           *WebStore
}

func NewChrome() *Chrome {
	c := &Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = &Tabs{o: c.o.Get("tabs")}
	c.Runtime = NewRuntime(c.o.Get("runtime"))
	c.Alarms = &Alarms{o: c.o.Get("alarms")}
	c.Bookmarks = &Bookmarks{o: c.o.Get("bookmarks")}
	c.BrowserAction = &BrowserAction{o: c.o.Get("browserAction")}
	c.BrowsingData = &BrowsingData{o: c.o.Get("browsingData")}
	c.Commands = &Commands{o: c.o.Get("commands")}
	c.ContextMenus = NewContextMenus(c.o.Get("contextMenus"))
	c.Cookies = &Cookies{o: c.o.Get("cookies")}
	c.Debugger = &Debugger{o: c.o.Get("debugger")}
	c.DeclarativeContent = NewDeclarativeContent(c.o.Get("declarativeContent"))
	c.DesktopCapture = &DesktopCapture{o: c.o.Get("desktopCapture")}
	c.Downloads = &Downloads{o: c.o.Get("downloads")}
	c.Enterprise = NewEnterprise(c.o.Get("enterprise"))
	c.Extension = NewExtension(c.o.Get("extension"))
	c.FileBrowserHandler = &FileBrowserHandler{o: c.o.Get("fileBrowserHandler")}
	c.FileSystemProvider = &FileSystemProvider{o: c.o.Get("fileSystemProvider")}
	c.FontSettings = &FontSettings{o: c.o.Get("fontSettings")}
	c.Gcm = NewGcm(c.o.Get("gcm"))
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
	c.Proxy = NewProxy(c.o.Get("proxy"))
	c.Sessions = NewSessions(c.o.Get("sessions"))
	c.Storage = NewStorage(c.o.Get("storage"))
	c.System = NewSystem(c.o.Get("system"))
	c.TabCapture = &TabCapture{o: c.o.Get("tabCapture")}
	c.TopSites = &TopSites{o: c.o.Get("topSites")}
	c.Tts = &Tts{o: c.o.Get("tts")}
	c.TtsEngine = &TtsEngine{o: c.o.Get("ttsEngine")}
	c.WebNavigation = &WebNavigation{o: c.o.Get("webNavigation")}
	c.WebRequest = NewWebRequest(c.o.Get("webRequest"))
	c.WebStore = &WebStore{o: c.o.Get("webstore")}
	c.Windows = NewWindows(c.o.Get("windows"))
	return c
}
