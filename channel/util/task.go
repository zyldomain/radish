package util

type Task struct {
	run  func()
}


func NewTask(r func()) *Task{
	return &Task{run:r}
}

func (t *Task)Run(){
	t.run()
}