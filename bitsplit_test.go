package bitsplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitU16(t *testing.T) {
	assert := assert.New(t)

	expectedHigh, expectedLow := uint8(0x12), uint8(0x34)

	actualHi, actualLow := SplitU16(0x1234)

	assert.Equal(expectedHigh, actualHi)
	assert.Equal(expectedLow, actualLow)
}

func TestSplitU24(t *testing.T) {
	assert := assert.New(t)

	expectedHigh, expectedMed, expectedLow := uint8(0x12), uint8(0x34), uint8(0x56)

	actualHigh, actualMed, actualLow := SplitU24(0x123456)

	assert.Equal(expectedHigh, actualHigh)
	assert.Equal(expectedMed, actualMed)
	assert.Equal(expectedLow, actualLow)
}

func TestSplitU32(t *testing.T) {
	assert := assert.New(t)

	expectedHigh, expectedLow := uint16(0xAABB), uint16(0x1122)

	actualHi, actualLow := SplitU32(0xAABB1122)

	assert.Equal(expectedHigh, actualHi)
	assert.Equal(expectedLow, actualLow)
}

func TestSplitU64(t *testing.T) {
	assert := assert.New(t)

	expectedHigh, expectedLow := uint32(0xAABBCCDD), uint32(0x11223344)

	actualHi, actualLow := SplitU64(0xAABBCCDD11223344)

	assert.Equal(expectedHigh, actualHi)
	assert.Equal(expectedLow, actualLow)
}

func TestJoinUInt8s(t *testing.T) {
	assert := assert.New(t)

	actual := JoinUInt8s(0x12, 0x34)

	assert.Equal(uint16(0x1234), actual)
}

func TestJoinUInt16s(t *testing.T) {
	assert := assert.New(t)

	actual := JoinUInt16s(0x1234, 0x5678)

	assert.Equal(uint32(0x12345678), actual)
}

func TestJoinUInt24s(t *testing.T) {
	assert := assert.New(t)

	actual := JoinUInt24s(0x12, 0x34, 0x56)

	assert.Equal(uint32(0x123456), actual)
}

func TestJoinUInt32s(t *testing.T) {
	assert := assert.New(t)

	actual := JoinUInt32s(0x12341234, 0x56785678)

	assert.Equal(uint64(0x1234123456785678), actual)
}
