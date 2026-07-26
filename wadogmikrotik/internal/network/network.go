package network

import (
    "time"

    "github.com/go-ping/ping"
)

type Pinger interface {
    Ping(host string, timeout time.Duration, attempts int, delay time.Duration) (bool, error)
}

type pingClient struct{}

func New() Pinger {
    return &pingClient{}
}

func (p *pingClient) Ping(host string, timeout time.Duration, attempts int, delay time.Duration) (bool, error) {
    for attempt := 0; attempt < attempts; attempt++ {
        pinger, err := ping.NewPinger(host)
        if err != nil {
            return false, err
        }
        pinger.SetPrivileged(true)
        pinger.Count = 1
        pinger.Timeout = timeout

        err = pinger.Run()
        if err == nil {
            stats := pinger.Statistics()
            if stats.PacketsRecv > 0 {
                return true, nil
            }
        }

        if attempt+1 < attempts {
            time.Sleep(delay)
        }
    }

    return false, nil
}
