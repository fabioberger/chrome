package chrome

import "github.com/gopherjs/gopherjs/js"

/*
* Types
 */
type Alarm struct {
	js.Object
	Name            string `js:"name"`
	ScheduledTime   string `js:"scheduledTime"`
	PeriodInMinutes string `js:"periodInMinutes"`
}

type Alarms struct {
	o js.Object
}

/*
* Methods:
 */

// Create creates an alarm. Near the time(s) specified by alarmInfo, the onAlarm event is fired.
// If there is another alarm with the same name (or no name if none is specified), it will be
// cancelled and replaced by this alarm.
// You must use time.Now().UnixNano() for "when" timestamp in alarmInfo for this to work.
func (a *Alarms) Create(name string, alarmInfo interface{}) {
	a.o.Call("create", name, alarmInfo)
}

// Get retrieves details about the specified alarm.
func (a *Alarms) Get(name string, callback func(alarm Alarm)) {
	a.o.Call("get", name, callback)
}

// GetAll gets an array of all the alarms.
func (a *Alarms) GetAll(callback func(alarms []Alarm)) {
	a.o.Call("getAll", callback)
}

// Clear clears the alarm with the given name.
func (a *Alarms) Clear(name string, callback func(wasCleared bool)) {
	a.o.Call("clear", name, callback)
}

// ClearAll clears all alarms.
func (a *Alarms) ClearAll(callback func(wasCleared bool)) {
	a.o.Call("clearAll", callback)
}

/*
* Events
 */
// OnAlarm is fired when an alarm has elapsed. Useful for event pages.
func (a *Alarms) OnAlarm(callback func(Alarm)) {
	a.o.Get("onAlarm").Call("addListener", callback)
}
