# 📁 ofolders - Professional Folder Management CLI

A powerful Windows CLI tool to organize and open your frequently used folder paths in Windows Explorer. Similar to Git workflow but for managing your files.

**Created by:** Carlos Sigua

## 🚀 Quick Start

### 1. Create a Group
```bash
ofolders group create trabajo
```

### 2. Add Paths to the Group
```bash
ofolders add "D:\OneDrive\2026\CONTRATOS" --alias "Contratos" --group "trabajo"
ofolders add "D:\OneDrive\2026\PROYECTOS" --alias "Proyectos" --group "trabajo"
```

### 3. Open the Group in Explorer
```bash
ofolders run trabajo
```
This will open all paths as tabs in your existing Explorer window.

---

## 📖 Complete Command Reference

### Group Management

#### Create a new group
```bash
ofolders group create <name>
```
**Example:**
```bash
ofolders group create work
ofolders group create personal
```

#### List all groups
```bash
ofolders group list
```
**Output:**
```
📁 Available Groups:
════════════════════════════════════
  • work (5 paths)
  • personal (3 paths)
```

#### List paths in a specific group
```bash
ofolders group list <name>
```
**Example:**
```bash
ofolders group list trabajo
```
**Output:**
```
📂 Group: trabajo
════════════════════════════════════
  1. [Contratos] D:\OneDrive\2026\CONTRATOS
  2. [Proyectos] D:\OneDrive\2026\CONTRATOS\PROYECTOS
```

#### Delete a group
```bash
ofolders group delete <name>
```
**Example:**
```bash
ofolders group delete personal
```

---

### Path Management

#### Add a path to a group
```bash
ofolders add <path> --alias <alias> --group <group>
```
**Example:**
```bash
ofolders add "D:\Projects\ClientA" --alias "ClientA" --group "work"
ofolders add "C:\Users\adminos\Documents" --alias "Documents" --group "personal"
```

#### Update a path in a group
```bash
ofolders update <alias> <new-path> --group <group>
```
**Example:**
```bash
ofolders update Contratos "D:\NewLocation\Contratos" --group trabajo
```

#### Remove a path from a group
```bash
ofolders remove <alias> --group <group>
```
**Example:**
```bash
ofolders remove Proyectos --group trabajo
```

---

### Execution

#### Open a group in Explorer
```bash
ofolders run <group>
```
**What it does:**
1. Finds your active Explorer window
2. Creates new tabs for each path in the group
3. Opens each path in its own tab

**Requirements:**
- Must have an Explorer window already open
- Windows explorer must be running

**Example:**
```bash
ofolders run trabajo
ofolders run personal
```

---

### Maintenance

#### Clean all paths from a group
```bash
ofolders clean <group>
```
**Example:**
```bash
ofolders clean trabajo
```

#### Clean all groups (reset everything)
```bash
ofolders clean --all
```
⚠️ **Warning:** This will remove all paths from all groups!

---

### Information

#### Show version and author
```bash
ofolders version
```
**Output:**
```
ofolders v1.0.0
Created by Carlos Sigua
```

#### Show statistics
```bash
ofolders info
```
**Output:**
```
📊 ofolders Statistics
════════════════════════════════════
Total groups: 2
Total paths: 8
Config location: C:\Users\adminos\AppData\Roaming\ofolders\config.json
════════════════════════════════════
Created by: Carlos Sigua
Version: v1.0.0
```

#### Show help
```bash
ofolders --help
ofolders group --help
ofolders add --help
```

---

## 💾 Configuration

Your configuration is stored at:
```
%APPDATA%\ofolders\config.json
```

Example config structure:
```json
{
  "groups": {
    "trabajo": {
      "name": "trabajo",
      "paths": [
        {
          "path": "D:\\OneDrive\\2026\\CONTRATOS",
          "alias": "Contratos"
        },
        {
          "path": "D:\\OneDrive\\2026\\CONTRATOS\\PROYECTOS",
          "alias": "Proyectos"
        }
      ]
    }
  }
}
```

---

## 🎯 Use Cases

### Daily Work Setup
```bash
ofolders run trabajo
# Opens all work-related folders in Explorer
```

### Project Management
```bash
ofolders group create project_alpha
ofolders add "C:\projects\alpha\src" --alias "Source" --group project_alpha
ofolders add "C:\projects\alpha\docs" --alias "Documentation" --group project_alpha
ofolders add "C:\projects\alpha\tests" --alias "Tests" --group project_alpha
ofolders run project_alpha
```

### Organize by Department
```bash
ofolders group create engineering
ofolders group create design
ofolders group create marketing

ofolders add "\\network\engineering\code" --alias "Code" --group engineering
ofolders add "\\network\design\assets" --alias "Assets" --group design
ofolders add "\\network\marketing\campaigns" --alias "Campaigns" --group marketing
```

---

## 🔧 Troubleshooting

### "Explorer window not found"
- Make sure you have a Windows Explorer window open before running `ofolders run`
- Click on a File Explorer window first, then run the command

### "Group not found"
- Use `ofolders group list` to see all available groups
- Group names are case-sensitive

### "Path already exists" error
- Each alias must be unique within a group
- Use `ofolders group list <group>` to see existing aliases

---

## 📝 Tips & Tricks

1. **Use descriptive aliases:** Instead of "Folder1", use "ProjectSource" or "ClientFiles"
2. **Organize by context:** Create separate groups for different projects/departments
3. **Use network paths:** Works with UNC paths like `\\server\share\folder`
4. **Quick access:** Add ofolders to your PATH for global access from any terminal

---

## 🐛 Reporting Issues

If you encounter any issues, check:
1. Explorer window is open
2. Paths exist and are accessible
3. Config file is valid JSON at `%APPDATA%\ofolders\config.json`

---

## 📄 License

Developed with ❤️ for efficient file management on Windows.
