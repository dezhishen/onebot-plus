MYDIR=C:/github/onebot-plus/pkg/plugin
INCLUDE=$GOPATH/include
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        # a="${file/message_element_/}"
        # b="${a/\.\//}"
        # c="${b/.proto/}"
        # echo "case \"$c\":
		# result.Data = &MessageSegmentGRPC_MessageElement${c^}{
		# 	MessageElement${c^}: msg.Data.(*MessageElement${c^}).ToGRPC(),
		# }"
        # echo "$file" 
        #echo "protoc --proto_path=$INCLUDE --proto_path=$MYDIR --go_out=plugins=grpc:\model\  $MYDIR/$file"
        protoc --proto_path=$INCLUDE --proto_path=C:/github --proto_path=C:/github/onebot-sdk/pkg/proto --proto_path=$MYDIR --go_out=plugins=grpc:../plugin  $MYDIR/$file
        #echo "$x $file $file = 2;"
    fi
done   