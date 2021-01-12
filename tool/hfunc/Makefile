version := $(shell /bin/date "+%Y-%m-%d %H:%M")

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" hfunc.go
	$(if $(shell command -v upx), upx hfunc)
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o hfunc-darwin hfunc.go
	$(if $(shell command -v upx), upx hfunc-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o hfunc.exe hfunc.go
	$(if $(shell command -v upx), upx hfunc.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o hfunc-linux hfunc.go
	$(if $(shell command -v upx), upx hfunc-linux)
