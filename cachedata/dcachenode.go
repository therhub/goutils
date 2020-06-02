package cachedata

type DCacheNode struct {
	Key   int
	Value int
	Prev  *DCacheNode
	Next  *DCacheNode
}
