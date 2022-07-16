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
	currentLevel := len(sl.sentinel.nexts) - 1

	startFrom := make([]*node, len(sl.sentinel.nexts))
	copy(startFrom, sl.sentinel.nexts)

	nowNode := sl.sentinel
	for {
		// fmt.Printf("%v\n", nowNode)
		if currentLevel == 0 {
			// continue next
			for nowNode.nexts[currentLevel] != nil && nowNode.nexts[currentLevel].id <= id {
				nowNode = nowNode.nexts[currentLevel]
			}

			startFrom[currentLevel] = nowNode
			break
		} else if len(nowNode.nexts)-1 >= currentLevel && nowNode.nexts[currentLevel] != nil && nowNode.nexts[currentLevel].id <= id {
			nowNode = nowNode.nexts[currentLevel]
		} else if currentLevel > 0 {
			startFrom[currentLevel] = nowNode
			currentLevel--
		} else {
			break
		}
	}

	newNode := &node{id, nil}
	newNode.addLevel()
	newNode.nexts[0] = startFrom[0].nexts[0]
	startFrom[0].nexts[0] = newNode
	newNodeLevel := 0

	for newNodeLevel < len(startFrom)-1 && tossCoin() {
		newNodeLevel++
		newNode.addLevel()
		newNode.nexts[newNodeLevel] = startFrom[newNodeLevel].nexts[newNodeLevel]
		if startFrom[newNodeLevel] != nil {
			startFrom[newNodeLevel].nexts[newNodeLevel] = newNode
		} else {
			sl.sentinel.nexts[newNodeLevel] = newNode
		}
	}
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
