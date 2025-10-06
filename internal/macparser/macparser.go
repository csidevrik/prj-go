// internal/macparser/macparser.go
package macparser

import (
	"fmt"
	"regexp"
	"strings"
)

type MACFormats struct {
	Raw            string
	Normalized     string
	LinuxFormat    string
	WindowsFormat  string
	HuaweiFormat   string
	CiscoFormat    string
	PossibleOrigin string
}

func DetectAndNormalize(mac string) (*MACFormats, error) {
	raw := strings.ToLower(strings.TrimSpace(mac))
	norm := regexp.MustCompile(`[^0-9a-f]`).ReplaceAllString(raw, "")
	if len(norm) != 12 {
		return nil, fmt.Errorf("MAC inválida: %s", mac)
	}

	format := &MACFormats{
		Raw:            raw,
		Normalized:     norm,
		LinuxFormat:    fmt.Sprintf("%s:%s:%s:%s:%s:%s", norm[0:2], norm[2:4], norm[4:6], norm[6:8], norm[8:10], norm[10:12]),
		WindowsFormat:  fmt.Sprintf("%s-%s-%s-%s-%s-%s", norm[0:2], norm[2:4], norm[4:6], norm[6:8], norm[8:10], norm[10:12]),
		HuaweiFormat:   fmt.Sprintf("%s-%s-%s", norm[0:4], norm[4:8], norm[8:12]),
		CiscoFormat:    fmt.Sprintf("%s.%s.%s", norm[0:4], norm[4:8], norm[8:12]),
		PossibleOrigin: detectOrigin(raw),
	}
	return format, nil
}

func detectOrigin(mac string) string {
	switch {
	case regexp.MustCompile(`^([0-9a-f]{2}[-:]){5}[0-9a-f]{2}$`).MatchString(mac):
		if strings.Contains(mac, "-") {
			return "Windows"
		}
		return "Linux/Mikrotik"
	case regexp.MustCompile(`^[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}$`).MatchString(mac):
		return "Huawei"
	case regexp.MustCompile(`^[0-9a-f]{4}\.[0-9a-f]{4}\.[0-9a-f]{4}$`).MatchString(mac):
		return "Cisco"
	default:
		return "Desconocido"
	}
}
