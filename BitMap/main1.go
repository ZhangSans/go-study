package main

import "fmt"

const (
    bitsPerByte  = 8  // 每个字节的位数（固定为8）
    maxByteValue = 255 // 单字节最大值（未直接使用，可移除）
)

type BitMap struct {
    data []byte // 字节数组存储位标志
    size int64  // 位图总容量（最多可表示的整数个数）
}

// 构造函数：创建指定容量的 Bitmap
func NewBitMap() *BitMap {
    size := int64(19_999_999_999) // 示例容量（可根据需求修改）
    numBytes := (size + bitsPerByte - 1) / bitsPerByte // 计算所需字节数
    return &BitMap{
        data: make([]byte, numBytes),
        size: size,
    }
}

// 设置某一位为1
func (b *BitMap) Set(index int64) {
    if index < 0 || index >= b.size {
        panic("index out of range")
    }
    byteIndex := index / bitsPerByte  // 确定字节位置
    bitIndex := index % bitsPerByte   // 确定位在字节中的位置
    b.data[byteIndex] |= 1 << bitIndex // 将指定位设为1
}

// 检查某一位是否为1
func (b *BitMap) Get(index int64) bool {
    if index < 0 || index >= b.size {
        panic("index out of range")
    }
    byteIndex := index / bitsPerByte
    bitIndex := index % bitsPerByte
    return (b.data[byteIndex] & (1 << bitIndex)) != 0 // 位与运算判断标志位
}

// 清除某一位（设为0）
func (b *BitMap) Clear(index int64) {
    if index < 0 || index >= b.size {
        panic("index out of range")
    }
    byteIndex := index / bitsPerByte
    bitIndex := index % bitsPerByte
    b.data[byteIndex] &= ^(1 << bitIndex) // 取反后位与运算
}

func main() {
    bm := NewBitMap()      // 创建 Bitmap 实例
    var id int64 = 123     
    bm.Set(id)             // 设置第123位为1
    fmt.Println(bm.Get(id)) // 输出 true
}
