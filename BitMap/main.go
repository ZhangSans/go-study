package main

import "fmt"

const (
	bitsPerByte  = 8
	maxByteValue = 255
)

type BitMap struct {
	data []byte
	size int64
}

func NewBitMap() *BitMap {
	size := int64(19999999999)
	numBytes := (size + bitsPerByte - 1) / bitsPerByte
	return &BitMap{
		data: make([]byte, numBytes),
		size: size,
	}
}

func (b *BitMap) Set(index int64) {
	if index < 0 || index >= b.size {
		panic("index out of range")
	}
	byteIndex := index / bitsPerByte
	bitIndex := index % bitsPerByte
	b.data[byteIndex] |= 1 << (bitsPerByte - 1 - bitIndex)
}
func (b *BitMap) Get(index int64) bool {
	if index < 0 || index >= b.size {
		panic("index out of range")
	}
	byteIndex := index / bitsPerByte
	bitIndex := index % bitsPerByte

	return (b.data[byteIndex] & (1 << (bitsPerByte - 1 - bitIndex))) != 0
}

func (b *BitMap) Clear(index int64) {
	if index < 0 || index >= b.size {
		panic("index out of range")
	}
	byteIndex := index / bitsPerByte
	bitIndex := index % bitsPerByte
	b.data[byteIndex] &= ^(1 << (bitsPerByte - 1 - bitIndex))
}

func main() {
	bitMap := NewBitMap()

	var id int64
	id = -1
	bitMap.Set(id)

	fmt.Println(bitMap.Get(id))
	id = 124
	fmt.Println(bitMap.Get(id))
}
