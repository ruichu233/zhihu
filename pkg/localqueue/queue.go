package localqueue

import (
	"context"
	"log"
	"sync"
	"time"
)

type QueueMode int

const (
	RealTimeMode QueueMode = iota // 实时处理模式
	BatchMode                     // 批处理模式
)

type BatchQueue struct {
	ch        chan interface{}
	batchSize int
	timeout   time.Duration
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	mode      QueueMode
}

func NewBatchQueue(size, batchSize int, timeout time.Duration, mode QueueMode) *BatchQueue {
	ctx, cancel := context.WithCancel(context.Background())
	return &BatchQueue{
		ch:        make(chan interface{}, size),
		batchSize: batchSize,
		timeout:   timeout,
		ctx:       ctx,
		cancel:    cancel,
		mode:      mode,
	}
}

func (q *BatchQueue) Push(data interface{}) {
	select {
	case q.ch <- data:
	case <-q.ctx.Done():
		return
	}
}

func (q *BatchQueue) Run(workers int, handler func(batch []interface{}) error) {
	log.Printf("Starting %d workers for batch queue", workers)
	defer log.Printf("Stopping %d workers for batch queue", workers)
	for i := 0; i < workers; i++ {
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			if q.mode == RealTimeMode {
				// 实时处理模式
				for {
					select {
					case data, ok := <-q.ch:
						if !ok {
							return
						}
						_ = handler([]interface{}{data})
					case <-q.ctx.Done():
						return
					}
				}
			} else {
				// 批处理模式
				batch := make([]interface{}, 0, q.batchSize)
				timer := time.NewTimer(q.timeout)
				defer timer.Stop()
				for {
					select {
					case data, ok := <-q.ch:
						if !ok {
							if len(batch) > 0 {
								_ = handler(batch)
							}
							return
						}
						batch = append(batch, data)
						if len(batch) >= q.batchSize {
							if err := handler(batch); err != nil {
								// 这里可以添加错误处理逻辑
							}
							batch = batch[:0]
							timer.Reset(q.timeout)
						}
					case <-timer.C:
						if len(batch) > 0 {
							if err := handler(batch); err != nil {
								// 这里可以添加错误处理逻辑
							}
							batch = batch[:0]
						}
						timer.Reset(q.timeout)
					case <-q.ctx.Done():
						if len(batch) > 0 {
							_ = handler(batch)
						}
						return
					}
				}
			}

		}()
	}
}

func (q *BatchQueue) Stop() {
	q.cancel()
	q.wg.Wait()
	close(q.ch)
}
