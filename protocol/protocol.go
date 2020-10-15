package protocol

// This package houses the datastore and the protocol functions to SET or GET from the datastore, concurrently.

import	(
	"sync"
)

// The SafeMap type with a Mutex.
type SafeMap struct {
	datastore   map[string]string
	mux sync.Mutex
}

// the safe map datastore var
var safeMap SafeMap

const (	
	OkResponse = "OK"
)

func init() {

	// initialise a SafeMap datastore, a map in this case
	datastore := make(map[string]string)
	datastore["greeting"] = "howzit" // from me
	safeMap = SafeMap{datastore: datastore}

}

func set(key, value string) string {
	safeMap.mux.Lock()
	safeMap.datastore[key] = value	
	safeMap.mux.Unlock()
	return OkResponse
}

func get(key string) string {

	// Note the datastore was initialised with a value in the init function

	safeMap.mux.Lock()	
	value, _ := safeMap.datastore[key]
	safeMap.mux.Unlock()
	return value

}

// the public usage function for the Protocol
func ExecuteCmd(cmd string, parameters []string) string {

	switch cmd {
	case "GET":
		key := parameters[0]
		return get(key)
	case "SET":
		key := parameters[0]
		value := parameters[1]
		return set(key, value)
	}

	return OkResponse
}
