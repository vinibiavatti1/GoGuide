package main

import (
	"errors"
	"fmt"
	"strings"
)

const Log string = `
2026-10-16T10:00:01Z INFO Application started successfully
2026-10-16T10:00:05Z INFO Loading configuration from environment
2026-10-16T10:00:12Z WARN Deprecated API endpoint called: /v1/legacy-login
2026-10-16T10:01:00Z INFO User session established for ID: 4892
2026-10-16T10:02:15Z ERROR Failed to connect to primary database cluster
2026-10-16T10:02:16Z INFO Retrying database connection on replica...
corrupted-line
2026-10-16T10:02:18Z INFO Database connection established on replica
2026-10-16T10:03:40Z WARN High memory usage detected: 85% utilized
2026-10-16T10:04:02Z INFO Processing batch job #1042
2026-10-16T10:05:11Z ERROR NullPointerException while parsing user payload
2026-10-16T10:05:12Z INFO Recovered from panic in worker goroutine
2026-10-16T10:06:30Z INFO Cache invalidated for key: user_settings_4892
corrupted-line
2026-10-16T10:07:00Z WARN Rate limit threshold reached for IP 192.168.1.50
2026-10-16T10:08:22Z INFO Incoming request GET /api/v2/status
2026-10-16T10:09:05Z ERROR Timeout waiting for downstream payment service response
2026-10-16T10:10:14Z INFO Scheduled background cleanup task started
2026-10-16T10:11:00Z WARN Configuration property 'max_retries' is unset, using default value 3
2026-10-16T10:12:45Z INFO Scheduled background cleanup task finished successfully
2026-10-16T10:13:30Z ERROR Unauthorized access attempt detected on /api/v2/admin
corrupted-line
2026-10-16T10:15:00Z INFO Shutting down server gracefully
`

var ErrMalformedEntry = errors.New("Malformed Entry")

func CountByLevel(log string) map[string]int {
	result := map[string]int{}
	for line := range strings.SplitSeq(log, "\n") {
		if len(line) == 0 {
			continue
		}
		lineLevel, err := getLevel(line)
		if err != nil {
			lineLevel = "Unknown"
		}
		v, ok := result[lineLevel]
		if !ok {
			result[lineLevel] = 1
		} else {
			result[lineLevel] = v + 1
		}
	}
	return result
}

func ExtractEntries(log string, level string) []string {
	result := []string{}
	for line := range strings.SplitSeq(log, "\n") {
		if len(line) == 0 {
			continue
		}
		lineLevel, err := getLevel(line)
		if err != nil {
			continue
		}
		if lineLevel == level {
			result = append(result, line)
		}
	}
	return result
}

func getLevel(entry string) (string, error) {
	parts := strings.Split(entry, " ")
	if len(parts) < 2 {
		return "", ErrMalformedEntry
	}
	return parts[1], nil
}

func main() {
	ch1 := make(chan map[string]int)
	ch2 := make(chan []string)

	go func() {
		ch1 <- CountByLevel(Log)
	}()
	go func() {
		ch2 <- ExtractEntries(Log, "ERROR")
	}()

	for range 2 {
		select {
		case x := <-ch1:
			fmt.Printf("Count by level: %v\n", x)
		case x := <-ch2:
			fmt.Printf("Extract Errors (n:%d): %s\n", len(x), strings.Join(x, "|"))
		}
	}
}
