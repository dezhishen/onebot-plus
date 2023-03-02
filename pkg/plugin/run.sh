# !/bin/bash
MYDIR=`dirname $0`
INCLUDE=$GOPATH/include
SDK_PATH=$MYDIR/../../../onebot-sdk
echo $MYDIR
echo $INCLUDE
echo $SDK_PATH
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        echo "$file"
        protoc --proto_path=$INCLUDE \
        --proto_path=$MYDIR \
        --proto_path=${SDK_PATH}/pkg/model \
        --go_out=plugins=grpc:./  \
        $file

    fi
done   

mv ./github.com/dezhishen/onebot-plus/pkg/plugin/*.pb.go ./
rm -rf ./github.com