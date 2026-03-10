// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

// +build darwin freebsd openbsd netbsd dragonfly
// +build !appengine

package logrus

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA

type Termios syscall.Termios
