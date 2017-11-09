package session

import (
	"sync"

	"github.com/wpajqz/linker"
)

const (
	OFFLINE = 0
	ONLINE  = 1
)

var (
	defaultSession = make(map[string]Session)
	mutex          sync.RWMutex
)

type (
	// socket信息、在线状态
	Session struct {
		Address string
		Ctx     *linker.Context
		Status  int // 0:不在线;1:在线
	}
)

func Get(key string) Session {
	mutex.RLock()
	defer mutex.RUnlock()

	return defaultSession[key]
}

func Set(key string, session Session) {
	mutex.Lock()
	defer mutex.Unlock()

	defaultSession[key] = session
}

func IsExist(key string) bool {
	mutex.RLock()
	defer mutex.RUnlock()

	if _, ok := defaultSession[key]; ok {
		return true
	}

	return false
}

func Delete(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(defaultSession, key)
}

func Default() map[string]Session {
	return defaultSession
}
