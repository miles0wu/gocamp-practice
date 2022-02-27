package fifth

import (
	"errors"
	"sync"
	"time"
)

type Breaker struct {
	sync.RWMutex
	status            bool
	metrics           []*Metric
	minReqThreshold   int64
	failRateThreshold float64
	windowSize        int
	lastTimestamp     time.Time
	retryDuration     time.Duration
}

func NewBreaker(
	windowSize int,
	minReqThreshold int64,
	failRateThreshold float64,
	retryDuration time.Duration,
) *Breaker {
	return &Breaker{
		windowSize:        windowSize,
		metrics:           make([]*Metric, 0, windowSize),
		minReqThreshold:   minReqThreshold,
		failRateThreshold: failRateThreshold,
		retryDuration:     retryDuration,
	}
}

func (b *Breaker) Launch() {
	go func() {
		for {
			b.AddMetric()
			time.Sleep(time.Microsecond * 100)
		}
	}()
}

func (b *Breaker) AddMetric() {
	if len(b.metrics) >= b.windowSize {
		b.metrics = b.metrics[1:]
	}

	b.metrics = append(b.metrics, NewMetric())
}

func (b *Breaker) RecordReqStatus(status bool) error {
	if len(b.metrics) == 0 {
		return errors.New("breaker did not launch")
	}

	m := b.metrics[len(b.metrics)-1]
	m.Record(status)

	return nil
}

func (b *Breaker) reqCount() int64 {
	cnt := int64(0)
	for _, m := range b.metrics {
		cnt += m.NumReq
	}
	return cnt
}

func (b *Breaker) failCount() int64 {
	cnt := int64(0)
	for _, m := range b.metrics {
		cnt += m.NumFail
	}
	return cnt
}

func (b *Breaker) confirmLoad() bool {
	b.RLock()
	defer b.RUnlock()
	numReq := b.reqCount()
	numFailed := b.failCount()
	if float64(numFailed)/float64(numReq) > b.failRateThreshold && numReq > b.minReqThreshold {
		return true
	}

	return false
}

func (b *Breaker) Monitor() {
	go func() {
		for {
			if b.status {
				if time.Since(b.lastTimestamp) > b.retryDuration {
					b.Lock()
					b.status = false
					b.Unlock()
				}
				continue
			}
			if b.confirmLoad() {
				b.Lock()
				b.status = true
				b.lastTimestamp = time.Now()
				b.Unlock()
			}
		}
	}()
}
