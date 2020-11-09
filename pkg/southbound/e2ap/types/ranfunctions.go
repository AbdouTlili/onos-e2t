// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type RanFunctionDescription string
type RanFunctionRevision int

type RanFunctionItem struct {
	Description RanFunctionDescription
	Revision    RanFunctionRevision
}

type RanFunctions map[int32]RanFunctionItem
