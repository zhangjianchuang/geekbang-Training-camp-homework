package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

var srv = &http.Server{Addr: ":8080"}

func starServer() error {
	fmt.Println("http server starting...")
	return srv.ListenAndServe()
}

func preDestroy() {
	fmt.Println("destroying...")
}

func listenFinishedEvent(ctx context.Context, cancel context.CancelFunc) error {
	// 监听系统kill信号量
	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)
	// 循环监听是为了当系统kill时能够做销毁前的处理
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done, do pre destroy...")
			preDestroy()
			fmt.Println("http server stopping...")
			return srv.Shutdown(ctx)
		case <-chanel:
			cancel()
			fmt.Println("killed, do cancel...")
		}
	}
}

func main() {
	// 获取一个上下文，注册cancel方法
	// 使用 errgroup 注册goroutine取消方法
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	group.Go(starServer)
	group.Go(func() error {
		return listenFinishedEvent(ctx, cancel)
	})

	// 等待所有goroutine完成，如果有异常将异常打印
	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all done!")

}
