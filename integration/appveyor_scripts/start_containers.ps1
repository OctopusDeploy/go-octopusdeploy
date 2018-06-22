$ErrorActionPreference = 'Stop'

# Load Functions
$functionFolder = Get-ChildItem -Path (Join-Path -Path $PWD -ChildPath 'integration\appveyor_scripts\functions')
foreach ($function in $functionFolder) { . $function.FullName }

Write-Output "Starting docker-compose. This may take several minutes"

$workingDirectory = Join-Path -Path $PWD -ChildPath 'integration\appveyor_scripts'

Start-ProcessAdvanced -FilePath 'docker-compose.exe' -ArgumentList 'up -d' -WorkingDirectory $workingDirectory -SleepTime 60 -EnableWaitMessage:$true

$octopusReady = $false
$maxRetries = 10

$localMachineIP = Get-DockerAdapterIP

Write-Output "Machine IP Address: $($localMachineIP)"

do {
    try {
        Write-Output "Trying to connect to Octopus..."
        $result = Invoke-WebRequest -UseBasicParsing -Uri "http://$($localMachineIP):81"
        if ($result.StatusCode -eq 200) {
            $octopusReady = $true
        }
    }
    catch {
        Write-Output "Next attempt in 15 seconds"
        Start-Sleep -Seconds 15
    }

    $retryCount++

}until($retryCount -eq $maxRetries -or $octopusReady)

if ($retryCount -eq $maxRetries)
{
    Write-Error "Failed to bring up Octopus Deploy container correctly."
}
else
{
    Write-Output "Docker containers have started."
}
