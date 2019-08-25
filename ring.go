package bestring

import (
	"fmt"
	"sync"
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

	nodeTree *avlTree
	locker   *sync.RWMutex
}

func NewBestRing(vnodes int) HashRing {
	if vnodes <= 0 {
		vnodes = defaultVodeNumber
	}

	ring := &bestRing{
		vNodeNumber: vnodes,
		realNodes:   make(map[uint32]*ServerNode),
		nodeTree:    newAVLTree(),
		locker:      new(sync.RWMutex),
	}

	return ring
}

func (b *bestRing) AddNode(addr string) (totalNodes int) {
	nodes := &ServerNode{addr}

	for i := 0; i < b.vNodeNumber; i++ {
		cur := checkSum([]byte(fmt.Sprintf("Addr:%s:Number:%d", addr, i)))

		b.locker.Lock()
		b.realNodes[cur] = nodes
		b.nodeTree.add(cur)
		b.locker.Unlock()
	}

	return b.nodeTree.total
}

func (b *bestRing) DeleteNode(addr string) (totalNodes int) {
	for i := 0; i < b.vNodeNumber; i++ {
		cur := checkSum([]byte(fmt.Sprintf("Addr:%s:Number:%d", addr, i)))

		b.locker.Lock()
		delete(b.realNodes, cur)
		totalNodes = b.nodeTree.delete(cur)
		b.locker.Unlock()
	}
	return
}

func (b *bestRing) GetNode(sec []byte) (addr string, err error) {
	if len(b.realNodes) == 0 {
		err = fmt.Errorf("empty ring")
		return
	}
	cur := checkSum(sec)

	b.locker.RLock()
	addr = b.realNodes[b.nodeTree.findLatestLeft(cur)].Addr
	b.locker.RUnlock()
	return
}
