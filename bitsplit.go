package bitsplit

func SplitU16(i uint16) (hi, lo uint8) {
	lo = uint8(i & 0xFF)
	hi = uint8(i >> 8)

	return
}

func SplitU24(i uint32) (hi, med, lo uint8) {
	// Why 24? Some weird formats might use this

	lo = uint8(i & 0xFF)
	med = uint8(i>>8) & 0xFF
	hi = uint8(i>>16) & 0xFF

	return
}

func SplitU32(i uint32) (hi, lo uint16) {
	lo = uint16(i & 0xFFFF)
	hi = uint16(i >> 16)

	return
}

func SplitU64(i uint64) (hi, lo uint32) {
	lo = uint32(i & 0xFFFFFFFF)
	hi = uint32(i >> 32)

	return
}

func JoinUInt8s(hi uint8, lo uint8) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}

func JoinUInt16s(hi uint16, lo uint16) uint32 {
	return (uint32(hi) << 16) | uint32(lo)
}

// NB: be careful about endianess in your application
func JoinUInt24s(hi, med, lo uint8) uint32 {
	// Since we don't have uint24, we'll use a 32bit with some zeroes leftover

	return (uint32(hi) << 16) | (uint32(med) << 8) | uint32(lo)
}

func JoinUInt32s(hi uint32, lo uint32) uint64 {
	return (uint64(hi) << 32) | uint64(lo)
}
