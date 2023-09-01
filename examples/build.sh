
#export CGO_ENABLED=1
mkdir -p ./plugins
export GOARCH=amd64
export GOOS=windows
go build -o ./plugins/plugin-windows-amd64.exe ./plugin/impl.go

export GOARM=7
export GOARCH=arm64
export GOOS=linux

go build -o ./plugins/plugin-linux-arm64v7 ./plugin/impl.go


export GOARCH=amd64
export GOOS=darwin
go build -o ./plugins/plugin-linux-amd64 ./plugin/impl.go

