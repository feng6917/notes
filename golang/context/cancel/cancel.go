package cancel

import (
	"context"
	"fmt"
	"time"
)

func CancelA() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("done!")
				return
			default:
				fmt.Println(time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func CancelA1() {
	var ctx, nctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, _ = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 终止
	cancel()
	time.Sleep(1 * time.Second)
	// 终止
	//ncancel()
	//time.Sleep(1 * time.Second)
}

func CancelA2() {
	var ctx, nctx context.Context
	var cancel, ncancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, ncancel = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
					// 子级终止
					ncancel()
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 父级终止
	cancel()
	time.Sleep(1 * time.Second)
}

func CancelA3() {
	var ctx, nctx context.Context
	var ncancel context.CancelFunc
	ctx, _ = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, ncancel = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
					// 子级终止
					ncancel()
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 父级终止
	// cancel()
	time.Sleep(1 * time.Second)
}

func CancelA4() {
	var ctx, nctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)

	nctx, _ = context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("child done!")
				return
			default:
				fmt.Println("child: ", time.Second)
			}
		}
	}(nctx)
	time.Sleep(5 * time.Second)
	// 终止
	cancel()
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
}
