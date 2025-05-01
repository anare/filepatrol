$serviceName = "FilePatrol"
$binaryPath = Join-Path (Get-Location) "runtime\bin\windows\filepatrol.exe"

Write-Host "Installing FilePatrol service from $binaryPath..."

& $binaryPath install
Start-Service -Name $serviceName

Write-Host "FilePatrol service installed and started."