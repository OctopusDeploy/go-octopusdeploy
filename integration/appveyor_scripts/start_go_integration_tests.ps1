. ".\integration\appveyor_scripts\Start-ProcessAdvanced.ps1"

$localMachineIP = Get-NetAdapter | Where-Object { $_.Name -like "*DockerNAT*" } | Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' }
$localMachineIP = $localMachineIP.IPAddress

if ([string]::IsNullOrEmpty($localMachineIP)) {
    Write-Error "Cannot get Docker Adapaters IP Address"
}
$OCTOPUS_URL = "http://$($localMachineIP):81" #Octopus URL
$OCTOPUS_APIKEY = Get-Content -Path (Join-Path -Path $PWD -ChildPath "octopus_key.txt")

Start-ProcessAdvanced -FilePath 'go' -ArgumentList "test -v -timeout 30s ./integration/..." -EnvironmentKeyValues @{ OCTOPUS_URL = $OCTOPUS_URL; OCTOPUS_APIKEY = $OCTOPUS_APIKEY } -Verbose
