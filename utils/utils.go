package utils

import (
	"os/user"
)

func HomeDir() string {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	return u.HomeDir
}

func Contains(strs []string, s string) bool {
	for _, v := range strs {
		if v == s {
			return true
		}
	}

	return false
}
