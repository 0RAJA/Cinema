package dao

import "sync"

//锁
var (
	rwMutexBox     = new(sync.RWMutex)
	rwMutexComment = new(sync.RWMutex)
	rwmutexMessage = new(sync.RWMutex)
	rwmutexMovie   = new(sync.RWMutex)
	rwmutexPlan    = new(sync.RWMutex)
	rwmutexSeat    = new(sync.RWMutex)
	rwmutexTicket  = new(sync.RWMutex)
	rwmutexUser    = new(sync.RWMutex)
)
