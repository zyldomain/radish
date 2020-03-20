package epoll

import (
	"fmt"
	"golang.org/x/sys/unix"
	"radish/channel/iface"
	"radish/channel/selector"
	"radish/channel/util"
	"sync"
	"time"
)

const (
	MillionSecond = int64(time.Nanosecond) * 1000000
	Second        = MillionSecond * 1000
)

type EpollEventLoop struct {
	selector iface.Selector
	tasks    *util.TaskList
	id       int64
	stop     bool
	running  bool

	lock sync.RWMutex
}

func NewEpollEventLoop(id int64) (*EpollEventLoop, error) {
	s, err := selector.OpenEpollSelector()
	if err != nil {
		return nil, err
	}
	return &EpollEventLoop{
		selector: s,
		tasks:    util.NewTaskList(),
		stop:     true,
		running:  false,
		id:       id,
	}, err
}

func (e *EpollEventLoop) StartWork() {
	e.running = true
	e.stop = false

	if e.selector == nil {
		s, err := selector.OpenEpollSelector()
		if err != nil {
			panic(err)
		}

		e.selector = s
	}

	go func() {
		for e.running {
			e.runAllTasks()

			tt := unix.NsecToTimespec(MillionSecond * 100)
			keys, err := e.selector.SelectWithTimeout(&tt)
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
	tasks := e.tasks

	e.tasks = util.NewTaskList()

	for _, task := range tasks.Iterator() {
		task.Run()
	}

	e.lock.Unlock()
}

func (e *EpollEventLoop) AddTask(task *util.Task) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.tasks.Add(task)
}

func (e *EpollEventLoop) processKeys(keys []iface.Key) {
	for _, key := range keys {
		if key.Flags&unix.EV_ERROR != 0 || key.Flags&unix.EV_EOF != 0 {
			unix.Close(key.Channel.FD())
			continue
		}
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

func (e *EpollEventLoop) Register(channel iface.Channel, interests []int16) {

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
