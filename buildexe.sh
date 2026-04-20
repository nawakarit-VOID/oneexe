#!/bin/bash
set -e
export PATH=/usr/local/go/bin:$PATH

echo "🔧 build resource..."
x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso

echo "📦 build exe..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
go build -o app.exe

echo "✅ done"


