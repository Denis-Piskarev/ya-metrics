package usecase

import (
	"go.uber.org/zap"
)

type Metric struct {
	m      MetricInterface // struct implementation interfacep
	logger *zap.SugaredLogger
}

func NewMetric(m MetricInterface, logger *zap.SugaredLogger) *Metric {
	return &Metric{
		m:      m,
		logger: logger,
	}
}
