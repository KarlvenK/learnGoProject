package sync_context

import (
	"context"
	"fmt"
	"time"
)

func WithCancel() {
	reqTask := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop", name, ctx.Err())
				return
			default:
				fmt.Println(name, "send request")
				time.Sleep(1 * time.Second)
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "worker1")
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)

	ctx, cancel = context.WithCancel(context.Background())

	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func WithValue() {
	type Option struct {
		Interval time.Duration
	}

	reqTask := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop", name, ctx.Err())
				return
			default:
				fmt.Println(name, "send request")
				op := ctx.Value("option").(*Option)
				time.Sleep(op.Interval * time.Second)
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "option", &Option{Interval: 1})

	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func WithTimeOut() {
	reqTask := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop", name, ctx.Err())
				return
			default:
				fmt.Println(name, "send request")
				time.Sleep(1 * time.Second)
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(4 * time.Second)
	fmt.Println("before cancel, goroutines ended before i call cancel()")
	cancel()
	time.Sleep(1 * time.Second)

}

func WithDeadline() {
	reqTask := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop", name, ctx.Err())
				return
			default:
				fmt.Println(name, "send request")
				time.Sleep(1 * time.Second)
			}
		}
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel, goroutines ended before i call cancel()")
	cancel()
	time.Sleep(1 * time.Second)
}
