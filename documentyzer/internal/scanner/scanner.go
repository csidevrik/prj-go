package scanner

import (
    "os"
    "path/filepath"

    "documentyzer/internal/meta"
)

type FolderStatus struct {
    Name       string
    Path       string
    HasMeta    bool
    HasReadme  bool
    HasInstall bool
    Meta       *meta.Meta // nil si no existe meta.json
}

func ScanScriptsDir(repoPath string) ([]FolderStatus, error) {
    scriptsPath := filepath.Join(repoPath, "scripts.d")

    entries, err := os.ReadDir(scriptsPath)
    if err != nil {
        return nil, err
    }

    var folders []FolderStatus

    for _, entry := range entries {
        if !entry.IsDir() {
            continue
        }

        folderPath := filepath.Join(scriptsPath, entry.Name())

        status := FolderStatus{
            Name:       entry.Name(),
            Path:       folderPath,
            HasMeta:    fileExists(filepath.Join(folderPath, "meta.json")),
            HasReadme:  fileExists(filepath.Join(folderPath, "README.md")),
            HasInstall: fileExists(filepath.Join(folderPath, "install.sh")),
        }

        if status.HasMeta {
            status.Meta, _ = meta.Load(folderPath)
        }

        folders = append(folders, status)
    }

    return folders, nil
}

func fileExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}
// ```

// ---

// ## Lo que te da el scanner ahora mismo

// Cuando llamas `ScanScriptsDir("/home/carlos/repos/prj-bash")` obtienes algo así:
// ```
// install-go    → HasMeta: false | HasReadme: false | HasInstall: true
// install-utils → HasMeta: false | HasReadme: false | HasInstall: true
// install-ohmyzsh → HasMeta: false | HasReadme: false | HasInstall: false
// ...