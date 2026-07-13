@echo off

echo ===================================
echo        KyraFS Release Builder
echo ===================================
echo.

:: Clean

if exist build rmdir /S /Q build
if exist dist rmdir /S /Q dist
if exist product rmdir /S /Q product

mkdir product
mkdir product\win

echo.
echo [1/2] Build Go Daemon...

go build -ldflags="-s -w" ^
-o product\win\kyrafs.exe ^
./cmd/kyrafs

if errorlevel 1 goto error

echo.
echo [2/2] Build Python Engine...

pyinstaller ^
--clean ^
--onefile ^
--name kyrafs-engine ^
--distpath product\win ^
engine\main.py

if errorlevel 1 goto error

echo.
echo ===================================
echo Build Success
echo ===================================

tree product

echo.
echo Next steps:
echo.
echo     cd product\win
echo     kyrafs init
echo     kyrafs serve
echo.

pause
exit

:error

echo.
echo Build Failed.
pause