package mapbench

import "sync"

type MyMap struct {
	sync.Mutex
	m map[int]int
}

var myMap *MyMap
var syncMap *sync.Map

func init() {
	myMap = &MyMap{
		m: make(map[int]int, 100),
	}

	syncMap = &sync.Map{}
}

func builtinMapStore(k, v int) {
	myMap.Lock()
	defer myMap.Unlock()
	myMap.m[k] = v
}

func builtinMapLookup(k int) int {
	myMap.Lock()
	defer myMap.Unlock()
	if v, ok := myMap.m[k]; !ok {
		return -1
	} else {
		return v
	}
}

func builtinMapDelete(k int) {
	myMap.Lock()
	defer myMap.Unlock()
	if _, ok := myMap.m[k]; !ok {
		return
	} else {
		delete(myMap.m, k)
	}
}

func syncMapStore(k, v int) {
	syncMap.Store(k, v)
}

func syncMapLookup(k int) int {
	v, ok := syncMap.Load(k)
	if !ok {
		return -1
	}

	return v.(int)
}

func syncMapDelete(k int) {
	syncMap.Delete(k)
}
