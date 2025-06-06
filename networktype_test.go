// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkType_String(t *testing.T) {
	testCases := []struct {
		cType          NetworkType
		expectedString string
	}{
		{NetworkTypeUnknown, ErrUnknownType.Error()},
		{NetworkTypeUDP4, "udp4"},
		{NetworkTypeUDP6, "udp6"},
		{NetworkTypeTCP4, "tcp4"},
		{NetworkTypeTCP6, "tcp6"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.expectedString,
			testCase.cType.String(),
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestNetworkType(t *testing.T) {
	testCases := []struct {
		typeString   string
		shouldFail   bool
		expectedType NetworkType
	}{
		{ErrUnknownType.Error(), true, NetworkTypeUnknown},
		{"udp4", false, NetworkTypeUDP4},
		{"udp6", false, NetworkTypeUDP6},
		{"tcp4", false, NetworkTypeTCP4},
		{"tcp6", false, NetworkTypeTCP6},
	}

	for i, testCase := range testCases {
		actual, err := NewNetworkType(testCase.typeString)
		if testCase.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t,
			testCase.expectedType,
			actual,
			"testCase: %d %v", i, testCase,
		)
	}
}
