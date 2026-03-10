// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

package nl

import (
	"syscall"
	"unsafe"
)

type RtMsg struct {
	syscall.RtMsg
}

func NewRtMsg() *RtMsg {
	return &RtMsg{
		RtMsg: syscall.RtMsg{
			Table:    syscall.RT_TABLE_MAIN,
			Scope:    syscall.RT_SCOPE_UNIVERSE,
			Protocol: syscall.RTPROT_BOOT,
			Type:     syscall.RTN_UNICAST,
		},
	}
}

func NewRtDelMsg() *RtMsg {
	return &RtMsg{
		RtMsg: syscall.RtMsg{
			Table: syscall.RT_TABLE_MAIN,
			Scope: syscall.RT_SCOPE_NOWHERE,
		},
	}
}

func (msg *RtMsg) Len() int {
	return syscall.SizeofRtMsg
}

func DeserializeRtMsg(b []byte) *RtMsg {
	return (*RtMsg)(unsafe.Pointer(&b[0:syscall.SizeofRtMsg][0]))
}

func (msg *RtMsg) Serialize() []byte {
	return (*(*[syscall.SizeofRtMsg]byte)(unsafe.Pointer(msg)))[:]
}
