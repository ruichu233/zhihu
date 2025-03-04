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
	for i := 0; i < workers; i++ {
		q.wg.Add(1)
		go func(workerID int) {
			defer q.wg.Done()
			if q.mode == RealTimeMode {
				// 实时处理模式
				for {
					select {
					case data, ok := <-q.ch:
						if !ok {
							return
						}
						if err := handler([]any{data}); err != nil {
							log.Printf("Worker %d: Error handling message: %v", workerID, err)
						}
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
								if err := handler(batch); err != nil {
									log.Printf("Worker %d: Error handling final batch: %v", workerID, err)
								}
							}
							log.Printf("Worker %d: Stopping batch queue", workerID)
							return
						}
						batch = append(batch, data)
						if len(batch) >= q.batchSize {
							if err := handler(batch); err != nil {
								log.Printf("Worker %d: Error handling batch: %v", workerID, err)
							}
							batch = batch[:0]
							timer.Reset(q.timeout)
						}
					case <-timer.C:
						if len(batch) > 0 {
							if err := handler(batch); err != nil {
								log.Printf("Worker %d: Error handling timeout batch: %v", workerID, err)
							}
							batch = batch[:0]
						}
						timer.Reset(q.timeout)
					case <-q.ctx.Done():
						if len(batch) > 0 {
							if err := handler(batch); err != nil {
								log.Printf("Worker %d: Error handling shutdown batch: %v", workerID, err)
							}
						}
						log.Printf("Worker %d: Stopping due to shutdown", workerID)
						return
					}
				}
			}
		}(i)
	}
}

func (q *BatchQueue) Stop() {
	q.cancel()
	q.wg.Wait()
	close(q.ch)
}
