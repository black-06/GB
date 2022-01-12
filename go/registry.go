package gb

import (
	"fmt"
	"sync"
)

var (
	gbs   = map[string]GB{}
	gbsMu sync.RWMutex
)

func RegisterGB(name string, gb GB) {
	gbsMu.Lock()
	if gb == nil {
		panic("RegisterGB: register gb failed")
	}
	gbs[name] = gb
	gbsMu.Unlock()
}

func GetGB(name string) (GB, error) {
	gbsMu.RLock()
	gb, exist := gbs[name]
	gbsMu.RUnlock()
	if !exist {
		return nil, fmt.Errorf("RegisterGB: unregistered gb name \"%+v\"", name)
	}
	return gb, nil
}
