// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"gotest.tools/assert"
	"testing"
)

func TestNewGlobalgNBID(t *testing.T) {

	g := e2apies.GlobalgNbId{
		PlmnId: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte("ONF"),
		},
		GnbId: &e2apies.GnbIdChoice{
			GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
				GnbId: &e2ap_commondatatypes.BitString{
					Value: 0x9ABCDEF000000000,
					Len:   28,
				},
			},
		},
	}

	cobject, err := newGlobalgNBID(&g)
	assert.NilError(t, err, "error converting to c struct")
	assert.Assert(t, cobject != nil)
	assert.Equal(t, int(cobject.plmn_id.size), 3, "expected plmn id to be 3 bytes")
	assert.Equal(t, int(cobject.gnb_id.present), 1, "expected choice to be 1 (gnb_id)")

	// Now do the reverse - C object back to struct
	g1, err := decodeGlobalGnbID(cobject)
	assert.NilError(t, err, "error converting back from c struct")
	assert.Assert(t, g1 != nil)
	assert.Equal(t, string(g1.PlmnId.Value), "ONF", "unexpected value for Plmn ID")
	switch choice := g1.GnbId.GnbIdChoice.(type) {
	case *e2apies.GnbIdChoice_GnbId:
		assert.Equal(t, int(choice.GnbId.Len), 28)
		assert.Equal(t, choice.GnbId.Value, uint64(0x9ABCDEF000000000))
	default:
		t.Fatalf("unexpected choice in GnbIdChoice %v", choice)
	}

	xer, err := xerEncodegNBID(&g)
	assert.NilError(t, err)
	t.Logf("XER GlobalgNbId: \n%s", string(xer))

	//per, err := perEncodegNBID(&g)
	//assert.NilError(t, err)
	//t.Logf("PER GlobalgNbId: \n%s", string(per))
}
