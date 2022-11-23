package star

import "sync"

type Singleton struct {
}

var one *Singleton
var mu sync.Mutex
var once sync.Once

func GetSingleton() *Singleton {
	// 双重检验
	if one == nil {
		mu.Lock()
		if one == nil {
			one = &Singleton{}
		}
		mu.Unlock()
	}
	return one
}

func GetSingleton2() *Singleton {
	once.Do(func() {
		one = &Singleton{}
	})
	return one
}
