package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

var rd *rand.Rand

func init() {
	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func tossCoin() bool {
	return rd.Int()%2 == 0
}

func (sl *SkipList) Print() {
	i := len(sl.sentinel.nexts) - 1

	for i >= 0 {
		fmt.Printf("hd -> ")
		for prevNode := sl.sentinel.nexts[i]; prevNode != nil; prevNode = prevNode.nexts[i] {
			fmt.Printf("%d -> ", prevNode.id)
		}
		fmt.Printf("nil\n")

		i--
	}
}
