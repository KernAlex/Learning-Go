package testProj

type Compare func(interface{}, interface{}) (int, error)

func AgreaterThanB(f Compare, A interface{}, B interface{}) (int, error){
	return f(A, B)
}

type bNode struct {
	left *bNode
	right *bNode
	value interface{}
}

type Pqueue struct {
	f Compare
	root bNode
	size int
}
/*
Creates New Empty priority queue
 */
func NewPq(f Compare) Pqueue {
	p := Pqueue{f,bNode{nil, nil, nil}, 0}
	return p
}
/*
needs a helper function
for inserting new elements into the pq
 */
func (p *Pqueue) InsertPQ(v interface{}){
	if p.root.left == nil {
		p.root.left = &bNode{nil, nil, v}
		p.size += 1
		return
	}

}
func helper(f Compare, size int, parent *bNode) {

}

/*
swaps for swimming
 */
func swap(left bool, grandParent *bNode, parent *bNode, child *bNode) {
	if left {
		grandParent.left = child
		parent.left, child.left = child.left, parent.left
		parent.right, child.right = child.right, parent.right
	} else {
		grandParent.right = child
		parent.left, child.left = child.left, parent.left
		parent.right, child.right = child.right, parent.right
	}
}
func (p *Pqueue) SizeOf() int {
	return p.size
}


