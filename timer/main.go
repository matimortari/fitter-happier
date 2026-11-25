package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		d, err := parseDuration(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "usage: timer [duration]")
			fmt.Fprintln(os.Stderr, "  duration examples: 10s, 5m, 1h30m, 90")
			os.Exit(1)
		}
		countdown(d)
		return
	}

	stopwatch()
}

func stopwatch() {
	start := time.Now()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	fmt.Fprintln(os.Stderr, "stopwatch running — Ctrl+C to stop")

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\r%s", formatElapsed(time.Since(start)))
		case <-done:
			elapsed := time.Since(start)
			fmt.Printf("\r%s\n", formatElapsed(elapsed))
			return
		}
	}
}

func countdown(d time.Duration) {
	deadline := time.Now().Add(d)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	fmt.Fprintln(os.Stderr, "countdown running — Ctrl+C to stop")

	for {
		remaining := time.Until(deadline)
		if remaining <= 0 {
			fmt.Printf("\r%s\n", formatElapsed(0))
			fmt.Fprintln(os.Stderr, "done")
			return
		}

		select {
		case <-ticker.C:
			fmt.Printf("\r%s", formatElapsed(remaining))
		case <-done:
			fmt.Printf("\r%s\n", formatElapsed(remaining))
			os.Exit(130)
		}
	}
}

// parseDuration accepts Go durations (10s, 5m, 1h30m) or a bare number of seconds.
func parseDuration(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty duration")
	}

	if d, err := time.ParseDuration(s); err == nil {
		if d <= 0 {
			return 0, fmt.Errorf("duration must be positive")
		}
		return d, nil
	}

	var secs int
	if _, err := fmt.Sscanf(s, "%d", &secs); err == nil && fmt.Sprintf("%d", secs) == s {
		if secs <= 0 {
			return 0, fmt.Errorf("duration must be positive")
		}
		return time.Duration(secs) * time.Second, nil
	}

	return 0, fmt.Errorf("invalid duration: %s", s)
}

func formatElapsed(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	totalMs := d.Milliseconds()
	h := totalMs / 3_600_000
	m := (totalMs % 3_600_000) / 60_000
	s := (totalMs % 60_000) / 1_000
	ms := (totalMs % 1_000) / 100 // tenths

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d.%d", h, m, s, ms)
	}
	return fmt.Sprintf("%02d:%02d.%d", m, s, ms)
}
