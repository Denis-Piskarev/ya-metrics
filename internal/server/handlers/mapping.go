package handlers

import (
	"context"
	"fmt"
	"strconv"

	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
)

// Mapping metric write
var metricWrite map[string]func(ctx context.Context, metric yametrics.Metric, name, value string) error = map[string]func(ctx context.Context, metric yametrics.Metric, name, value string) error{
	"counter": func(ctx context.Context, metric yametrics.Metric, name, value string) error {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}

		_, err = metric.WriteCounter(ctx, name, val)
		if err != nil {
			return err
		}

		return nil
	},
	"gauge": func(ctx context.Context, metric yametrics.Metric, name, value string) error {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		_, err = metric.WriteGauge(ctx, name, val)
		if err != nil {
			return err
		}

		return nil
	},
}

// Mapping metric get
var metricGet map[string]func(ctx context.Context, metric yametrics.Metric, name string) (string, error) = map[string]func(ctx context.Context, metric yametrics.Metric, name string) (string, error){
	"counter": func(ctx context.Context, metric yametrics.Metric, name string) (string, error) {
		val, err := metric.GetCounter(ctx, name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%v", val), nil
	},
	"gauge": func(ctx context.Context, metric yametrics.Metric, name string) (string, error) {
		val, err := metric.GetGauge(ctx, name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%v", val), nil
	},
}
