package bestring

import (
	"container/list"
)

type avlNode struct {
	value    uint32
	height   int
	isDelete bool

	left  *avlNode
	right *avlNode
}

type avlTree struct {
	height int
	total  int

	root *avlNode
}

func newAVLTree() *avlTree {
	return &avlTree{}
}

func (a *avlTree) add(ele uint32) {
	a.root = a.insert(a.root, ele)
	a.height = a.root.height
}

func (a *avlTree) delete(ele uint32) int {
	var node *avlNode = a.root

	for node != nil {
		if node.value == ele {
			node.isDelete = true
		}

		if node.value > ele {
			node = node.left
		} else {
			node = node.right
		}
	}

	a.total -= 1
	return a.total
}

func (a *avlTree) findLatestLeft(ele uint32) uint32 {
	stack := list.New()

	var node *avlNode = a.root
	var res = make([]uint32, 0, a.total)

	for stack.Len() != 0 || node != nil {
		for node != nil && node.left != nil {
			stack.PushBack(node)
			node = node.left
		}

		if stack.Len() != 0 || node != nil {
			if node == nil {
				node = stack.Remove(stack.Back()).(*avlNode)
			}

			if !node.isDelete {
				res = append(res, node.value)
			}

			if len(res) >= 2 && res[len(res)-2] <= ele && ele < res[len(res)-1] {
				return res[len(res)-2]
			}
			node = node.right
		}
	}

	if ele < res[0] || ele >= res[len(res)-1] {
		return res[len(res)-1]
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] <= ele && res[i+1] > ele {
			return res[i]
		}
	}

	return res[len(res)-1]
}

func (a *avlTree) insert(node *avlNode, ele uint32) *avlNode {
	if node == nil {
		a.total += 1
		return &avlNode{value: ele, height: 1}
	}

	stack := list.New()

	for {
		if ele == node.value {
			node.isDelete = false
			return a.root
		}

		if ele < node.value {
			if node.left != nil {
				stack.PushBack(node)
				node = node.left
			} else {
				node.left = &avlNode{value: ele, height: 1}
				a.ajustHeight(node)
				break
			}
		} else {
			if node.right != nil {
				stack.PushBack(node)
				node = node.right
			} else {
				node.right = &avlNode{value: ele, height: 1}
				a.ajustHeight(node)
				break
			}
		}
	}
	a.total += 1

	var leftHeight, rightHeight, dif int
	var father *avlNode

	for stack.Len() > 0 {
		node = stack.Remove(stack.Back()).(*avlNode)

		leftHeight = a.getNodeHeight(node.left)
		rightHeight = a.getNodeHeight(node.right)
		dif = leftHeight - rightHeight

		if dif == 2 {
			if node.left.value > ele {
				node = a.llRotation(node)
			} else {
				node = a.lrRotation(node)
			}
		}

		if dif == -2 {
			if node.right.value > ele {
				node = a.rlRotation(node)
			} else {
				node = a.rrRotation(node)
			}
		}

		if (dif == 2 || dif == -2) && stack.Len() > 0 {
			father = stack.Back().Value.(*avlNode)

			if father.value < ele {
				father.right = node
				a.ajustHeight(father)
			} else {
				father.left = node
				a.ajustHeight(father)
			}
		}
	}
	return node
}

func (a *avlTree) getNodeHeight(node *avlNode) int {
	if node == nil {
		return 0
	}

	return node.height
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *avlTree) ajustHeight(node *avlNode) {
	node.height = maxInt(a.getNodeHeight(node.left), a.getNodeHeight(node.right)) + 1
}

func (a *avlTree) llRotation(node *avlNode) *avlNode {
	nodeLeft := node.left
	node.left = node.left.right
	nodeLeft.right = node

	a.ajustHeight(node)
	a.ajustHeight(nodeLeft)
	return nodeLeft
}

func (a *avlTree) rrRotation(node *avlNode) *avlNode {
	nodeRight := node.right
	node.right = node.right.left
	nodeRight.left = node

	a.ajustHeight(node)
	a.ajustHeight(nodeRight)
	return nodeRight
}

func (a *avlTree) lrRotation(node *avlNode) *avlNode {
	node.left = a.rrRotation(node.left)
	return a.llRotation(node)
}

func (a *avlTree) rlRotation(node *avlNode) *avlNode {
	node.right = a.llRotation(node.right)
	return a.rrRotation(node)
}
