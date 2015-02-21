package main

import (
	"fmt"
	"time"

	"github.com/fabioberger/chrome"
	QUnit "github.com/fabioberger/qunit"
	"honnef.co/go/js/dom"
)

var (
	doc = dom.GetWindow().Document()
)

func main() {
	c := chrome.NewChrome()
	/*
	* Alarm Tests
	 */

	QUnit.Module("Alarms")

	// Create New Alarm
	alarmOps := chrome.Object{
		"when": time.Now().UnixNano() + 1000000,
	}
	c.Alarms.Create("test_alarm", alarmOps)

	// Get the Alarm created above
	c.Alarms.Get("test_alarm", func(alarm chrome.Alarm) {
		QUnit.Test("Get()", func(assert QUnit.QUnitAssert) {
			assert.Equal(alarm.Name, "test_alarm", "Get")
		})

		// Clear the Alarm retrieved above
		c.Alarms.Clear("test_alarm", func(wasCleared bool) {
			QUnit.Test("Clear()", func(assert QUnit.QUnitAssert) {
				assert.Equal(wasCleared, true, "Clear")
			})

			// Create two more Alarms
			c.Alarms.Create("test_alarm2", alarmOps)
			c.Alarms.Create("test_alarm3", alarmOps)

			// Get both Alarms created above
			c.Alarms.GetAll(func(alarms []chrome.Alarm) {
				QUnit.Test("GetAll()", func(assert QUnit.QUnitAssert) {
					assert.Equal(alarms[0].Name, "test_alarm2", "GetAll")
					assert.Equal(alarms[1].Name, "test_alarm3", "GetAll")
				})
				fmt.Println("finished running getAll")
				// Clear both Alarms above
				c.Alarms.ClearAll(func(wasCleared bool) {
					QUnit.Test("Clear()", func(assert QUnit.QUnitAssert) {
						assert.Equal(wasCleared, true, "Clear")
					})
				})
			})
		})
	})

}
