// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

// +build cgo,linux

package system

/*
#include <unistd.h>
*/
import "C"

func GetClockTicks() int {
	return int(C.sysconf(C._SC_CLK_TCK))
}
