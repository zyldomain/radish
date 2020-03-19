package channel

import (
	"golang.org/x/sys/unix"
	"radish/channel/util"
	"sync"
)

type EpollEventLoop struct {
	selector Selector
	tasks    *util.TaskList
	id       int64
	stop     bool
	running  bool

	lock sync.RWMutex
}

func NewEpollEventLoop() (*EpollEventLoop, error) {
	s, err := OpenEpollSelector()
	if err != nil {
		return nil, err
	}
	return &EpollEventLoop{
		selector: s,
		tasks:    util.NewTaskList(),
		stop:     true,
		running:  false,
	}, err
}

func (e *EpollEventLoop) StartWork() {
	e.running = true
	e.stop = false

	if e.selector == nil {
		s, err := OpenEpollSelector()
		if err != nil {
			panic(err)
		}

		e.selector = s
	}

	go func() {
		for e.running {

			e.runAllTasks()

			//t := time.Now().Add(2 * time.Second)
			//_, _ := unix.TimeToTimespec(t)
			keys, err := e.selector.Select()
			if err != nil {
				e.reBuildSelector()
			}
			e.processKeys(keys)
		}
	}()
}

func (e *EpollEventLoop) runAllTasks() {
	e.lock.Lock()
	tasks := e.tasks

	e.tasks = util.NewTaskList()

	e.lock.Unlock()

	for _, task := range tasks.Iterator() {
		task.Run()
	}
}

func (e *EpollEventLoop) AddTask(task *util.Task) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.tasks.Add(task)
}

func (e *EpollEventLoop) processKeys(keys []Key) {
	for _, key := range keys {
		if key.Filter == unix.EVFILT_READ {
			list := util.NewArrayList()

			key.Channel.Unsafe().Read(list)

			for _, o := range list.Iterator() {
				key.Channel.Read(o)
			}
		}

		if key.Filter == unix.EVFILT_WRITE {
			//TODO
		}

	}
}

func (e *EpollEventLoop) Register(channel Channel, interests []int16) {

	channel.SetEventLoop(e)
	doRegister := func() {

		unix.SetNonblock(channel.FD(), true)
		for _, filter := range interests {
			e.selector.AddInterests(channel, filter)
		}
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

func (e *EpollEventLoop) reBuildSelector() {

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
