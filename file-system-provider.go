// WARNING: This API only works on Chrome OS

package chrome

import "github.com/gopherjs/gopherjs/js"

type FileSystemProvider struct {
	o js.Object
}

/*
* Methods:
 */

/* Mount mounts a file system with the given fileSystemId and displayName. displayName will be shown in the left panel of Files.app. displayName can contain any characters including '/', but cannot be an empty string. displayName must be descriptive but doesn't have to be unique. The fileSystemId must not be an empty string. */
func (f *FileSystemProvider) Mount(options map[string]interface{}, callback func()) {
	f.o.Call("mount", options, callback)
}

// Unmount unmounts a file system with the given fileSystemId. It must be called after onUnmountRequested is invoked.
// Also, the providing extension can decide to perform unmounting if not requested (eg. in case of lost connection, or a file error).
func (f *FileSystemProvider) Unmount(options map[string]interface{}, callback func()) {
	f.o.Call("unmount", options, callback)
}

// GetAll returns all file systems mounted by the extension.
func (f *FileSystemProvider) GetAll(callback func(fileSystems []map[string]interface{})) {
	f.o.Call("getAll", callback)
}

// Get returns information about a file system with the passed fileSystemId.
func (f *FileSystemProvider) Get(fileSystemId string, callback func(fileSystemInfo map[string]interface{})) {
	f.o.Call("get", fileSystemId, callback)
}

/*
* Events
 */

// OnUnmountRequested raised when unmounting for the file system with the fileSystemId identifier is requested.
// In the response, the unmount API method must be called together with successCallback .
// If unmounting is not possible (eg. due to a pending operation), then errorCallback must be called.
func (f *FileSystemProvider) OnUnmountRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onUnmountRequested").Call("addListener", callback)
}

// OnGetMetadataRequested raised when metadata of a file or a directory at entryPath is requested.
// The metadata must be returned with the successCallback call. In case of an error, errorCallback must be called.
func (f *FileSystemProvider) OnGetMetadataRequested(callback func(options map[string]interface{}, successCallback func(metadata map[string]interface{}), errorCallback func(err string))) {
	f.o.Get("onGetMetadataRequested").Call("addListener", callback)
}

// OnReadDirectoryRequested raised when contents of a directory at directoryPath are requested.
// The results must be returned in chunks by calling the successCallback several times. In case of an error, errorCallback must be called.
func (f *FileSystemProvider) OnReadDirectoryRequested(callback func(options map[string]interface{}, successCallback func(entries []map[string]interface{}, hasMore bool), errorCallback func(err string))) {
	f.o.Get("onReadDirectoryRequested").Call("addListener", callback)
}

// OnOpenFileRequested raised when opening a file at filePath is requested. If the file does not exist,
// then the operation must fail. Maximum number of files opened at once can be specified with MountOptions.
func (f *FileSystemProvider) OnOpenFileRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onOpenFileRequested").Call("addListener", callback)
}

// OnCloseFileRequested raised when opening a file previously opened with openRequestId is requested to be closed.
func (f *FileSystemProvider) OnCloseFileRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onCloseFileRequested").Call("addListener", callback)
}

// OnReadFileRequested raised when reading contents of a file opened previously with openRequestId is requested.
// The results must be returned in chunks by calling successCallback several times. In case of an error, errorCallback must be called.
func (f *FileSystemProvider) OnReadFileRequested(callback func(options map[string]interface{}, successCallback func(data interface{}, hasMore bool), errorCallback func(err string))) {
	f.o.Get("onReadFileRequested").Call("addListener", callback)
}

// OnCreateDirectoryRequested raised when creating a directory is requested. The operation must fail with the EXISTS error
// if the target directory already exists. If recursive is true, then all of the missing directories on the directory path must be created.
func (f *FileSystemProvider) OnCreateDirectoryRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onCreateDirectoryRequested").Call("addListener", callback)
}

// OnDeleteEntryRequested raised when deleting an entry is requested. If recursive is true, and the entry is a directory,
// then all of the entries inside must be recursively deleted as well.
func (f *FileSystemProvider) OnDeleteEntryRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onDeleteEntryRequested").Call("addListener", callback)
}

// OnCreateFileReqested raised when creating a file is requested. If the file already exists, then errorCallback must
// be called with the EXISTS error code.
func (f *FileSystemProvider) OnCreateFileReqested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onCreateFileReqested").Call("addListener", callback)
}

// OnCopyEntryRequested raised when copying an entry (recursively if a directory) is requested. If an error occurs,
// then errorCallback must be called.
func (f *FileSystemProvider) OnCopyEntryRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onCopyEntryRequested").Call("addListener", callback)
}

// OnMoveEntryRequested raised when moving an entry (recursively if a directory) is requested. If an error occurs,
// then errorCallback must be called.
func (f *FileSystemProvider) OnMoveEntryRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onMoveEntryRequested").Call("addListener", callback)
}

// OnTruncateRequested raised when truncating a file to a desired length is requested. If an error occurs,
// then errorCallback must be called.
func (f *FileSystemProvider) OnTruncateRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onTruncateRequested").Call("addListener", callback)
}

// OnWriteFileRequested raised when writing contents to a file opened previously with openRequestId is requested.
func (f *FileSystemProvider) OnWriteFileRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onWriteFileRequested").Call("addListener", callback)
}

/* OnAbortRequested raised when aborting an operation with operationRequestId is requested. The operation executed with operationRequestId must be immediately stopped and successCallback of this abort request executed. If aborting fails, then errorCallback must be called. Note, that callbacks of the aborted operation must not be called, as they will be ignored. Despite calling errorCallback, the request may be forcibly aborted. */
func (f *FileSystemProvider) OnAbortRequested(callback func(options map[string]interface{}, successCallback func(), errorCallback func(err string))) {
	f.o.Get("onAbortRequested").Call("addListener", callback)
}
