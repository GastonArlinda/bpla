package storage

import "sync"

type Storage interface{
	Read() Metrics
	Write(met Metrics)
}

type MetricsStorage struct {
	metrics Metrics
	mtx sync.Mutex
}

func NewMetricsStorage() Storage {
	return &MetricsStorage{}
}

func (m *MetricsStorage) Read() Metrics {
	return m.metrics
}

func (m *MetricsStorage) Write(met Metrics) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.metrics = met
}
