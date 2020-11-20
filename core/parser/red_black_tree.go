package parser

const (
	RedBlackTreeRed   int = 0
	RedBlackTreeBlack int = 1

	RedBlackTreeNone      int = 0
	RedBlackTreeLeftDone  int = 1
	RedBlackTreeRightDone int = 2
)

type RedBlackTreeNode struct {
	left   *RedBlackTreeNode
	right  *RedBlackTreeNode
	parent *RedBlackTreeNode
	value  interface{}
	color  int
}

type RedBlackTree struct {
	compare func(left, right interface{}) int
	root    *RedBlackTreeNode
	count   int
}

func NewRedBlackTree(compare func(left, right interface{}) int) *RedBlackTree {
	return &RedBlackTree{
		compare: compare,
	}
}

func (t *RedBlackTree) Iterate(on func(interface{})) {
	if t.root == nil {
		return
	}
	cursor := t.root
	state := 0
	for {
		if state == RedBlackTreeNone {
			if cursor.left == nil {
				state = RedBlackTreeLeftDone
			} else {
				cursor = cursor.left
			}
			continue
		}
		if state == RedBlackTreeLeftDone {
			on(cursor.value)
			if cursor.right == nil {
				state = RedBlackTreeRightDone
			} else {
				state = RedBlackTreeNone
				cursor = cursor.right
			}
			continue
		}
		if state == RedBlackTreeRightDone {
			if cursor.parent == nil {
				break
			}
			if cursor.parent.left == cursor {
				state = RedBlackTreeLeftDone
			} else {
				state = RedBlackTreeRightDone
			}
			cursor = cursor.parent
			continue
		}
	}
}

func (t *RedBlackTree) Get(value interface{}) interface{} {
	node := t.getFromNode(t.root, value)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *RedBlackTree) Has(value interface{}) bool {
	return t.getFromNode(t.root, value) != nil
}

func (t *RedBlackTree) Size() int {
	return t.count
}

func (t *RedBlackTree) getFromNode(node *RedBlackTreeNode, value interface{}) *RedBlackTreeNode {
	result := 0
	cursor := node
	for cursor != nil {
		result = t.compare(value, node.value)
		if result == 0 {
			return cursor
		}
		if result < 0 {
			cursor = cursor.left
		} else {
			cursor = cursor.right
		}
	}
	return nil
}

func (t *RedBlackTree) Add(value interface{}) {
	if t.root == nil {
		t.root = &RedBlackTreeNode{
			value: value,
			color: RedBlackTreeBlack,
		}
		t.count++
		return
	}

	t.addToNode(t.root, value)
}

func (t *RedBlackTree) addToNode(node *RedBlackTreeNode, value interface{}) {
	result := t.compare(value, node.value)
	if result < 0 {
		if node.left == nil {
			node.left = &RedBlackTreeNode{
				value:  value,
				parent: node,
			}
			t.count++
			t.fixNodeAdd(node.left)
			return
		}
		t.addToNode(node.left, value)
		return
	}
	if result > 0 {
		if node.right == nil {
			node.right = &RedBlackTreeNode{
				value:  value,
				parent: node,
			}
			t.count++
			t.fixNodeAdd(node.right)
			return
		}
		t.addToNode(node.right, value)
		return
	}
	node.value = value
}

func (t *RedBlackTree) fixNodeAdd(node *RedBlackTreeNode) {
	node.color = RedBlackTreeRed
	for node != nil && node != t.root && node.parent.color != RedBlackTreeRed {
		parent := node.parent
		grandParent := parent.parent
		if parent == grandParent.left {
			uncle := grandParent.right
			if uncle != nil && uncle.color == RedBlackTreeRed {
				uncle.color = RedBlackTreeBlack
				parent.color = RedBlackTreeBlack
				grandParent.color = RedBlackTreeRed
				node = grandParent
			} else {
				if node == parent.right {
					node = parent
					t.leftRotate(parent)
				}
				node.parent.color = RedBlackTreeBlack
				grandParent.color = RedBlackTreeRed
				t.rightRotate(grandParent)
			}
		} else {
			uncle := grandParent.left
			if uncle != nil && uncle.color == RedBlackTreeRed {
				uncle.color = RedBlackTreeBlack
				parent.color = RedBlackTreeBlack
				grandParent.color = RedBlackTreeRed
				node = grandParent
			} else {
				if node == parent.left {
					node = parent
					t.rightRotate(parent)
				}
				node.parent.color = RedBlackTreeBlack
				grandParent.color = RedBlackTreeRed
				t.leftRotate(grandParent)
			}
		}
	}
	t.root.color = RedBlackTreeBlack
}

