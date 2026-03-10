// Copyright IBM Corp. 2014, 2018
// SPDX-License-Identifier: Apache-2.0

package configs

type HugepageLimit struct {
	// which type of hugepage to limit.
	Pagesize string `json:"page_size"`

	// usage limit for hugepage.
	Limit uint64 `json:"limit"`
}
