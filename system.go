package chrome

import "github.com/gopherjs/gopherjs/js"

type System struct {
	o       js.Object
	Cpu     *Cpu
	Memory  *Memory
	Storage *SysStorage
}

func NewSystem(systemObj js.Object) *System {
	s := new(System)
	s.o = systemObj
	if systemObj.String() != "undefined" {
		s.Cpu = &Cpu{o: systemObj.Get("cpu")}
		s.Memory = &Memory{o: systemObj.Get("memory")}
		s.Storage = &SysStorage{o: systemObj.Get("storage")}
	}
	return s
}

type Cpu struct {
	o js.Object
}

type Memory struct {
	o js.Object
}

type SysStorage struct {
	o js.Object
}

/*
* Types
 */

type StorageUnitInfo struct {
	js.Object
	Id       string `js:"id"`
	Name     string `js:"name"`
	Type     string `js:"type"`
	Capacity int64  `js:"capacity"`
}

/*
* Methods
 */

// GetInfo queries basic CPU information of the system.
func (c *Cpu) GetInfo(callback func(info Object)) {
	c.o.Call("getInfo", callback)
}

// GetInfo get physical memory information.
func (m *Memory) GetInfo(callback func(info map[string]int64)) {
	m.o.Call("getInfo", callback)
}

// GetInfo get the storage information from the system. The argument passed to
// the callback is an array of StorageUnitInfo objects.
func (s *SysStorage) GetInfo(callback func(info []StorageUnitInfo)) {
	s.o.Call("getInfo", callback)
}

// EjectDevice ejects a removable storage device.
func (s *SysStorage) EjectDevice(id string, callback func(result string)) {
	s.o.Call("ejectDevice", id, callback)
}

// GetAvailableCapacity get the available capacity of a specified |id| storage device.
// The |id| is the transient device ID from StorageUnitInfo.
func (s *SysStorage) GetAvailableCapacity(callback func(info Object)) {
	s.o.Call("getAvailableCapacity", callback)
}

/*
* Events
 */

// OnAttached fired when a new removable storage is attached to the system.
func (s *SysStorage) OnAttached(callback func(info StorageUnitInfo)) {
	s.o.Get("onAttached").Call("addListener", callback)
}

// OnDetached fired when a removable storage is detached from the system.
func (s *SysStorage) OnDetached(callback func(id string)) {
	s.o.Get("onDetached").Call("addListener", callback)
}
