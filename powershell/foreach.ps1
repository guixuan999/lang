# foreach
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

Write-Host -----------------------------
Write-Host

# for
for($i=1; $i -le 5; $i++) {
    Write-Host $i
}

Write-Host -----------------------------
Write-Host

# while
$i = 5
while($i -gt 0) {
    Write-Host $i
    $i--
}

Write-Host -----------------------------
Write-Host

# do ... while
do {
    Write-Host Hello
    $i++
} while ($i -lt 10)

Write-Host -----------------------------
Write-Host

# if
$a = 5
if($a -gt 0 -and $a -lt 10) {
    Write-Host "<0, 10>"
} elseif($a -ge 10) {
    Write-Host "[10, INFINIT>"
} else {
    Write-Host "<-INFINIT, 0]"
}

Write-Host -----------------------------
Write-Host

# switch (style 1)
$b = 5
switch($b) {
    "b" {
        Write-Host I am b
    }
    "c" {
        Write-Host I am c
    }
    10 {
        Write-Host I am 10
    }
    default {
         Write-Host I am none...
    }
}

Write-Host -----------------------------
Write-Host

# switch (style 2)
switch($b) {
    {$b -eq "b"} {
        Write-Host I am b
    }
    {$b -eq "c"} {
        Write-Host I am c
    }
    {$b -le 10} {
        Write-Host I am less than or equal to 10
    }
    default {
         Write-Host I am none...
    }
}