package main

import (
	"encoding/hex"
)

type cycle struct {
	list     [256]byte
	position int
	skip     int
}

func newCycle() (c cycle) {
	for i := range c.list {
		c.list[i] = byte(i)
	}
	c.skip = 1
	return
}

func (c *cycle) checksum() int {
	return int(c.list[0]) * int(c.list[1])
}

func (c *cycle) run(i int) {
	c.reverse(i)
	c.position += c.skip
	c.skip++
}

func (c *cycle) reverse(i int) {
	from := c.position
	to := from + i - 1
	c.position = to
	for {
		if from > to {
			break
		}
		c.list[from%256], c.list[to%256] = c.list[to%256], c.list[from%256]
		from++
		to--
	}
}

func (c *cycle) denseHash() string {
	return denseHash(c.list)
}

func denseHash(sparseHash [256]byte) string {
	var denseBytes [16]byte
	for i := 0; i < 16; i++ {
		var arr [16]byte
		copy(arr[:], sparseHash[(i*16):((i+1)*16)])
		denseBytes[i] = mergeBytes(arr)
	}
	return toHex(denseBytes)
}

func mergeBytes(bytes [16]byte) (b byte) {
	for i := 0; i < 16; i++ {
		b ^= bytes[i]
	}
	return
}

func toHex(bytes [16]byte) (s string) {
	return hex.EncodeToString(bytes[:])
}
