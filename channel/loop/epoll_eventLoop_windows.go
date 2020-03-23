// +build windows

package loop

import (
	"errors"
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

	lock *sync.RWMutex
	tlock sync.Mutex

	channelPkg map[iface.Channel]chan *iface.Pkg
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
		channelPkg:make(map[iface.Channel]chan *iface.Pkg),
		lock:&sync.RWMutex{},
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
			e.lock.RLock()
			for c, ch := range e.channelPkg{
				stop := false
				for !stop{
					select {
					case p := <-ch:
						if p.Event == iface.READ ||  p.Event == iface.CONNECTED {
							c.Read(p.Data)
						}else if p.Event == iface.CLOSED {
							subc, ok := p.Data.(iface.Channel)
							if !ok {
								panic("wrong type")
							}
							e.RemoveChannel(subc)
						}else if p.Event == iface.WRITE {
							crw, ok := c.(iface.ChannelRW)
							if !ok {
								panic("wrong type")
							}
							crw.AddWriteMsg(p)
						}
					default:
						stop = true
						break
					}
				}
			}
			e.lock.RUnlock()

		}
	}()
}

func (e *EpollEventLoop) runAllTasks() {
	e.tlock.Lock()
	e.ttasks, e.tasks = e.tasks, e.ttasks
	e.tasks.RemoveAll()
	e.tlock.Unlock()
	for _, task := range e.ttasks.Iterator() {
		task.Run()
	}
	e.ttasks.RemoveAll()
}

func (e *EpollEventLoop) AddTask(task *util.Task) {
	e.tlock.Lock()

	defer e.tlock.Unlock()
	e.tasks.Add(task)
}

func (e *EpollEventLoop) Register(channel iface.Channel) {

	channel.SetEventLoop(e)
	e.lock.Lock()

	e.channelPkg[channel] = make(chan *iface.Pkg,1000)


	c, ok := channel.(iface.ChannelRW)

	if !ok {
		panic(errors.New("wrong type"))
	}
	e.lock.Unlock()
	go c.WriteLoop()
	go c.ReadLoop()

	doRegister := func() {
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

func (e *EpollEventLoop) AddPackage(c iface.Channel, pkg *iface.Pkg){
	e.lock.RLock()
	ch, ok := e.channelPkg[c]

	if !ok{
		panic("unkown error")
	}
	e.lock.RUnlock()


	ch <- pkg
}
func (e *EpollEventLoop)RemoveChannel(ch iface.Channel){
	e.lock.Lock()
	delete(e.channelPkg,ch)

	e.lock.Unlock()
}
