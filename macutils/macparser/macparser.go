package macparser

import (
	"fmt"
	"regexp"
	"strings"
)

// Estructura principal
type MACFormats struct {
	Raw            string
	Normalized     string
	LinuxFormat    string
	WindowsFormat  string
	HuaweiFormat   string
	CiscoFormat    string
	PossibleOrigin string
	MACType        string // Unicast, Multicast, Broadcast
	OUI            string // Fabricante según OUI
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
		MACType:        getMACType(norm),
		OUI:            lookupOUI(norm),
	}

	return format, nil
}

// Detección de tipo de MAC por el primer byte
func getMACType(norm string) string {
	if norm == "ffffffffffff" {
		return "Broadcast"
	}
	b := norm[0:2]
	bVal := parseHexByte(b)

	if bVal&1 == 1 {
		return "Multicast"
	}
	return "Unicast"
}

// Conversión simple hex string a byte (por claridad)
func parseHexByte(hexStr string) byte {
	var b byte
	fmt.Sscanf(hexStr, "%x", &b)
	return b
}

// OUI básico
func lookupOUI(norm string) string {
	oui := norm[0:6]
	ouiMap := map[string]string{
		"001a2b": "Intel Corp.",
		"fcfbfb": "Apple Inc.",
		"3cd92b": "Ubiquiti Networks",
		"b827eb": "Raspberry Pi Foundation",
		"000c29": "VMware, Inc.",
		"f4f5e8": "Huawei Technologies",
	}
	if vendor, ok := ouiMap[oui]; ok {
		return vendor
	}
	return "Fabricante desconocido"
}

// Detección del formato de entrada
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
