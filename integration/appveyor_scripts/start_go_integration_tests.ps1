$ErrorActionPreference = 'Stop'

# Load Functions
$functionFolder = Get-ChildItem -Path (Join-Path -Path $PWD -ChildPath 'integrations\appveyor_scripts\functions')
foreach ($function in $functionFolder) { . $function.FullName }

$OCTOPUS_URL = Get-DockerAdapterIP
$OCTOPUS_APIKEY = Get-Content -Path (Join-Path -Path $PWD -ChildPath "octopus_key.txt")

Start-ProcessAdvanced -FilePath 'go' -ArgumentList "test -v -timeout 30s ./integration/..." -EnvironmentKeyValues @{ OCTOPUS_URL = $OCTOPUS_URL; OCTOPUS_APIKEY = $OCTOPUS_APIKEY } -Verbose
