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
--------------------------sudo apt install mingw-w64 ***-1
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
go build -o app.exe
-----------------------

--------------------------sudo apt install mingw-w64 ***-1
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
go build -ldflags="-s -w -H windowsgui" -o app.exe

---------------------------------------------------------
      VALUE "LegalCopyright", "Copyright (C) 2026 Nawakarit, GNU General Public License v3.0"

      VALUE "LegalCopyright", "Copyright © 2026 Nawakarit, GNU General Public License v3.0"



2-GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ 
CC=x86_64-w64-mingw32-gcc \ 
go build -o app.exe

3-GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ 
CC=x86_64-w64-mingw32-gcc \ 
go build -o app.exe

4-CC=x86_64-w64-mingw32-gcc fyne package -os windows

***5-rsrc -manifest app.rc -ico icon.ico -o rsrc.syso
go build -o app.exe

***6-x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso

fyne package -os windows

--x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \ 
CC=x86_64-w64-mingw32-gcc \ 
go build -o app.exe

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 
CC=x86_64-w64-mingw32-gcc 
go build -o app.exe

CC=x86_64-w64-mingw32-gcc fyne package -os windows

fyne package -os windows

#!/bin/bash
7--
echo "🔧 build resource..."
x86_64-w64-mingw32-windres app.rc -O coff -o rsrc.syso

echo "📦 build exe..."
CC=x86_64-w64-mingw32-gcc fyne package -os windows

echo "✅ done"

