package bitslice

func (t *BitSlice) Or(bslice *BitSlice) {
	if len(bslice.data) > len(t.data) {
		diff := len(bslice.data) - len(t.data)
		post := make([]uint64, diff)
		t.data = append(t.data, post...)
	}

	if bslice.length > t.length {
		t.length = bslice.length
	}

	for i, b := range bslice.data {
		t.data[i] |= b
	}
}

func (t *BitSlice) And(bslice *BitSlice) {
	// Logic for making this operation non-destructive
	// If one is longer than the other, the excess bits will be carried over and not lost
	/*
		if bslice.length != t.length {
			if bslice.length > t.length {
				t.bufferOnes(bslice.length - t.length)
			} else {
				bslice.bufferOnes(t.length - bslice.length)
			}
		}
	*/

	if bslice.length > t.length {
		t.length = bslice.length

		if len(bslice.data) > len(t.data) {
			extra := make([]uint64, len(bslice.data)-len(t.data))
			t.data = append(t.data, extra...)
		}
	}

	for i, b := range bslice.data {
		t.data[i] &= b
	}
}

func (t *BitSlice) Xor(bslice *BitSlice) {
	if bslice.length > t.length {
		t.length = bslice.length

		if len(bslice.data) > len(t.data) {
			extra := make([]uint64, len(bslice.data)-len(t.data))
			t.data = append(t.data, extra...)
		}
	}

	for i, b := range bslice.data {
		t.data[i] ^= b
	}
}

func (t *BitSlice) Not() {
	for i := 0; i < len(t.data); i++ {
		t.data[i] = ^t.data[i]
		if i == len(t.data)-1 {
			mask := createMask(64 - t.unusedSpace())
			t.data[i] &= mask
		}
	}
}

func (t *BitSlice) bufferOnes(amount int) {
	// Helper function to add amount ones to the end of a bitslice
	space := t.unusedSpace()
	oldLength := len(t.data)
	add := (amount - space) / 64
	if (amount-space)%64 > 0 {
		add++
	}
	extra := make([]uint64, add)
	t.data = append(t.data, extra...)

	t.length += amount
	for amount > 0 {
		offset := uint(64 - space)
		mask := createMask(amount) << offset
		t.data[oldLength-1] |= mask
		amount -= space
		oldLength++
		// Set available space to the entirety of a uint64
		space = 64
	}
}

func createMask(length int) uint64 {
	val := uint64(0)
	if length > 63 {
		return MAXUINT
	} else if length < 0 {
		return val
	}

	val = MAXUINT
	val >>= uint(64 - length)

	return val
}
