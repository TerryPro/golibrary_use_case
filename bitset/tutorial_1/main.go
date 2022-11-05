package main

import (
	"fmt"
	"math/rand"

	"github.com/bits-and-blooms/bitset"
)

func use_1() {
	b := bitset.New(64)
	// 放入一个数
	b.Set(10)
	fmt.Println(b.DumpAsBits()) // 000000000000000000000000000000000000000000000000010000000000
	// 删除一个值
	b.Clear(10)
	fmt.Println(b.DumpAsBits()) // 000000000000000000000000000000000000000000000000000000000000
	// 长度
	b.Set(1).Set(3)
	fmt.Println(b.DumpAsBits()) // 000000000000000000000000000000000000000000000000000000000000

	fmt.Println(b.Len()) // 64
	// 测试
	fmt.Println(b.Test(3)) // true
}

func use_2() {
	fmt.Printf("Hello from BitSet!\n")
	var b bitset.BitSet
	// play some Go Fish
	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52))
		card2 := uint(rand.Intn(52))
		b.Set(card1)
		if b.Test(card2) {
			fmt.Println("Go Fish!")
		}
		b.Clear(card1)
	}

	// Chaining
	b.Set(10).Set(11)

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}

}
func main() {
	use_1()
	use_2()
}
