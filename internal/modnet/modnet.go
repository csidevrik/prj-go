// internal/modnet/modnet.go
package modnet

import (
	"time"

	"github.com/go-ping/ping"
)

func Monitor(listIP []string) map[string]string {
	results := make(map[string]string)
	for _, ip := range listIP {
		pinger, _ := ping.NewPinger(ip)
		pinger.Count = 1
		pinger.Timeout = 3 * time.Second
		err := pinger.Run()
		if err != nil {
			results[ip] = "🔴 bad"
		} else {
			results[ip] = "🟢 ok"
		}
	}
	return results
}
