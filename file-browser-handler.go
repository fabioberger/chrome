// WARNING: FileBrowserHandler only works on Chrome OS

package chrome

import "github.com/gopherjs/gopherjs/js"

type FileBrowserHandler struct {
	o js.Object
}

/*
* Types
 */

type FileHandlerExecuteEventDetails struct {
	Entries []interface{} `js:"entries"`
	Tab_id  int           `js:"tab_id"`
}

/*
* Methods:
 */

/* SelectFile prompts user to select file path under which file should be saved. When the file is selected, file access permission required to use the file (read, write and create) are granted to the caller. The file will not actually get created during the function call, so function caller must ensure its existence before using it. The function has to be invoked with a user gesture. */
func (f *FileBrowserHandler) SelectFile(selectionParams map[string]interface{}, callback func(result map[string]interface{})) {
	f.o.Call("selectFile", selectionParams, callback)
}

/*
* Events
 */

// OnExecute fired when file system action is executed from ChromeOS file browser.
func (f *FileBrowserHandler) OnExecute(callback func(id string, details FileHandlerExecuteEventDetails)) {
	f.o.Get("onExecute").Call("addListener", callback)
}
