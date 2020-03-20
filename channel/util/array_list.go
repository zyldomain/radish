package util

import "errors"

type ArrayList struct {
	elems []interface{}
	size  int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elems: make([]interface{}, 16),
		size:  0,
	}
}

func (a *ArrayList) Add(obj interface{}) {
	a.elems[a.size] = obj
	a.size++
	if a.size >= len(a.elems)/2 {
		tmp := make([]interface{}, 2*len(a.elems))
		copy(tmp, a.elems)
		a.elems = tmp
	}
}

func (a *ArrayList) Remove(index int) {
	if index >= a.size {
		panic(errors.New("越界"))
	}

	for i := index; i < a.size; i++ {
		a.elems[i] = a.elems[i+1]
	}

	a.size--
}

func (a *ArrayList) Get(index int) interface{} {
	if index >= a.size {
		panic(errors.New("越界"))
	}

	return a.elems[index]
}

func (a *ArrayList) Iterator() []interface{} {
	return a.elems[:a.size]
}
func (a *ArrayList) Size() int {
	return a.size
}
