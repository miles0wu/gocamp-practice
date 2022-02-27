package fifth

import (
	"sync/atomic"
	"time"
)

type Metric struct {
	NumReq    int64
	NumFail   int64
	Timestamp time.Time
}

func NewMetric() *Metric {
	return &Metric{
		Timestamp: time.Now(),
	}
}

func (m *Metric) Record(status bool) {
	if !status {
		atomic.AddInt64(&m.NumFail, 1)
	}

	atomic.AddInt64(&m.NumReq, 1)
}
