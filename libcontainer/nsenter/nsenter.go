// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

// +build linux,!gccgo

package nsenter

/*
#cgo CFLAGS: -Wall
extern void nsexec();
void __attribute__((constructor)) init(void) {
	nsexec();
}
*/
import "C"
