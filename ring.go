package bestring

import (
	"fmt"
)

const (
	defaultVodeNumber = 3
)

type ServerNode struct {
	Addr string
}

type HashRing interface {
	AddNode(addr string) (totalNodes int)
	DeleteNode(addr string) (totalNodes int)
	GetNode(sec []byte) (addr string, err error)
}

type bestRing struct {
	vNodeNumber int
	realNodes   map[uint32]*ServerNode
	nodeTree    *avlTree
}

func NewBestRing(vnodes int) HashRing {
	if vnodes <= 0 {
		vnodes = defaultVodeNumber
	}

	ring := &bestRing{
		vNodeNumber: vnodes,
		realNodes:   make(map[uint32]*ServerNode),
		nodeTree:    newAVLTree(),
	}

	return ring
}

func (b *bestRing) AddNode(addr string) (totalNodes int) {
	nodes := &ServerNode{addr}

	for i := 0; i < b.vNodeNumber; i++ {
		cur := checkSum([]byte(fmt.Sprintf("Addr:%s:Number:%d", addr, i)))

		b.realNodes[cur] = nodes
		b.nodeTree.add(cur)
	}

	return len(b.realNodes)
}

func (b *bestRing) DeleteNode(addr string) (totalNodes int) {
	return
}

func (b *bestRing) GetNode(sec []byte) (addr string, err error) {
	if len(b.realNodes) == 0 {
		err = fmt.Errorf("empty ring")
		return
	}

	cur := checkSum(sec)
	return b.realNodes[b.nodeTree.findLatestLeft(cur)].Addr, nil
}
