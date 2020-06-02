package cachedata

// LRU 最少最近算法
type LruCacheList struct {
	limit   int
	HashMap map[int]*DCacheNode
	head    *DCacheNode
	end     *DCacheNode
}

// func Constuctor(capacity int) LruCacheList {
// 	lruList := LruCacheList{limit: capacity}

// }

// func (l *LruCacheList) Get(key int) int {

// }
