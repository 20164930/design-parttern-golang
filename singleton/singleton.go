package singleton

import "sync"

// 懒汉式加载
type singleton struct {
}

var instance *singleton
var mutex sync.Mutex

// 非并发安全
//func GetInstance() *singleton {
//	if instance == nil {
//		instance = &singleton{}
//	}
//	return instance
//}

// 双重判断，优点是对象一旦被创建出来，多线程下不会阻塞
func GetInstance() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}


// sync.Once初始化，能保证只在首次调用时执行一次，内部实现也是双重锁
// 通过done的状态来判断是否执行过
var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
