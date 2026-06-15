package models

type PathEntry struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

type Group struct {
	Name  string       `json:"name"`
	Paths []PathEntry `json:"paths"`
}

type Config struct {
	Groups map[string]Group `json:"groups"`
}
