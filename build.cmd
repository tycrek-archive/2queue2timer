@echo off
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
set CC=C:\Mingw\bin\gcc.exe
set CXX=C:\Mingw\bin\g++.exe
go build -buildmode=exe 2q2t.go