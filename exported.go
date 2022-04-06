package shutters

// shutterManager provides a package global resource release manager.
var shutterManager = &ShutterManager{}

// Add binds in a new function to the shutter manager, which will be called on
// service shutdown to release resources.
func Add(s Shutter) {
	shutterManager.Add(s)
}

// Close calls all resource releasing functions stored in the shutter manager.
func Close() {
	shutterManager.Close()
}
