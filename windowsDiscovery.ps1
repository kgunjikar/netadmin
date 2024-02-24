# Discover Windows machines in the network via Active Directory
$computers = Get-ADComputer -Filter 'OperatingSystem -Like "Windows"' | Select-Object -ExpandProperty Name

# Loop through each computer and try to get system info
foreach ($computer in $computers) {
    Write-Host "Getting system info for: $computer"
    # Attempt to get system information
    try {
        $systemInfo = Get-WmiObject -Class Win32_OperatingSystem -ComputerName $computer
        # Display basic system information
        Write-Host "System: $($systemInfo.CSName)"
        Write-Host "Version: $($systemInfo.Version)"
        Write-Host "Manufacturer: $($systemInfo.Manufacturer)"
        Write-Host "Last Boot Up Time: $($systemInfo.LastBootUpTime)"
    } catch {
        Write-Host "Failed to get system info for: $computer"
    }
}
