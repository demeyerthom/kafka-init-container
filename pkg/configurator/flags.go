package configurator

import "fmt"

func CreateFlag(flag string, value interface{}) string {
	return fmt.Sprintf("--%s %v", flag, value)
}

func CreateConfigFlag(flag string, value interface{}) string {
	return CreateFlag("config", fmt.Sprintf("%s=%v", flag, value))
}
