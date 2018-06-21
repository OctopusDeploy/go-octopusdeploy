function Start-DockerCompose
{
    [CmdletBinding()]
    Param
    (
        # Directory to start Docker Compose In
        [Parameter(Mandatory=$false)]
        [string]
        $WorkingDirectory=$PWD.Path,

        # Arguments to start Docker Compose With
        [Parameter(Mandatory=$true)]
        [string]
        $ArgumentList
    )

    $ErrorActionPreference = 'Stop'

    # Setup stdin\stdout redirection
    $StartInfo = New-Object System.Diagnostics.ProcessStartInfo -Property @{
                    FileName = 'docker-compose'
                    Arguments = 'up -d'
                    UseShellExecute = $false
                    RedirectStandardOutput = $true
                    RedirectStandardError = $true
                    WorkingDirectory = $WorkingDirectory
                }

    # Create new process
    $Process = New-Object System.Diagnostics.Process

    # Assign previously created StartInfo properties
    $Process.StartInfo = $StartInfo

    # Register Object Events for stdin\stdout reading
    $OutEvent = Register-ObjectEvent -InputObject $Process -EventName OutputDataReceived -Action {
        Write-Host $Event.SourceEventArgs.Data
    }
    $ErrEvent = Register-ObjectEvent -InputObject $Process -EventName ErrorDataReceived -Action {
        Write-Host $Event.SourceEventArgs.Data
    }

    # Start process
    [void]$Process.Start()

    # Begin reading stdin\stdout
    $Process.BeginOutputReadLine()
    $Process.BeginErrorReadLine()

    # Start a timer to provide some feedback
    $stopwatch =  [system.diagnostics.stopwatch]::StartNew()

    # Do something else while events are firing
    do
    {

        Start-Sleep -Seconds 60
        Write-Output "Currently waited $($stopwatch.Elapsed.Minutes) minutes for docker-compose to finish.."
    }
    while (!$Process.HasExited)

    # Unregister events
    $OutEvent.Name, $ErrEvent.Name |
        ForEach-Object {Unregister-Event -SourceIdentifier $_}
}

Write-Output "Starting docker-compose. This may take several minutes"

$workingDirectory = Join-Path -Path $PWD -ChildPath 'integration\appveyor_scripts'

Start-DockerCompose -WorkingDirectory $workingDirectory -ArgumentList "up -d"

$octopusReady = $false
$maxRetries = 10

$localMachineIP = Get-NetAdapter | Where-Object { $_.Name -like "*DockerNAT*" } | Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' }
$localMachineIP = $localMachineIP.IPAddress

if ([string]::IsNullOrEmpty($localMachineIP)) {
    Write-Error "Cannot get Docker Adapaters IP Address"
}

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
