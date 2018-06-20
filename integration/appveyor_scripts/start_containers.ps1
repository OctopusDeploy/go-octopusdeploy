Push-Location -Path 'integration\appveyor_scripts'
Write-Output "Starting docker-compose. This may take a while..."
Start-Process -FilePath 'docker-compose' -ArgumentList 'up -d' -Wait -NoNewWindow
Pop-Location

# Wait for Octopus to be ready
$octopusReady = $false
$retryCount

do {
    try {
        Write-Output "Trying to connect to Octopus..."
        $result = Invoke-WebRequest -UseBasicParsing -Uri "http://$($env:LOCAL_MACHINE_IP):81"
        if ($result.StatusCode -eq 200) {
            $octopusReady = $true
        }
    }
    catch {
        Write-Output "Next attempt in 15 seconds"
        Start-Sleep -Seconds 15
    }

    $retryCount++

}until($retryCount -eq 5 -or $octopusReady)

if ($retryCount -eq 5)
{
    Write-Error "Failed to bring up Octopus Deploy container correctly."
}
else
{
    Write-Output "Docker containers have started."
}
