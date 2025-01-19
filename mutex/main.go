package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	enableFeatureKeiba *bool
	mutex              = &sync.RWMutex{}
	errInvalidValueSet = errors.New("entity: invalid value set")
)

func EnableFeatureKeiba() bool {
	mutex.RLock()
	defer mutex.RUnlock()
	return *enableFeatureKeiba
}

func SetEnableFeatureKeiba(i int, value bool) error {
	mutex.Lock()
	defer mutex.Unlock()
	if enableFeatureKeiba != nil {
		return errInvalidValueSet
	}
	enableFeatureKeiba = &value
	return nil
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			if err := SetEnableFeatureKeiba(i, true); err != nil {
				return
			}
		}()
		go func() {
			fmt.Println(i, EnableFeatureKeiba())
		}()
	}
	time.Sleep(10 * time.Second)
}
