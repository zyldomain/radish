package util

import "errors"

type TaskList struct {
	elems []*Task
	size  int
}

func NewTaskList() *TaskList {
	return &TaskList{
		elems: make([]*Task, 16),
		size:  0,
	}
}

func (a *TaskList) Add(task *Task) {
	a.elems[a.size] = task
	a.size++
	if a.size >= len(a.elems)/2 {
		tmp := make([]*Task, 2*len(a.elems))
		copy(tmp, a.elems)
		a.elems = tmp
	}
}

func (a *TaskList) Remove(index int) {
	if index >= a.size {
		panic(errors.New("越界"))
	}

	for i := index; i < a.size; i++ {
		a.elems[i] = a.elems[i+1]
	}

	a.size--
}

func (a *TaskList) Get(index int) *Task {
	if index >= a.size {
		panic(errors.New("越界"))
	}

	return a.elems[index]
}

func (a *TaskList) Iterator() []*Task {
	return a.elems[:a.size]
}
