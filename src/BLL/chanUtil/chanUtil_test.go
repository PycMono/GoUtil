package chanUtil

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	firstSuccessCh_init  = make(chan bool)
	secondSuccessCh_init = make(chan bool)
	threeSuccessCh_init  = make(chan bool)
	fourSuccessCh_init   = make(chan bool)
	wg                   sync.WaitGroup
)

func TestNewChanUtil(t *testing.T) {
	RegisterInitSuccess("first", firstSuccessCh_init)
	RegisterInitSuccess("second", secondSuccessCh_init)
	RegisterInitSuccess("three", threeSuccessCh_init)
	RegisterInitSuccess("four", fourSuccessCh_init)

	go func() {
		<-firstSuccessCh_init
		fmt.Println("first")
	}()

	go func() {
		<-secondSuccessCh_init
		fmt.Println("second")
	}()

	go func() {
		<-threeSuccessCh_init
		fmt.Println("three")
	}()

	go func() {
		<-fourSuccessCh_init
		fmt.Println("four")
	}()

	go func() {
		time.Sleep(time.Minute * 1)
		Notify()
	}()

	time.Sleep(time.Minute * 20)
}
