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
mkdir product\win\storage

echo.
echo [1/3] Build Go Daemon...

go build -ldflags="-s -w" ^
-o product\win\kyrafs.exe ^
./cmd/kyrafs

if errorlevel 1 goto error

echo.
echo [2/3] Build Python Engine...

pyinstaller ^
--clean ^
--onefile ^
--name kyrafs-engine ^
--distpath product\win ^
engine\main.py

if errorlevel 1 goto error

echo.
echo [3/3] Copy Configuration...

copy config\kyrafs.ini product\win\ >nul

echo.
echo ===================================
echo Build Success
echo ===================================

tree product

pause
exit

:error

echo.
echo Build Failed.
pause