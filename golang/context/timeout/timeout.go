package timeout

import (
	"context"
	"fmt"
	"time"
)

func newContextWithTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*3)
}

func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("val: ", i)
		}
	}
}

func deal2(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("val: ", i)
			cancel()
		}
	}
}

func TimeoutA() {
	ctx, cancel := newContextWithTimeOut()
	defer cancel()
	deal(ctx)
}

func TimeoutA1() {
	ctx, cancel := newContextWithTimeOut()
	defer cancel()
	deal2(ctx, cancel)
}
