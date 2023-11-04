package data

import "sync"

var (
	wg   *sync.WaitGroup
	once sync.Once
)

func GetWaitGroup() *sync.WaitGroup {
	once.Do(func() {
		wg = &sync.WaitGroup{}
	})
	return wg
}
