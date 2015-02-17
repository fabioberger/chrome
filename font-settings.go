package chrome

import "github.com/gopherjs/gopherjs/js"

type FontSettings struct {
	o js.Object
}

/*
* Types
 */

type FontName struct {
	js.Object
	FontId      string `js:"fontId"`
	DisplayName string `js:"displayName"`
}

/*
* Methods:
 */

// ClearFont clears the font set by this extension, if any.
func (f *FontSettings) ClearFont(details Object, callback func()) {
	f.o.Call("clearFont", details, callback)
}

// GetFont gets the font for a given script and generic font family.
func (f *FontSettings) GetFont(details Object, callback func(details Object)) {
	f.o.Call("getFont", details, callback)
}

// SetFont sets the font for a given script and generic font family.
func (f *FontSettings) SetFont(details Object, callback func()) {
	f.o.Call("setFont", details, callback)
}

// GetFontList gets a list of fonts on the system.
func (f *FontSettings) GetFontList(callback func(results []FontName)) {
	f.o.Call("getFontList", callback)
}

// ClearDefaultFontSize clears the default font size set by this extension, if any.
func (f *FontSettings) ClearDefaultFontSize(details Object, callback func()) {
	f.o.Call("clearDefaultFontSize", details, callback)
}

// GetDefaultFontSize gets the default font size.
func (f *FontSettings) GetDefaultFontSize(details Object, callback func(details Object)) {
	f.o.Call("getDefaultFontSize", details, callback)
}

// SetDefaultFontSize sets the default font size.
func (f *FontSettings) SetDefaultFontSize(details Object, callback func()) {
	f.o.Call("setDefaultFontSize", details, callback)
}

// ClearDefaultFixedFontSize clears the default fixed font size set by this extension, if any.
func (f *FontSettings) ClearDefaultFixedFontSize(details Object, callback func()) {
	f.o.Call("clearDefaultFixedFontSize", details, callback)
}

// GetDefaultFixedFontSize gets the default size for fixed width fonts.
func (f *FontSettings) GetDefaultFixedFontSize(details Object, callback func(details Object)) {
	f.o.Call("getDefaultFixedFontSize", details, callback)
}

// SetDefaultFixedFontSize sets the default size for fixed width fonts.
func (f *FontSettings) SetDefaultFixedFontSize(details Object, callback func()) {
	f.o.Call("setDefaultFixedFontSize", details, callback)
}

// ClearMinimumFontSize lears the minimum font size set by this extension, if any.
func (f *FontSettings) ClearMinimumFontSize(details Object, callback func()) {
	f.o.Call("clearMinimumFontSize", details, callback)
}

// GetMinimumFontSize gets the minimum font size.
func (f *FontSettings) GetMinimumFontSize(details Object, callback func(details Object)) {
	f.o.Call("getMinimumFontSize", details, callback)
}

// SetMinimumFontSize sets the minimum font size.
func (f *FontSettings) SetMinimumFontSize(details Object, callback func()) {
	f.o.Call("setMinimumFontSize", details, callback)
}

/*
* Events
 */

// OnFontChanged fired when a font setting changes.
func (f *FontSettings) OnFontChanged(callback func(details Object)) {
	f.o.Get("onFontChanged").Call("addListener", callback)
}

// OnDefaultFontSizeChanged fired when the default font size setting changes.
func (f *FontSettings) OnDefaultFontSizeChanged(callback func(details Object)) {
	f.o.Get("onDefaultFontSizeChanged").Call("addListener", callback)
}

// OnDefaultFixedFontSizeChanged fired when the default fixed font size setting changes.
func (f *FontSettings) OnDefaultFixedFontSizeChanged(callback func(details Object)) {
	f.o.Get("onDefaultFixedFontSizeChanged").Call("addListener", callback)
}

// OnMinimumFontSizeChanged fired when the minimum font size setting changes.
func (f *FontSettings) OnMinimumFontSizeChanged(callback func(details Object)) {
	f.o.Get("onMinimumFontSizeChanged").Call("addListener", callback)
}
