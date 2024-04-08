package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
用signal.Notify和http.Server的Shutdown结合起来实现服务器的优雅终止
Handler其会在2s后返回 hello。
创建一个 http.Server 实例，指定端口与 Handler。
声明一个 processed chan，其用来保证服务优雅的终止后再退出主 goroutine。
新启一个 goroutine，其会监听 os.Interrupt 信号，一旦服务被中断即调用服务的 Shutdown 方法，
确保活跃连接的正常返回（本代码使用的 Context 超时时间为 3s，大于服务 Handler 的处理时间，所以不会超时）。
处理完成后，关闭 processed 通道，最后主 goroutine 退出。
*/
var addr = flag.String("server addr", ":8080", "server address")

func main() {
	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "hello")
	})

	// server
	srv := http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// make sure idle connections returned
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		/*Shutdown() 方法会优雅地关闭服务器，等待所有正在进行的连接处理完成后再关闭。
		如果在规定的超时时间内没有成功关闭服务器，会返回一个错误。
		*/
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed, err: %v\n", err)
		}
		log.Println("server gracefully shutdown")

		close(processed)
	}()

	//serve
	err := srv.ListenAndServe()
	if http.ErrServerClosed != err {
		log.Fatalf("Server not gracefully shutdown, err: %v\n", err)
	}

	// waiting for goroutine above processed
	<-processed
}
