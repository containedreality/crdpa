package crdpa

import (
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	ROUNDS = 56
)

type CRDPA struct {
	roundkeys []uint32
}

func round(a uint32, b uint32, c uint32, d uint32, k uint32) (uint32, uint32, uint32, uint32) {
	a = bits.RotateLeft32(a, -8)
	a ^= k
	a += b

	b = bits.RotateLeft32(b, 3)
	b ^= a
	b += c

	c = bits.RotateLeft32(c, -13)
	c ^= b
	c += d

	d = bits.RotateLeft32(d, 7)
	d ^= c
	d += b

	return a, b, c, d
}

func invround(a uint32, b uint32, c uint32, d uint32, k uint32) (uint32, uint32, uint32, uint32) {
	d -= b
	d ^= c
	d = bits.RotateLeft32(d, -7)

	c -= d
	c ^= b
	c = bits.RotateLeft32(c, 13)

	b -= c
	b ^= a
	b = bits.RotateLeft32(b, -3)

	a -= b
	a ^= k
	a = bits.RotateLeft32(a, 8)

	return a, b, c, d
}

func keyschedule(key []byte) []uint32 {
	var roundkeys []uint32

	a := binary.LittleEndian.Uint32(key[0:4])
	b := binary.LittleEndian.Uint32(key[4:8])
	c := binary.LittleEndian.Uint32(key[8:12])
	d := binary.LittleEndian.Uint32(key[12:16])
	e := binary.LittleEndian.Uint32(key[16:20])
	f := binary.LittleEndian.Uint32(key[20:24])
	g := binary.LittleEndian.Uint32(key[24:28])
	h := binary.LittleEndian.Uint32(key[28:32])

	for i := range ROUNDS {
		roundkeys = append(roundkeys, a)

		a, b, c, d = round(a, b, c, d, uint32(i))
		a, f, g, h = round(a, e, f, g, h)
	}

	return roundkeys
}

func (crdpa *CRDPA) Encrypt(dst, src []byte) {
	a := binary.LittleEndian.Uint32(src[0:4])
	b := binary.LittleEndian.Uint32(src[4:8])
	c := binary.LittleEndian.Uint32(src[8:12])
	d := binary.LittleEndian.Uint32(src[12:16])

	for i := range ROUNDS {
		a, b, c, d = round(a, b, c, d, crdpa.roundkeys[i])
	}

	binary.LittleEndian.PutUint32(dst[0:4], a)
	binary.LittleEndian.PutUint32(dst[4:8], b)
	binary.LittleEndian.PutUint32(dst[8:12], c)
	binary.LittleEndian.PutUint32(dst[12:16], d)
}

func (crdpa *CRDPA) Decrypt(dst, src []byte) {
	a := binary.LittleEndian.Uint32(src[0:4])
	b := binary.LittleEndian.Uint32(src[4:8])
	c := binary.LittleEndian.Uint32(src[8:12])
	d := binary.LittleEndian.Uint32(src[12:16])

	for i := ROUNDS - 1; i >= 0; i-- {
		a, b, c, d = invround(a, b, c, d, crdpa.roundkeys[i])
	}

	binary.LittleEndian.PutUint32(dst[0:4], a)
	binary.LittleEndian.PutUint32(dst[4:8], b)
	binary.LittleEndian.PutUint32(dst[8:12], c)
	binary.LittleEndian.PutUint32(dst[12:16], d)
}

func New(key []byte) (*CRDPA, error) {
	var crdpa CRDPA
	if len(key) != 32 {
		return &crdpa, errors.New("invalid key length")
	}

	crdpa.roundkeys = keyschedule(key)

	return &crdpa, nil
}
