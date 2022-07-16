package skiplist

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRandomNumber(t *testing.T) {
	n, front := 1000000, 0
	for i := 0; i < n; i++ {
		if tossCoin() {
			front++
		}
	}

	if math.Abs(float64(front)/float64(n)-0.5) > 0.01 {
		t.Fatalf("probability of front is %f, which exceeds the allowable error range.",
			math.Abs(float64(front)/float64(n)-0.5),
		)
	}
}

func TestLv0(t *testing.T) {
	array := []int{2, 8, 10, 11, 13, 19, 20, 22, 26, 30}
	sl := Create(array)
	lv0 := sl.sentinel

	node := lv0.nexts[0]
	for i := 0; i < len(array); i++ {
		if node == nil || node.id != array[i] {
			if node != nil {
				t.Logf("array[%d]=%d, node[%d]=%d", i, array[i], i, node.id)
			}
			t.Fatalf("level 0 does not match at %d", i)
		}

		node = node.nexts[0]
	}

	// Print
	// sl.Print()
}

func TestSearch(t *testing.T) {
	array := []int{2, 8, 10, 11, 13, 19, 20, 22, 26, 30}
	sl := Create(array)

	// Print
	sl.Print()

	nd, ok := sl.Search(2)
	if !(ok && nd.id == 2) {
		t.Fatalf("search 2 error")
	}

	nd, ok = sl.Search(20)
	if !(ok && nd.id == 20) {
		t.Fatalf("search 20 error")
	}

	nd, ok = sl.Search(33)
	if ok || nd != nil && nd.id == 33 {
		t.Fatalf("search 33 error, get result %v", nd)
	}
}

func TestSearchWithLargeAmountOfData(t *testing.T) {
	rd = rand.New(rand.NewSource(time.Now().UnixNano()))

	check := make(map[int]bool, 0)
	array := make([]int, 0)

	n := 10000
	for i := 0; i < n; i++ {
		x := rd.Int() % n
		if _, ok := check[x]; ok {
			continue
		}

		check[x] = true
		array = append(array, x)
	}

	// sort
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	// create skip list
	sl := Create(array)
	for i := 0; i < n; i++ {
		x := rd.Int() % n

		_, ckok := check[x]
		nd, slok := sl.Search(x)

		checkResult(t, x, ckok, slok, nd)
	}
}

func checkResult(t *testing.T, x int, ckok, slok bool, nd *node) {
	if ckok != slok {
		t.Fatalf("array and skiplist no match")
	}

	if !ckok && (nd != nil && nd.id == x) {
		t.Fatalf("get a value that does not exist in array, %d", x)
	}

	if ckok && (nd == nil || nd.id != x) {
		t.Fatalf("get a value that unexpected in array, %d", x)
	}
}
