$ErrorActionPreference = "Stop"

$AppName = "cp_tester"
$BinaryName = "cp.exe"
$Repo = "mbicl/cp_tester"
$InstallDir = "$env:LOCALAPPDATA\Programs\$AppName"

# Detect architecture
$Arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

# Get latest version
$Release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
$Version = $Release.tag_name

if (-not $Version) {
    Write-Error "Could not determine latest version"
    exit 1
}

Write-Host "Installing $AppName $Version (windows/$Arch)..."

$Filename = "cp_windows_${Arch}.exe"
$DownloadUrl = "https://github.com/$Repo/releases/download/$Version/$Filename"

# Create install directory
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

$DestPath = Join-Path $InstallDir $BinaryName

Write-Host "Downloading $DownloadUrl..."
Invoke-WebRequest -Uri $DownloadUrl -OutFile $DestPath

# Add to PATH if not already present
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to user PATH. Restart your terminal for changes to take effect."
}

Write-Host "$AppName $Version installed to $DestPath"
