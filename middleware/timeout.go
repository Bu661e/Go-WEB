package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	// 使用 context.WithTimeout 为当前上下文设置一个超时时间（3秒）。
	// 如果超时时间到达，ctx.Done() 通道会被关闭，表示请求超时。
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	r.WithContext(ctx)
	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return
	// channel结束后，会发送一个零值，并立即处理这个分支
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Request Timeout"))
		return
	}
	ctx.Done()

}
