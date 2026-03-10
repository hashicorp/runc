// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

package mount

// GetMounts retrieves a list of mounts for the current running process.
func GetMounts() ([]*Info, error) {
	return parseMountTable()
}

// Mounted looks at /proc/self/mountinfo to determine of the specified
// mountpoint has been mounted
func Mounted(mountpoint string) (bool, error) {
	entries, err := parseMountTable()
	if err != nil {
		return false, err
	}

	// Search the table for the mountpoint
	for _, e := range entries {
		if e.Mountpoint == mountpoint {
			return true, nil
		}
	}
	return false, nil
}
