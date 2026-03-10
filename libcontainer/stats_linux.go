// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

package libcontainer

import "github.com/opencontainers/runc/libcontainer/cgroups"
import "github.com/opencontainers/runc/libcontainer/intelrdt"

type Stats struct {
	Interfaces    []*NetworkInterface
	CgroupStats   *cgroups.Stats
	IntelRdtStats *intelrdt.Stats
}
