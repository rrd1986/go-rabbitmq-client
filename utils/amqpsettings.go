package utils

import (
	"os"
	"strconv"
)

//const AmqpDefaultPort = 5672
//const AmqpDefaultVirtualHost = "2bb143d6-8dfb-4a9f-9afe-aca5a384ab46"
//const AmqpDefaultUsername = "1766b4f3-2124-4f12-afe2-61a462bbaff9"
//const AmqpDefaultPassword = "UcaoJH9N9QxMCJP7AJocttVCK0h8OcNp"

var AmqpDefaultHostname = "localhost"
var AmqpDefaultVirtualHost = ""
var AmqpDefaultPort = 5673
var AmqpDefaultUsername = "guest"
var AmqpDefaultPassword = "guest"
var DefaultDeployConsumer = true
var DefaultExchangeName = "internal.exchange"

// TODO: later support multiple serial no to be registered to this service using an api contract
var OpccDefaultSerialNo = "taxy5nacnJBSuF"

// Overide the defaults from env values
func SetAmqpSettings() {
	val, ok := os.LookupEnv(AmqpEnvHostname)
	if ok {
		AmqpDefaultHostname = val
	}
	val, ok = os.LookupEnv(AmqpEnvVirtualHost)
	if ok {
		AmqpDefaultVirtualHost = val
	}
	val, ok = os.LookupEnv(AmqpEnvPort)
	if ok {
		AmqpDefaultPort, _ = strconv.Atoi(val)
	}
	val, ok = os.LookupEnv(AmqpEnvUsername)
	if ok {
		AmqpDefaultUsername = val
	}
	val, ok = os.LookupEnv(AmqpEnvPassword)
	if ok {
		AmqpDefaultPassword = val
	}
	val, ok = os.LookupEnv(OpccSerialNo)
	if ok {
		OpccDefaultSerialNo = val
	}
	val, ok = os.LookupEnv(AmqpEnvExchangeName)
	if ok {
		DefaultExchangeName = val
	}
}
