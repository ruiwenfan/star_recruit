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

// 参考了网上一些博客，可以直接用go提供的once函数
func GetSingleton2() *Singleton {
	once.Do(func() {
		one = &Singleton{}
	})
	return one
}
