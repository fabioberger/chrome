package chrome

import "github.com/gopherjs/gopherjs/js"

type Notification struct {
	o *js.Object
}

/*
* Types
 */

type NotificationOptions struct {
	*js.Object
	Type           string   `js:"type"`
	IconUrl        string   `js:"iconUrl"`
	AppIconMaskUrl string   `js:"appIconMaskUrl"`
	Title          string   `js:"title"`
	Message        string   `js:"message"`
	ContextMessage string   `js:"contextMessage"`
	Priority       int      `js:"priority"`
	EventTime      int64    `js:"eventTime"`
	Buttons        []Object `js:"buttons"`
	ImageUrl       string   `js:"imageUrl"`
	Items          []Object `js:"items"`
	Progress       int      `js:"progress"`
	IsClickable    bool     `js:"isClickable"`
}

/*
* Methods
 */

// Create creates and displays a notification.
func (n *Notification) Create(notificationId string, options NotificationOptions, callback func(notificationId string)) {
	n.o.Call("create", notificationId, options, callback)
}

// Update updates an existing notification.
func (n *Notification) Update(notificationId string, options NotificationOptions, callback func(wasUpdated bool)) {
	n.o.Call("update", notificationId, options, callback)
}

// Clear clears the specified notification.
func (n *Notification) Clear(notificationId string, callback func(notificationId string)) {
	n.o.Call("clear", notificationId, callback)
}

// GetAll retrieves all the notifications.
func (n *Notification) GetAll(callback func(notifications Object)) {
	n.o.Call("getAll", callback)
}

// GetPermissionLevel retrieves whether the user has enabled notifications from this app or extension.
func (n *Notification) GetPermissionLevel(callback func(level string)) {
	n.o.Call("getPermissionLevel", callback)
}

/*
* Events
 */

// OnClosed the notification closed, either by the system or by user action.
func (n *Notification) OnClosed(callback func(notificationId string, byUser bool)) {
	n.o.Get("onClosed").Call("addListener", callback)
}

// OnClicked the user clicked in a non-button area of the notification.
func (n *Notification) OnClicked(callback func(notificationId string)) {
	n.o.Get("onClicked").Call("addListener", callback)
}

// OnButtonClicked the user pressed a button in the notification.
func (n *Notification) OnButtonClicked(callback func(notificationId string, buttonIndex int)) {
	n.o.Get("onButtonClicked").Call("addListener", callback)
}

// OnPermissionLevelChanged the user changes the permission level.
func (n *Notification) OnPermissionLevelChanged(callback func(level string)) {
	n.o.Get("onPermissionLevelChanged").Call("addListener", callback)
}

// OnShowSettings the user clicked on a link for the app's notification settings.
func (n *Notification) OnShowSettings(callback func()) {
	n.o.Get("onShowSettings").Call("addListener", callback)
}
