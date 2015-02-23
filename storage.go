package chrome

import "github.com/gopherjs/gopherjs/js"

type Storage struct {
	o       *js.Object
	Sync    *js.Object
	Local   *js.Object
	Managed *js.Object
}

func NewStorage(storageObj *js.Object) *Storage {
	s := new(Storage)
	s.o = storageObj
	if s.o.String() != "undefined" {
		s.Sync = storageObj.Get("sync")
		s.Local = storageObj.Get("local")
		s.Managed = storageObj.Get("managed")
	}
	return s
}

/*
* Types
 */

type StorageChange struct {
	OldValue interface{} `js:"oldValue"`
	NewValue interface{} `js:"newValue"`
}

/*
* Events
 */

// OnChanged fired when one or more items change.
func (s *Storage) OnChanged(callback func(changes map[string]StorageChange, areaName string)) {
	s.o.Get("onChanged").Call("addListener", callback)
}
