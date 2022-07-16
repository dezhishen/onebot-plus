
#export CGO_ENABLED=1

export GOARCH=amd64
export GOOS=windows
go build -o ./plugin/impl/plugin-windows-amd64.exe ../plugin/impl

export GOARM=7
export GOARCH=arm64
export GOOS=linux

go build -o ./plugin/impl/plugin-linux-arm64v7 ../plugin/impl

export GOARCH=amd64
export GOOS=darwin
go build -o ./plugin/impl/plugin-linux-amd64 ../plugin/impl


export GOARCH=amd64
export GOOS=linux
go build -o ./plugin/impl/plugin-linux-amd64 ../plugin/impl
