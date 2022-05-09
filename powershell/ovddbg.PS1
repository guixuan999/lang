[CmdletBinding()]
param (
	[Alias('s')]
	[ValidateSet('enable', 'disable', 'remove')]
	[string]$subcmd,

	#[Parameter(ParameterSetName = 'Switches')]
	[Alias('h')]
    [switch]$help
)

Function usage() {
    $SCRIPT_NAME = $MyInvocation.MyCommand.Name
    Write-Host Usage: ".\$SCRIPT_NAME [enable|disable|remove|-h]"
     Write-Host "if no parameter is feed, debug status will be reported."
}


if ($help.IsPresent) {
	usage
	exit 1
}

$pre_location = get-location
write-host $pre_location

if($subcmd -eq 'enable') {
	
	set-location HKLM:\SYSTEM\CurrentControlSet\Services\OVD
	set-itemproperty . debug 1
	Write-Host enabled.
}

if($subcmd -eq 'disable') {
	set-location HKLM:\SYSTEM\CurrentControlSet\Services\OVD
	set-itemproperty . debug 0
	Write-Host disabled.
}

if($subcmd -eq 'remove') {
	set-location HKLM:\SYSTEM\CurrentControlSet\Services\OVD
	remove-itemproperty . debug
	Write-Host property debug removed.
}

if($subcmd -eq '') {
	set-location HKLM:\SYSTEM\CurrentControlSet\Services\OVD
	write-host debug status:(get-itemproperty .).debug
}

set-location $pre_location