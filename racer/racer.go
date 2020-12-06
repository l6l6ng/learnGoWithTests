package racer

import (
	"fmt"
	"net/http"
	"time"
)

var ten = 10 * time.Second

func Racer(a, b string, timeout time.Duration) (winner string, error error) {
	//startA := time.Now()
	//http.Get(a)
	//aDuration := time.Since(startA)
	//
	//startB := time.Now()
	//http.Get(b)
	//bDuration := time.Since(startB)

	//aDuration := measureResponseTime(a)
	//bDuration := measureResponseTime(a)
	//
	//if aDuration < bDuration {
	//	return a
	//}
	//
	//return b

	return ConfigurableRacer(a, b, ten)
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
