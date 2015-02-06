package chrome

import "github.com/gopherjs/gopherjs/js"

type Alarms struct {
	o js.Object
}

/*
* Methods:
 */

// You must use time.Now().UnixNano() for "when" timestamp in alarmInfo
func (a *Alarms) Create(name string, alarmInfo interface{}) {
	a.o.Call("create", name, alarmInfo)
}

func (a *Alarms) Get(name string, callback func(alarm map[string]interface{})) {
	a.o.Call("get", name, callback)
}

func (a *Alarms) GetAll(callback func(alarms []map[string]interface{})) {
	a.o.Call("getAll", callback)
}

func (a *Alarms) Clear(name string, callback func(wasCleared bool)) {
	a.o.Call("clear", name, callback)
}

func (a *Alarms) ClearAll(callback func(wasCleared bool)) {
	a.o.Call("clearAll", callback)
}

/*
* Events
 */

func (a *Alarms) OnAlarm(callback func(map[string]interface{})) {
	a.o.Get("onAlarm").Call("addListener", callback)
}
