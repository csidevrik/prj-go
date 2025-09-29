# Script PowerShell: iniciar-estructura.ps1

Add-Type -AssemblyName System.Windows.Forms
$folderBrowser = New-Object System.Windows.Forms.FolderBrowserDialog
$folderBrowser.Description = "Selecciona la carpeta raíz para CUMPLIMIENTO30PUNTOS"
$null = $folderBrowser.ShowDialog()
$basePath = $folderBrowser.SelectedPath

if (-not $basePath) {
    Write-Host "Cancelado por el usuario."
    exit
}

# Estructura principal
$mainFolder = Join-Path $basePath "c30p"
New-Item -ItemType Directory -Path $mainFolder -Force | Out-Null

# Subcarpetas clave
New-Item -ItemType Directory -Path (Join-Path $mainFolder "backend\internal\files") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "backend\internal\metadata") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "backend\internal\compliance") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "backend\internal\api") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "backend\cmd\server") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "frontend") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "tools") -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $mainFolder "dist\manifest") -Force | Out-Null

# Carpeta de datos y meses
$dataFolder = Join-Path $mainFolder "data"
New-Item -ItemType Directory -Path $dataFolder -Force | Out-Null

$meses = @(
    "01ENERO", "02FEBRERO", "03MARZO", "04ABRIL", "05MAYO", "06JUNIO",
    "07JULIO", "08AGOSTO", "09SEPTIEMBRE", "10OCTUBRE", "11NOVIEMBRE", "12DICIEMBRE"
)

foreach ($mes in $meses) {
    New-Item -ItemType Directory -Path (Join-Path $dataFolder $mes) -Force | Out-Null
}

# JSON inicial
$cumplimientoDir = Join-Path $dataFolder "cumplimiento"
New-Item -ItemType Directory -Path $cumplimientoDir -Force | Out-Null
$jsonPath = Join-Path $cumplimientoDir "cumplimiento30puntos.json"
"[]" | Set-Content -Path $jsonPath -Encoding UTF8

Write-Host "✅ Estructura creada correctamente en: $mainFolder"
