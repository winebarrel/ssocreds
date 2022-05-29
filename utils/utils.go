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
