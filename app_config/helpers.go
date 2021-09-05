package AppConfig

import "os"

func env(variable string, defaultValue string) (value string) {
	value = os.Getenv(variable)

	if value == "" {
		return defaultValue
	} else {
		return value
	}
}
