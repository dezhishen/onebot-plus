# !/bin/bash
MYDIR="."
INCLUDE=$GOPATH/include
SDK_PATH=$MYDIR/../../../../onebot-sdk
echo $MYDIR
echo $INCLUDE
echo $SDK_PATH
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        protoc --proto_path=$INCLUDE \
        --proto_path=${SDK_PATH}/pkg/model \
        --proto_path=$MYDIR \
        --go_out=plugins=grpc:./  \
        $file
    fi
done   

mv ./github.com/dezhishen/onebot-plus/pkg/plugin/base/*.pb.go ./
rm -rf ./github.com