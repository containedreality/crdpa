package crdpa

import (
	"fmt"
	"log"
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
