@echo off
echo Compiling P92 SSL Tools as tps2015_sslinfo.exe...
go build -o tps2015_sslinfo.exe main.go
if %errorlevel% equ 0 (
    echo.
    echo Compilation successful: tps2015_sslinfo.exe
) else (
    echo.
    echo Compilation failed!
    pause
)
