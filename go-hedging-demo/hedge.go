package main

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"
)

// HedgedTransport 实现了 http.RoundTripper 接口
type HedgedTransport struct {
	Transport   http.RoundTripper // 底层真正的 Transport
	MaxAttempts int               // 最大并发请求数（包括最初的1次）
	HedgeDelay  time.Duration     // 触发对冲的延迟时间
}

func (ht *HedgedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 如果没有设置，使用默认行为
	transport := ht.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	attempts := ht.MaxAttempts
	if attempts <= 0 {
		attempts = 1
	}

	// 使用带有取消功能的 context 控制整个对冲生命周期
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	// 结果通道，用于接收第一个成功的响应或错误
	type result struct {
		resp *http.Response
		err  error
	}
	resCh := make(chan result, attempts)
	var wg sync.WaitGroup

	// 启动一个请求的闭包函数
	doRequest := func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 克隆请求，防止并发修改
			cloneReq := req.Clone(ctx)
			resp, err := transport.RoundTrip(cloneReq)

			// 只有当请求不是因为 context 取消而失败时，才尝试写入结果
			if !errors.Is(err, context.Canceled) {
				select {
				case resCh <- result{resp: resp, err: err}:
				default:
					// 通道已满或已不再需要，直接丢弃（如果 resp 不为空，需要关闭 Body 以防泄露）
					if resp != nil && resp.Body != nil {
						resp.Body.Close()
					}
				}
			}
		}()
	}

	// 1. 发起第一个请求
	doRequest()

	// 2. 控制对冲的定时器和尝试次数
	timer := time.NewTimer(ht.HedgeDelay)
	defer timer.Stop()

	errs := make([]error, 0, attempts)
	requestsSent := 1

	for {
		select {
		case res := <-resCh:
			// 收到结果
			if res.err == nil {
				// 成功！立即取消其他还在飞行的请求
				cancel()
				// 等待后台 goroutine 清理完成 (可选，这里为了简单不阻塞)
				return res.resp, nil
			}
			// 如果这个请求失败了，记录错误
			errs = append(errs, res.err)
			// 如果所有发出的请求都失败了，且已经达到最大尝试次数，返回错误
			if len(errs) == attempts {
				return nil, errors.Join(errs...)
			}

			// 如果一个请求失败了，且还没达到最大尝试次数，我们不应该死等 Timer，
			// 而应该立刻触发下一个对冲请求（这里为了简化逻辑，依然依赖下一次 Timer 或失败循环）
			// 实际生产级实现可以在这里直接触发 doRequest()

		case <-timer.C:
			// 对冲延迟到达
			if requestsSent < attempts {
				// 触发对冲请求
				doRequest()
				requestsSent++
				// 重置定时器，准备下一次可能的对冲
				timer.Reset(ht.HedgeDelay)
			}

		case <-ctx.Done():
			// 整个请求超时或被调用方取消
			return nil, ctx.Err()
		}
	}
}
