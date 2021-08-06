package racer

import (
	"time"
	"net/http"
	"fmt"
)

// v1
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	startA := time.Now()
// 	http.Get(url)
// 	return time.Since(startA)
// }

// v2
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
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