func (t *RedBlackTree) rightRotate(node *RedBlackTreeNode) {
	if node == nil {
		return
	}
	cursor := node.left

	node.left = cursor.right
	if cursor.right != nil {
		cursor.right.parent = node
	}

	cursor.right = node
	cursor.parent = node.parent
	node.parent = cursor

	if cursor.parent == nil {
		t.root = cursor
	} else if cursor.parent.left == node {
		cursor.parent.left = cursor
	} else {
		cursor.parent.right = cursor
	}
}

func (t *RedBlackTree) leftRotate(node *RedBlackTreeNode) {
	if node == nil {
		return
	}
	cursor := node.right

	node.right = cursor.left
	if cursor.left != nil {
		cursor.left.parent = node
	}

	cursor.left = node
	cursor.parent = node.parent
	node.parent = cursor

	if cursor.parent == nil {
		t.root = cursor
	} else if cursor.parent.right == node {
		cursor.parent.right = cursor
	} else {
		cursor.parent.left = cursor
	}
}

func (t *RedBlackTree) Remove(value interface{}) {
	if t.root == nil {
		return
	}

	cursor := t.root
	result := 0
	for cursor != nil {
		result = t.compare(value, cursor.value)
		if result == 0 {
			break
		}
		if result < 0 {
			cursor = cursor.left
		} else {
			cursor = cursor.right
		}
	}

	if cursor == nil {
		return
	}
	if cursor.left == nil && cursor.right == nil {
		t.removeLeafNode(cursor)
		return
	}
	if cursor.left == nil {
		cursor.value = cursor.right.value
		t.removeLeafNode(cursor.right)
		return
	}
	if cursor.right == nil {
		cursor.value = cursor.left.value
		t.removeLeafNode(cursor.left)
		return
	}
	leftMax := t.getMaxNode(cursor.left)
	cursor.value = leftMax.value
	t.removeLeafNode(leftMax)
}

func (t *RedBlackTree) getMaxNode(node *RedBlackTreeNode) *RedBlackTreeNode {
	cursor := node
	for cursor.right != nil {
		cursor = cursor.right
	}
	if cursor.left != nil {
		cursor.color, cursor.left.color = cursor.left.color, cursor.color
		t.rightRotate(cursor)
	}
	return cursor
}

func (t *RedBlackTree) removeLeafNode(node *RedBlackTreeNode) {
	if node.parent == nil {
		t.root = nil
		t.count = 0
		return
	}
	cursor := node
	for cursor != nil {
		if cursor.color == RedBlackTreeRed {
			cursor.color = RedBlackTreeBlack
			cursor = nil
		} else {
			parent := cursor.parent
			if parent == nil {
				cursor = nil
			} else if cursor == parent.left {
				brother := parent.right
				if brother.color == RedBlackTreeRed {
					parent.color = RedBlackTreeRed
					brother.color = RedBlackTreeBlack
					t.leftRotate(parent)
				} else {
					if brother.right != nil && brother.right.color == RedBlackTreeRed {
						brother.color = parent.color
						parent.color = RedBlackTreeBlack
						brother.right.color = RedBlackTreeBlack
						t.leftRotate(parent)
						cursor = nil
					} else if brother.left != nil && brother.left.color == RedBlackTreeRed {
						brother.color = RedBlackTreeRed
						brother.left.color = RedBlackTreeBlack
						t.rightRotate(brother)
					} else {
						brother.color = RedBlackTreeRed
						cursor = parent
					}
				}
			} else {
				brother := parent.left
				if brother.color == RedBlackTreeRed {
					parent.color = RedBlackTreeRed
					brother.color = RedBlackTreeBlack
					t.rightRotate(parent)
				} else {
					if brother.left != nil && brother.left.color == RedBlackTreeRed {
						brother.color = parent.color
						parent.color = RedBlackTreeBlack
						brother.left.color = RedBlackTreeBlack
						t.rightRotate(parent)
						cursor = nil
					} else if brother.right != nil && brother.right.color == RedBlackTreeRed {
						brother.color = RedBlackTreeRed
						brother.right.color = RedBlackTreeBlack
						t.leftRotate(brother)
					} else {
						brother.color = RedBlackTreeRed
						cursor = parent
					}
				}
			}
		}
	}
	if node == node.parent.left {
		node.parent.left = nil
	} else {
		node.parent.right = nil
	}
	t.count--
}

func (t *RedBlackTree) Clear() {
	t.root = nil
	t.count = 0
}
