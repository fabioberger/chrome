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

// AddRules takes an array of rule instances as its first parameter and a callback function that is called on completion
func (e *OnPageChanged) AddRules(rules []map[string]interface{}, callback func(details map[string]interface{})) {
	e.o.Call("addRules", rules, callback)
}

// RemoveRules accepts an optional array of rule identifiers as its first parameter and a callback function as its second parameter.
func (e *OnPageChanged) RemoveRules(ruleIdentifiers []string, callback func()) {
	e.o.Call("removeRules", ruleIdentifiers, callback)
}

// GetRules accepts an optional array of rule identifiers with the same semantics as removeRules and a callback function.
func (e *OnPageChanged) GetRules(ruleIdentifiers []string, callback func(details map[string]interface{})) {
	e.o.Call("getRules", ruleIdentifiers, callback)
}
