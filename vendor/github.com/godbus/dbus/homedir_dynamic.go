// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

// +build !static_build

package dbus

import (
	"os/user"
)

func lookupHomeDir() string {
	u, err := user.Current()
	if err != nil {
		return "/"
	}
	return u.HomeDir
}
