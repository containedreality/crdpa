package crdpa

import (
	"crypto/cipher"
	"fmt"
	"log"
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	key := []byte("AAAABBBBCCCCDDDDEEEEFFFFGGGGHHHH")

	b, err := New(key)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	plaintext := []byte("AAAABBBBCCCCDDDD")
	buf := make([]byte, 16)
	buf2 := make([]byte, 16)

	b.Encrypt(buf, plaintext)

	fmt.Println(buf)

	b.Decrypt(buf2, buf)
	fmt.Println(buf2)
}

func TestEntropy(t *testing.T) {
	key := make([]byte, 32)
	iv := make([]byte, 16)
	for i := range key {
		key[i] = byte(i * 3)
	}
	for i := range iv {
		iv[i] = byte(i * 7)
	}

	block, err := New(key)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	stream := cipher.NewCTR(block, iv)

	dataSize := (1024 * 1024) * 10
	plaintext := make([]byte, dataSize)
	ciphertext := make([]byte, dataSize)

	stream.XORKeyStream(ciphertext, plaintext)

	var freq [256]float64
	for _, b := range ciphertext {
		freq[b]++
	}

	var entropy float64
	length := float64(dataSize)
	for _, count := range freq {
		if count > 0 {
			p := count / length
			entropy -= p * math.Log2(p)
		}
	}

	t.Logf("entropy: %f", entropy)

	if entropy < 7.99 {
		t.Errorf("entropy too low: got %f, expected >= 7.99", entropy)
	}
}

func TestAvalanche(t *testing.T) {
	key := make([]byte, 32)
	block, err := New(key)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	pt1 := make([]byte, 16)
	pt2 := make([]byte, 16)
	pt2[0] ^= 0x01

	ct1 := make([]byte, 16)
	ct2 := make([]byte, 16)

	block.Encrypt(ct1, pt1)
	block.Encrypt(ct2, pt2)

	var diffBits int
	for i := 0; i < 16; i++ {
		xor := ct1[i] ^ ct2[i]
		for bit := 0; bit < 8; bit++ {
			if (xor & (1 << bit)) != 0 {
				diffBits++
			}
		}
	}

	t.Logf("bit flip differences: %d / 128 bits changed", diffBits)

	if diffBits < 40 || diffBits > 88 {
		t.Errorf("avalanche effect out of bounds: %d bits changed (ideal ~64)", diffBits)
	}
}

func BenchmarkCRDPA(b *testing.B) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}

	block, err := New(key)
	if err != nil {
		log.Println(err)
		b.Fail()
	}

	plaintext := []byte("AAAABBBBCCCCDDDD")

	buf := make([]byte, 16)

	for b.Loop() {
		block.Encrypt(buf, plaintext)
	}
}
