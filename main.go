package main

import (
	"fmt"
	"unsafe"
)

func main() {
	inputs := []S{
		S("aaaaaaaaa"),
		S("tgkvnvn1r"),
	}
	for _, input := range inputs {
		F(input)
		F(input)
	}
}

func hash(a interface{}) int {
	s := a.(S)
	return s.Hash()
}

type S string

func (v S) Hash() int {
	h := Hash(0, 0, *(*[]byte)(unsafe.Pointer(&v)))
	return int(h)
}

func F(k interface{}) {
	h1 := hash(k)
	defer func() {
		h2 := hash(k)
		if h1 != h2 {
			// panic: hashes differ: tgkvnvn1r 5561414296527422833 1767073917199631498
			panic(fmt.Sprintf("hashes differ: %v %v %v", k, h1, h2))
		}
	}()
}

// Copied implementation from  github.com/dchest/siphash v1.0.0

const BlockSize = 8

// Hash returns the 64-bit SipHash-2-4 of the given byte slice with two 64-bit
// parts of 128-bit key: k0 and k1.
func Hash(k0, k1 uint64, p []byte) uint64 {
	// Initialization.
	v0 := k0 ^ 0x736f6d6570736575
	v1 := k1 ^ 0x646f72616e646f6d
	v2 := k0 ^ 0x6c7967656e657261
	v3 := k1 ^ 0x7465646279746573
	t := uint64(len(p)) << 56

	// Compression.
	for len(p) >= BlockSize {
		m := uint64(p[0]) | uint64(p[1])<<8 | uint64(p[2])<<16 | uint64(p[3])<<24 |
			uint64(p[4])<<32 | uint64(p[5])<<40 | uint64(p[6])<<48 | uint64(p[7])<<56
		v3 ^= m

		// Round 1.
		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)

		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2

		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0

		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)

		// Round 2.
		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)

		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2

		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0

		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)

		v0 ^= m
		p = p[BlockSize:]
	}

	// Compress last block.
	switch len(p) {
	case 7:
		t |= uint64(p[6]) << 48
		fallthrough
	case 6:
		t |= uint64(p[5]) << 40
		fallthrough
	case 5:
		t |= uint64(p[4]) << 32
		fallthrough
	case 4:
		t |= uint64(p[3]) << 24
		fallthrough
	case 3:
		t |= uint64(p[2]) << 16
		fallthrough
	case 2:
		t |= uint64(p[1]) << 8
		fallthrough
	case 1:
		t |= uint64(p[0])
	}

	v3 ^= t

	// Round 1.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	// Round 2.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	v0 ^= t

	// Finalization.
	v2 ^= 0xff

	// Round 1.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	// Round 2.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	// Round 3.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	// Round 4.
	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)

	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2

	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0

	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	return v0 ^ v1 ^ v2 ^ v3
}
