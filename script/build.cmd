:: Script for building an application for Windows

cd ..

:: Del old data
IF EXIST bin RMDIR bin /Q /S
IF EXIST debug.out DEL debug.out /Q /S
IF EXIST test.out DEL test.out /Q /S

IF EXIST go.mod DEL go.mod /Q /S
IF EXIST go.sum DEL go.sum /Q /S

:: Set env variable
set GOROOT=c:\go
set GOOS=windows

:: Code format
go fmt

go mod init github.com/karpovdl/timef

:: Code build x64
set GOARCH=amd64
go build -ldflags "-s -w" -i -v

:: Code build x86
set GOARCH=386
go build -ldflags "-s -w" -i -v
