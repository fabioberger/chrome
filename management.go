package chrome

import "github.com/gopherjs/gopherjs/js"

type Management struct {
	o *js.Object
}

/*
* Types
 */

type IconInfo struct {
	*js.Object
	Size int    `js:"size"`
	Url  string `js:"url"`
}

type ExtensionInfo struct {
	*js.Object
	Id                   string     `js:"id"`
	Name                 string     `js:"name"`
	ShortName            string     `js:"shortName"`
	Description          string     `js:"description"`
	Version              string     `js:"version"`
	MayDisable           bool       `js:"mayDisable"`
	Enabled              string     `js:"enabled"`
	DisabledReason       string     `js:"disabledReason"`
	IsApp                bool       `js:"isApp"`
	Type                 string     `js:"type"`
	AppLaunchUrl         string     `js:"appLaunchUrl"`
	HomepageUrl          string     `js:"homepageUrl"`
	UpdateUrl            string     `js:"updateUrl"`
	OfflineEnabled       bool       `js:"offlineEnabled"`
	OptionsUrl           string     `js:"optionsUrl"`
	Icons                []IconInfo `js:"icons"`
	Permissions          []string   `js:"permissions"`
	HostPermissions      []string   `js:"hostPermissions"`
	InstallType          string     `js:"installType"`
	LaunchType           string     `js:"launchType"`
	AvailableLaunchTypes []string   `js:"availableLaunchTypes"`
}

/*
* Methods
 */

// GetAll returns a list of information about installed extensions and apps.
func (m *Management) GetAll(callback func(result []ExtensionInfo)) {
	m.o.Call("getAll", callback)
}

// Get returns information about the installed extension, app, or theme that has the given ID.
func (m *Management) Get(id string, callback func(result ExtensionInfo)) {
	m.o.Call("get", id, callback)
}

// GetSelf returns information about the calling extension, app, or theme. Note:
// This function can be used without requesting the 'management' permission in the manifest.
func (m *Management) GetSelf(callback func(result ExtensionInfo)) {
	m.o.Call("getSelf", callback)
}

// GetPermissionWarningsById returns a list of permission warnings for the given extension id.
func (m *Management) GetPermissionWarningsById(id string, callback func(permissionWarnings []string)) {
	m.o.Call("getPermissionWarningsById", id, callback)
}

// GetPermissionWarningsByManifest returns a list of permission warnings for the given extension manifest
// string. Note: This function can be used without requesting the 'management' permission in the manifest.
func (m *Management) GetPermissionWarningsByManifest(manifestStr string, callback func(permissionWarnings []string)) {
	m.o.Call("getPermissionWarningsByManifest", manifestStr, callback)
}

// SetEnabled enables or disables an app or extension.
func (m *Management) SetEnabled(id, string, enabled bool, callback func()) {
	m.o.Call("setEnabled", id, enabled, callback)
}

// Uninstall uninstalls a currently installed app or extension.
func (m *Management) Uninstall(id string, options Object, callback func()) {
	m.o.Call("uninstall", id, options, callback)
}

// UninstallSelf uninstalls the calling extension. Note: This function can be used without
// requesting the 'management' permission in the manifest.
func (m *Management) UninstallSelf(options Object, callback func()) {
	m.o.Call("uninstallSelf", options, callback)
}

// LaunchApp launches an application.
func (m *Management) LaunchApp(id string, callback func()) {
	m.o.Call("launchApp", id, callback)
}

// CreateAppShortcut display options to create shortcuts for an app. On Mac, only packaged app shortcuts can be created.
func (m *Management) CreateAppShortcut(id string, callback func()) {
	m.o.Call("createAppShortcut", id, callback)
}

// SetLaunchType set the launch type of an app.
func (m *Management) SetLaunchType(id string, launchType string, callback func()) {
	m.o.Call("setLaunchType", id, launchType, callback)
}

// GenerateAppForLink generate an app for a URL. Returns the generated bookmark app.
func (m *Management) GenerateAppForLink(url string, title string, callback func(result ExtensionInfo)) {
	m.o.Call("generateAppForLink", url, title, callback)
}

/*
* Events
 */

// OnInstalled fired when an app or extension has been installed.
func (m *Management) OnInstalled(callback func(info ExtensionInfo)) {
	m.o.Get("onInstalled").Call("addListener", callback)
}

// OnUninstalled fired when an app or extension has been uninstalled.
func (m *Management) OnUninstalled(callback func(id string)) {
	m.o.Get("onUninstalled").Call("addListener", callback)
}

// OnEnabled fired when an app or extension has been enabled.
func (m *Management) OnEnabled(callback func(info ExtensionInfo)) {
	m.o.Get("onEnabled").Call("addListener", callback)
}

// OnDisabled fired when an app or extension has been disabled.
func (m *Management) OnDisabled(callback func(info ExtensionInfo)) {
	m.o.Get("onDisabled").Call("addListener", callback)
}
