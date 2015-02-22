package chrome

import "github.com/gopherjs/gopherjs/js"

type Gcm struct {
	o                js.Object
	MAX_MESSAGE_SIZE int
}

func NewGcm(gcmObj js.Object) *Gcm {
	g := new(Gcm)
	g.o = gcmObj
	if g.o.String() != "undefined" {
		g.MAX_MESSAGE_SIZE = g.o.Get("MAX_MESSAGE_SIZE").Int()
	}
	return g
}

/*
* Methods:
 */

// Register registers the application with GCM. The registration ID will be returned by the callback.
// If register is called again with the same list of senderIds, the same registration ID will be returned.
func (g *Gcm) Register(senderIds []string, callback func(registrationId string)) {
	g.o.Call("register", senderIds, callback)
}

// Unregister unregisters the application from GCM.
func (g *Gcm) Unregister(callback func()) {
	g.o.Call("unregister", callback)
}

// Send sends a message according to its contents.
func (g *Gcm) Send(message Object, callback func(messageId string)) {
	g.o.Call("send", message, callback)
}

/*
* Events
 */

// OnMessage fired when a message is received through GCM.
func (g *Gcm) OnMessage(callback func(message Object)) {
	g.o.Get("onMessage").Call("addListener", callback)
}

// OnMessageDeleted fired when a GCM server had to delete messages sent by an app server to the application.
// See Messages deleted event section of Cloud Messaging documentation for details on handling this event.
func (g *Gcm) OnMessageDeleted(callback func()) {
	g.o.Get("onMessageDeleted").Call("addListener", callback)
}

// OnSendError fired when it was not possible to send a message to the GCM server.
func (g *Gcm) OnSendError(callback func(error Object)) {
	g.o.Get("onSendError").Call("addListener", callback)
}
