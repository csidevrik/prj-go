package meta

import (
    "encoding/json"
    "os"
    "path/filepath"
    "time"
)

type Distro struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Status  string `json:"status"` // "ok", "wip", "untested"
}

type Meta struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Distros     []Distro `json:"distros"`
    Version     string   `json:"version"`
    LastUpdated string   `json:"last_updated"`
    RawURL      string   `json:"raw_url"`
    Tags        []string `json:"tags"`
}

const MetaFileName = "meta.json"

func Load(folderPath string) (*Meta, error) {
    path := filepath.Join(folderPath, MetaFileName)
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var m Meta
    if err := json.Unmarshal(data, &m); err != nil {
        return nil, err
    }
    return &m, nil
}

func Save(folderPath string, m *Meta) error {
    m.LastUpdated = time.Now().Format("2006-01-02")
    data, err := json.MarshalIndent(m, "", "  ")
    if err != nil {
        return err
    }
    path := filepath.Join(folderPath, MetaFileName)
    return os.WriteFile(path, data, 0644)
}

func Exists(folderPath string) bool {
    _, err := os.Stat(filepath.Join(folderPath, MetaFileName))
    return err == nil
}