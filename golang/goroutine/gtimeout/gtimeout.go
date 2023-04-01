package gtimeout

import (
	"context"
	"fmt"
	"time"
)

func GTimeoutF1(in chan struct{}) {

	time.Sleep(1 * time.Second)
	in <- struct{}{}

}

func GTimeoutF2(in chan struct{}) {
	time.Sleep(3 * time.Second)
	in <- struct{}{}
}



func GTimeoutDoCtxTime() {
	tm, ctx := context.WithTimeout(context.Background(), 2*time.Second)

	defer ctx()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("3 s")
	case <-time.After(4 * time.Second):
		fmt.Println("4 s")
	case <-tm.Done():
		fmt.Println("Done")
	}
}
