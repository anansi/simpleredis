package protocol

// the datastore
var datastore map[string]string

const OkResponse = "OK"

func init() {

	// initialise a datastore, a map in this case
	datastore = make(map[string]string)
	datastore["greeting"] = "howzit" // from me

}

func set(key, value string) string {

	datastore[key] = value
	return OkResponse
}

func get(key string) string {

	// Note the datastore was initialised with a value in the init function
	value, _ := datastore[key]

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
