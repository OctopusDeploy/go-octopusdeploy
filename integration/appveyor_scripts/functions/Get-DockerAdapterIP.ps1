function Get-DockerAdapterIP
{
    [CmdletBinding()]
    Param
    (
        # Directory to start Docker Compose In
        [Parameter(Mandatory=$false)]
        [string]
        $AdapterName="*DockerNAT*"
    )

    $ErrorActionPreference = 'Stop'

    $localMachineIP = Get-NetAdapter | Where-Object { $_.Name -like $AdapterName } | Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' }
    $localMachineIP = $localMachineIP.IPAddress

    if ([string]::IsNullOrEmpty($localMachineIP)) {
        Write-Error "Cannot get Docker Adapaters IP Address"
    }

    return $localMachineIP
}
