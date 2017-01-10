package bitslice

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type BitSliceSuite struct{}

var _ = Suite(&BitSliceSuite{})

func (s *BitSliceSuite) TestNewEmptyBitSlice(c *C) {
	nebs := NewEmptyBitSlice()
	c.Assert(nebs.length, Equals, 0)
	c.Assert(nebs.data, HasLen, 0)
}

func (s *BitSliceSuite) TestNewBitSlice(c *C) {
	nbs := NewBitSlice(8)
	c.Assert(nbs.length, Equals, 8)
	c.Assert(nbs.data, HasLen, 1)

	nbs1 := NewBitSlice(12)
	c.Assert(nbs1.length, Equals, 12)
	c.Assert(nbs1.data, HasLen, 2)

	nbs2 := NewBitSlice(16)
	c.Assert(nbs2.length, Equals, 16)
	c.Assert(nbs2.data, HasLen, 2)
}

func (s *BitSliceSuite) TestUnusedSpace(c *C) {
	bs := NewBitSlice(9)
	c.Assert(bs.unusedSpace(), Equals, 7)

	bs1 := NewBitSlice(3)
	c.Assert(bs1.unusedSpace(), Equals, 5)

	bs2 := NewBitSlice(8)
	c.Assert(bs2.unusedSpace(), Equals, 0)

	bs3 := NewEmptyBitSlice()
	c.Assert(bs3.unusedSpace(), Equals, 0)
}

func (s *BitSliceSuite) TestSet(c *C) {
	bs := NewBitSlice(12)
	err := bs.Set(12)
	c.Assert(err, ErrorMatches, "Index 12 out of range")

	err = bs.Set(11)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, byte(8))

	err = bs.Set(2)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, byte(4))
}

func (s *BitSliceSuite) TestUnset(c *C) {
	bs := NewBitSlice(12)

	err := bs.Set(11)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, byte(8))
	err = bs.Unset(11)
	c.Assert(err, IsNil)
	c.Assert(bs.data[1], Equals, byte(0))

	err = bs.Set(2)
	c.Assert(err, IsNil)
	err = bs.Set(3)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, byte(12))
	err = bs.Unset(2)
	c.Assert(err, IsNil)
	c.Assert(bs.data[0], Equals, byte(8))
}
