package cancel

import (
	"context"
	"fmt"
	"time"
)

// 正常终止
func cancelA() {
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

// 父级终止 子级不终止    父级终止 子级终止
func cancelA1() {
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

// 子级终止 父级终止     子级终止结束，父级继续运行，任务完成终止结束
func cancelA2() {
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

// 子级终止 父级不终止   子级终止结束，父级不受影响
func cancelA3() {
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

// 验证a2 是否跟在父级函数创建子函数有关，验证结果，无关系，父级终止，子级即终止
func cancelA4() {
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
