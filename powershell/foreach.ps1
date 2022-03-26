$Platforms = "windows", "linux"
$Archs = "386", "amd64", "arm", "arm64"
foreach  ($OS in $Platforms) {
    foreach($ARCH in $Archs) {
        if($OS -eq "windows" -and $ARCH -eq "arm64") {
            # windows/arm64 is not support now, skip it!
            continue
        }
        Write-Host $OS --- $ARCH
    }
}
