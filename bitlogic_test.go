package bitslice

import (
	. "gopkg.in/check.v1"
)

func (s *BitSliceSuite) TestOr(c *C) {
	// Test working on equally sized bitslices
	{
		bs1 := NewBitSlice(64)
		err := bs1.Set(2)
		c.Assert(err, IsNil)
		err = bs1.Set(6)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(64)
		err = bs2.Set(5)
		c.Assert(err, IsNil)

		bs1.Or(bs2)
		c.Assert(bs1.data[0], Equals, uint64(100))
	}

	// Test working on unequally sized bitslices represented by same number of buffers
	{
		bs1 := NewBitSlice(128)
		err := bs1.Set(70)
		c.Assert(err, IsNil)
		err = bs1.Set(67)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(80)
		err = bs2.Set(2)
		c.Assert(err, IsNil)
		err = bs2.Set(64)
		c.Assert(err, IsNil)

		bs1.Or(bs2)
		c.Assert(bs1.length, Equals, 128)
		c.Assert(bs1.data[1], Equals, uint64(73))
		c.Assert(bs1.data[0], Equals, uint64(4))
	}

	// Test that it works on unequally sized bitslices represented by a different number of buffers
	{
		bs1 := NewBitSlice(64)
		err := bs1.Set(0)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(96)
		err = bs2.Set(65)
		c.Assert(err, IsNil)

		bs1.Or(bs2)
		c.Assert(bs1.length, Equals, 96)
		c.Assert(bs1.data, HasLen, 2)
		c.Assert(bs1.data[1], Equals, uint64(2))
		c.Assert(bs1.data[0], Equals, uint64(1))
	}
}

func (s *BitSliceSuite) TestAnd(c *C) {
	// Test working on equally sized bitslices
	{
		bs1 := NewBitSlice(64)
		err := bs1.Set(2)
		c.Assert(err, IsNil)
		err = bs1.Set(6)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(64)
		err = bs2.Set(2)
		c.Assert(err, IsNil)

		bs1.And(bs2)
		c.Assert(bs1.data[0], Equals, uint64(4))
	}

	// Test working on unequally sized bitslices represented by same number of buffers
	{
		bs1 := NewBitSlice(128)
		err := bs1.Set(70)
		c.Assert(err, IsNil)
		err = bs1.Set(64)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(80)
		err = bs2.Set(2)
		c.Assert(err, IsNil)
		err = bs2.Set(64)
		c.Assert(err, IsNil)

		bs1.And(bs2)
		c.Assert(bs1.length, Equals, 128)
		c.Assert(bs1.data[1], Equals, uint64(1))
		c.Assert(bs1.data[0], Equals, uint64(0))
	}

	// Test that it works on unequally sized bitslices represented by a different number of buffers
	{
		bs1 := NewBitSlice(64)
		err := bs1.Set(0)
		c.Assert(err, IsNil)

		bs2 := NewBitSlice(96)
		err = bs2.Set(65)
		c.Assert(err, IsNil)

		bs1.And(bs2)
		c.Assert(bs1.length, Equals, 96)
		c.Assert(bs1.data, HasLen, 2)
		c.Assert(bs1.data[1], Equals, uint64(2))
		c.Assert(bs1.data[0], Equals, uint64(0))
	}
}

func (s *BitSliceSuite) TestBufferOnes(c *C) {
	// Test adding ones to only one uint64
	{
		bs := NewBitSlice(1)
		err := bs.Set(0)
		c.Assert(err, IsNil)

		bs.bufferOnes(3)
		c.Assert(bs.data[0], Equals, uint64(15))
	}

	// Test adding ones to a multi uint bitslice
	{
		bs := NewBitSlice(95)
		bs.bufferOnes(1)
		c.Assert(bs.data[1], Equals, uint64(2147483648))
		c.Assert(bs.data[0], Equals, uint64(0))
	}

	// Test adding overflowing ones
	{
		bs := NewBitSlice(1)
		err := bs.Set(0)
		c.Assert(err, IsNil)

		bs.bufferOnes(64)
		c.Assert(bs.data[1], Equals, uint64(1))
		c.Assert(bs.data[0], Equals, MAXUINT)
	}
}

func (s *BitSliceSuite) TestXor(c *C) {
	// TODO: Create this
}

func (s *BitSliceSuite) TestNot(c *C) {
	// TODO: Create this
}

func (s *BitSliceSuite) TestCreateMask(c *C) {
	c.Assert(createMask(3), Equals, uint64(7))
	c.Assert(createMask(1), Equals, uint64(1))
	c.Assert(createMask(5), Equals, uint64(31))
}
