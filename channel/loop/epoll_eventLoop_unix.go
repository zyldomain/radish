//+build  linux darwin netbsd freebsd openbsd dragonfly

package loop

import (
	"fmt"
	"radish/channel/iface"
	"radish/channel/selector"
	"radish/channel/util"
	"sync"
)

type EpollEventLoop struct {
	selector iface.Selector
	tasks    *util.TaskList
	ttasks   *util.TaskList
	objList  *util.ArrayList
	id       int64
	stop     bool
	running  bool

	lock sync.RWMutex
}

func NewEpollEventLoop(id int64) (*EpollEventLoop, error) {
	s, err := selector.OpenSelector()
	if err != nil {
		return nil, err
	}
	return &EpollEventLoop{
		selector: s,
		tasks:    util.NewTaskList(),
		ttasks:   util.NewTaskList(),
		objList:  util.NewArrayList(),
		stop:     true,
		running:  false,
		id:       id,
	}, err
}

func (e *EpollEventLoop) StartWork() {
	e.running = true
	e.stop = false

	if e.selector == nil {
		s, err := selector.OpenSelector()
		if err != nil {
			panic(err)
		}

		e.selector = s
	}

	go func() {
		for e.running {
			e.runAllTasks()
			keys, err := e.selector.SelectWithTimeout(util.MillionSecond / 10)
			if err != nil {
				//e.reBuildSelector()
				fmt.Println("1异常断开", e.id)
				fmt.Println(err)
				fmt.Println("2异常断开")
			}
			e.processKeys(keys)
		}
	}()
}

func (e *EpollEventLoop) runAllTasks() {
	e.lock.Lock()
	e.ttasks, e.tasks = e.tasks, e.ttasks
	e.tasks.RemoveAll()
	e.lock.Unlock()
	for _, task := range e.ttasks.Iterator() {
		task.Run()
	}
	e.ttasks.RemoveAll()
}

func (e *EpollEventLoop) AddTask(task *util.Task) {
	e.lock.Lock()

	defer e.lock.Unlock()
	e.tasks.Add(task)
}

func (e *EpollEventLoop) Register(channel iface.Channel) {

	channel.SetEventLoop(e)
	doRegister := func() {
		channel.SetNonBolcking()
		e.selector.AddRead(channel)
		channel.SetActive()
		channel.Pipeline().ChannelActive(channel)
	}

	if e.InEventLoop() {

		doRegister()
	} else {
		if !e.running {
			e.StartWork()
		}
		e.AddTask(util.NewTask(doRegister))
	}

}

func (e *EpollEventLoop) InEventLoop() bool {
	return false
}

func (e *EpollEventLoop) ID() int64 {
	return e.id
}

func (e *EpollEventLoop) Shutdown() {
	e.stop = true
	e.running = false
	//TODO释放selector
	e.selector = nil
}
func (e *EpollEventLoop) AddPackage(ch iface.Channel, pkg *iface.Pkg) {}

func (e *EpollEventLoop) RemoveChannel(ch iface.Channel) {}
