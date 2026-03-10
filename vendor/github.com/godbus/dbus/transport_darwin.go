// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

package dbus

func (t *unixTransport) SendNullByte() error {
	_, err := t.Write([]byte{0})
	return err
}
