package chrome

import "github.com/gopherjs/gopherjs/js"

type Sessions struct {
	o                   js.Object
	MAX_SESSION_RESULTS int
}

func NewSessions(sessionsObj js.Object) *Sessions {
	s := new(Sessions)
	s.o = sessionsObj
	if s.o.String() != "undefined" {
		s.MAX_SESSION_RESULTS = s.o.Get("MAX_SESSION_RESULTS").Int()
	}
	return s
}

/*
* Types
 */

type Filter struct {
	js.Object
	MaxResults int `js:"maxResults"`
}

type Session struct {
	js.Object
	LastModified int    `js:"lastModified"`
	Tab          Tab    `js:"tab"`
	Window       Window `js:"window"`
}

type Device struct {
	js.Object
	DeviceName string    `js:"deviceName"`
	Sessions   []Session `js:"sessions"`
}

/*
* Methods
 */

// GetRecentlyClosed gets the list of recently closed tabs and/or windows.
func (s *Sessions) GetRecentlyClosed(filter Filter, callback func(sessions []Session)) {
	s.o.Call("getRecentlyClosed", filter, callback)
}

// GetDevices retrieves all devices with synced sessions.
func (s *Sessions) GetDevices(filter Filter, callback func(devices []Device)) {
	s.o.Call("getDevices", filter, callback)
}

// Restore reopens a windows.Window or tabs.Tab, with an optional callback to run when the entry has been restored.
func (s *Sessions) Restore(sessionId string, callback func(restoredSession Session)) {
	s.o.Call("restore", sessionId, callback)
}

/*
* Events
 */

// OnChanged fired when recently closed tabs and/or windows are changed. This event does not monitor synced sessions changes.
func (s *Sessions) OnChanged(callback func()) {
	s.o.Get("onChanged").Call("addListener", callback)
}
