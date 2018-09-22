package pkg

import "fmt"

func CreateFlag(flag string) string {
	return fmt.Sprintf("--%s", flag)
}

func CreateConfig(flag string, value string) string {
	return fmt.Sprintf("%s=%s", flag, value)
}
