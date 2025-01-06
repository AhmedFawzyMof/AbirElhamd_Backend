package middleware

import (
	"fmt"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if err.Error() == fmt.Sprintf("stat %s: no such file or directory", path) {
		return false, nil
	}
	return false, err
}
