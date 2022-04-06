package shutters

import (
	// Standard Library Imports
	"sync"
)

// Shutter is the sole type that needs to be implemented in order to
// stack functions to be called for performing graceful shutdowns.
type Shutter func()

// NewManager returns a pointer to a new instance of a shutter manager for
// advanced use cases.
func NewManager() *ShutterManager {
	return &ShutterManager{}
}

// ShutterManager provides a manager to enable safe release of provisioned
// resources.
type ShutterManager struct {
	mutex    sync.RWMutex
	shutters []Shutter
}

// Add binds in a new function to be called on service shutdown to release
// resources.
func (m *ShutterManager) Add(s Shutter) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.shutters = append(m.shutters, s)
}

// Close calls all stored resource releasing functions.
func (m *ShutterManager) Close() {
	m.mutex.RLock()
	for _, closer := range m.shutters {
		// Close all the things!
		closer()
	}
	m.mutex.RUnlock()

	// Clear the closer array.
	m.mutex.Lock()
	m.shutters = nil
	m.mutex.Unlock()
}
