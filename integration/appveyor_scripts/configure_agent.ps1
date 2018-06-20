Import-Module -Name NetNat
Get-NetNat | Remove-NetNat -Confirm:$false # https://github.com/docker/for-win/issues/598
New-NetFirewallRule -DisplayName "Allow Access to MSSQL from Docker" -Direction Inbound -LocalPort 1433 -Protocol TCP -Action Allow
