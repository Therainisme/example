package skiplist

type SkipList struct {
	sentinel *node
}

func Create(array []int) *SkipList {
	sl := &SkipList{
		sentinel: &node{
			id:    -1,
			nexts: make([]*node, 0),
		},
	}

	// linked list
	sl.sentinel.addLevel()
	prevNode := sl.sentinel

	for _, x := range array {
		newNode := &node{}
		newNode.id = x
		newNode.nexts = append(newNode.nexts, nil)

		prevNode.nexts[0] = newNode
		prevNode = newNode
	}

	// add level
	currentLevel := 0
	for sl.sentinel.nexts[currentLevel] != nil {
		currentLevel++

		sl.sentinel.addLevel()

		prevNode := sl.sentinel
		for prevLevelNode := sl.sentinel.nexts[currentLevel-1]; prevLevelNode != nil; prevLevelNode = prevLevelNode.nexts[currentLevel-1] {
			if tossCoin() {
				prevLevelNode.addLevel()

				prevNode.nexts[currentLevel] = prevLevelNode
				prevNode = prevLevelNode
			}
		}
	}

	return sl
}

func (sl *SkipList) Insert(id int) {
	// todo
}

func (sl *SkipList) Search(id int) (*node, bool) {
	currentLevel := len(sl.sentinel.nexts) - 1

	for nowNode := sl.sentinel; ; {
		// fmt.Printf("%v\n", nowNode)
		if nowNode.id == id {
			return nowNode, true
		} else if len(nowNode.nexts)-1 >= currentLevel && nowNode.nexts[currentLevel] != nil && nowNode.nexts[currentLevel].id <= id {
			nowNode = nowNode.nexts[currentLevel]
		} else if currentLevel > 0 {
			currentLevel--
		} else {
			return nil, false
		}
	}
}

func (sl *SkipList) Delete(id int) {
	// todo
}

func (s *node) addLevel() {
	s.nexts = append(s.nexts, nil)
}

func (s *node) getTopLevel() *node {
	return s.nexts[len(s.nexts)-1]
}

type node struct {
	id    int
	nexts []*node
}
