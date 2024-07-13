package repeat

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

// RepeatNet sends request with repeating
func RepeatNet(ctx context.Context, client *http.Client, req *http.Request) error {
	sec := 1
	tickerResend := time.NewTicker(time.Duration(sec) * time.Second)
	var urlErr *net.OpError

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		case <-tickerResend.C:
			// Sending new request
			err := cliSend(client, req)

			// Check if error is OpError
			if errors.As(err, &urlErr) {
				if sec == 5 && err != nil {
					return err
				}
			} else {
				return err
			}

			// If err is nil then return
			if err == nil {
				return nil
			}

			// Resetting ticker with new interval
			sec += 2
			tickerResend.Reset(time.Duration(sec) * time.Second)
		}
	}
}

// Sends to server
func cliSend(client *http.Client, reqw *http.Request) error {
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
