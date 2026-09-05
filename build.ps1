$ErrorActionPreference = 'Stop'
Set-Location $PSScriptRoot
$env:CGO_ENABLED = '1'
go mod tidy
go build -ldflags='-H=windowsgui' -o SubFuffeR.exe .
