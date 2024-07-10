package memyandex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (m *MemStatsYaSt) SendAllMetricsToServer(ctx context.Context, addr string) error {
	metrics, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Sending request with compressed data
	client := http.Client{Timeout: 20 * time.Second}
	reqw, err := http.NewRequest("POST", fmt.Sprintf(AllMetricsURL, addr), bytes.NewBuffer(metrics))
	if err != nil {
		return err
	}
	reqw.Header.Set("Content-Type", "application/json")
	reqw.Header.Set("Content-Encoding", "gzip")
	reqw.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(reqw)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not expected status code: %d", resp.StatusCode)
	}

	return nil
}
