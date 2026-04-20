#!/bin/bash
set -e
export PATH=/usr/local/go/bin:$PATH

echo "🔧 build resource..."
x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso

echo "📦 build exe..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ CC=x86_64-w64-mingw32-gcc \ go build -o app.exe

echo "✅ done"


2-GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ 
CC=x86_64-w64-mingw32-gcc \ 
go build -o app.exe

3-GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ CC=x86_64-w64-mingw32-gcc \ go build -o app.exe

4-CC=x86_64-w64-mingw32-gcc fyne package -os windows

5-rsrc -manifest app.manifest -ico icon.ico -o rsrc.syso
go build -o app.exe

6-x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso
fyne package -os windows
