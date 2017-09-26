package config

var IS_LOCAL = false
var IS_STAGING = false
var MODE = ""

func SetMode(mode string) {
	MODE = mode
	if mode == "local" {
		IS_LOCAL = true
	}
	if mode == "staging" {
		IS_STAGING = true
	}
}
func HOST() string {
	if MODE == "local" {
		return "http://localhost:8088"
	}

	if MODE == "staging" {
		return "https://openthaiand.org"
	}

	if MODE == "production" {
		return "https://openthaiand.org"
	}
	//if mode is not set it's always in production
	return "https://openthaiand.org"
}
