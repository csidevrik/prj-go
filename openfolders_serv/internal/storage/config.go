package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adminos/openfolders_serv/internal/models"
)

type ConfigStore struct {
	filePath string
	config   *models.Config
}

func NewConfigStore() (*ConfigStore, error) {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		appDataPath = os.Getenv("USERPROFILE")
	}

	configDir := filepath.Join(appDataPath, "ofolders")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	filePath := filepath.Join(configDir, "config.json")

	store := &ConfigStore{
		filePath: filePath,
		config:   &models.Config{Groups: make(map[string]models.Group)},
	}

	if err := store.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return store, nil
}

func (cs *ConfigStore) Load() error {
	data, err := os.ReadFile(cs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, cs.config); err != nil {
		return err
	}

	return nil
}

func (cs *ConfigStore) Save() error {
	data, err := json.MarshalIndent(cs.config, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(cs.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

func (cs *ConfigStore) CreateGroup(name string) error {
	if _, exists := cs.config.Groups[name]; exists {
		return fmt.Errorf("group '%s' already exists", name)
	}

	cs.config.Groups[name] = models.Group{
		Name:  name,
		Paths: []models.PathEntry{},
	}

	return cs.Save()
}

func (cs *ConfigStore) DeleteGroup(name string) error {
	if _, exists := cs.config.Groups[name]; !exists {
		return fmt.Errorf("group '%s' not found", name)
	}

	delete(cs.config.Groups, name)
	return cs.Save()
}

func (cs *ConfigStore) AddPath(groupName, path, alias string) error {
	group, exists := cs.config.Groups[groupName]
	if !exists {
		return fmt.Errorf("group '%s' not found", groupName)
	}

	// Check if alias already exists in this group
	for _, p := range group.Paths {
		if p.Alias == alias {
			return fmt.Errorf("alias '%s' already exists in group '%s'", alias, groupName)
		}
	}

	group.Paths = append(group.Paths, models.PathEntry{
		Path:  path,
		Alias: alias,
	})

	cs.config.Groups[groupName] = group
	return cs.Save()
}

func (cs *ConfigStore) RemovePath(groupName, alias string) error {
	group, exists := cs.config.Groups[groupName]
	if !exists {
		return fmt.Errorf("group '%s' not found", groupName)
	}

	found := false
	newPaths := []models.PathEntry{}

	for _, p := range group.Paths {
		if p.Alias != alias {
			newPaths = append(newPaths, p)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("path with alias '%s' not found in group '%s'", alias, groupName)
	}

	group.Paths = newPaths
	cs.config.Groups[groupName] = group
	return cs.Save()
}

func (cs *ConfigStore) UpdatePath(groupName, alias, newPath string) error {
	group, exists := cs.config.Groups[groupName]
	if !exists {
		return fmt.Errorf("group '%s' not found", groupName)
	}

	found := false
	for i, p := range group.Paths {
		if p.Alias == alias {
			group.Paths[i].Path = newPath
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("path with alias '%s' not found in group '%s'", alias, groupName)
	}

	cs.config.Groups[groupName] = group
	return cs.Save()
}

func (cs *ConfigStore) GetGroup(name string) (models.Group, error) {
	group, exists := cs.config.Groups[name]
	if !exists {
		return models.Group{}, fmt.Errorf("group '%s' not found", name)
	}
	return group, nil
}

func (cs *ConfigStore) ListGroups() map[string]models.Group {
	return cs.config.Groups
}

func (cs *ConfigStore) GetConfigPath() string {
	return cs.filePath
}

func (cs *ConfigStore) CleanGroup(name string) error {
	group, exists := cs.config.Groups[name]
	if !exists {
		return fmt.Errorf("group '%s' not found", name)
	}

	group.Paths = []models.PathEntry{}
	cs.config.Groups[name] = group
	return cs.Save()
}

func (cs *ConfigStore) CleanAll() error {
	cs.config.Groups = make(map[string]models.Group)
	return cs.Save()
}
