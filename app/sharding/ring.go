package sharding

type Ring struct{
	  nodes []uint32
	  nodeMap map[uint32]string
}


func NewRing() *Ring {
    return &Ring{
        nodes:   make([]uint32, 0),
        nodeMap: make(map[uint32]string),
    }
}

func(ring *Ring) addNode(uint32 hash, string node) {
	 ring.nodes = append(ring.nodes, hash)
	 ring.nodeMap[hash] = node
}

func(ring *Ring) removeNode(uint32 hash) {
	 delete(ring.nodeMap, hash)
	 for i, h := range ring.nodes {
		 if h == hash {
			 ring.nodes = append(ring.nodes[:i], ring.nodes[i+1:]...)
			 break
		 }
	 }
}

func(ring *Ring) getNode(uint32 hash) string {
	 if len(ring.nodes) == 0 {
		 return ""
	 }

	 for _, h := range ring.nodes {
		 if h >= hash {
			 return ring.nodeMap[h]
		 }
	 }
	
	 return ring.nodeMap[ring.nodes[0]]
}	