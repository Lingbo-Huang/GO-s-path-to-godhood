package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray []bool
	size     int
	hashes   []func(data []byte) uint32
}

func NewBloomFilter(size int, hashes []func(data []byte) uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		size:     size,
		hashes:   hashes,
	}
}

func (bf *BloomFilter) Add(data []byte) {
	for _, hash := range bf.hashes {
		position := hash(data) % uint32(bf.size)
		bf.bitArray[position] = true
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hash := range bf.hashes {
		position := hash(data) % uint32(bf.size)
		if !bf.bitArray[position] {
			return false
		}
	}
	return true
}

func hash1(data []byte) uint32 {
	h := fnv.New32a()
	h.Write(data)
	return h.Sum32()
}

func hash2(data []byte) uint32 {
	var hash uint32
	for _, b := range data {
		hash = hash*31 + uint32(b)
	}
	return hash
}

func main() {
	hashes := []func(data []byte) uint32{hash1, hash2}
	bf := NewBloomFilter(1000, hashes)

	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))

	fmt.Println(bf.Contains([]byte("hello"))) // Output: true
	fmt.Println(bf.Contains([]byte("world"))) // Output: true
	fmt.Println(bf.Contains([]byte("foo")))   // Output: false
}
