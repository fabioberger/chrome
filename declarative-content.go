package chrome

import "github.com/gopherjs/gopherjs/js"

type DeclarativeContent struct {
	o             js.Object
	OnPageChanged OnPageChanged
}

type OnPageChanged struct {
	o js.Object
}

/*
* Events
 */

func (e *OnPageChanged) AddRules(rules []map[string]interface{}, callback func(details map[string]interface{})) {
	e.o.Call("addRules", rules, callback)
}

func (e *OnPageChanged) RemoveRules(ruleIdentifiers []string, callback func()) {
	e.o.Call("removeRules", ruleIdentifiers, callback)
}

func (e *OnPageChanged) GetRules(ruleIdentifiers []string, callback func(details map[string]interface{})) {
	e.o.Call("getRules", ruleIdentifiers, callback)
}
