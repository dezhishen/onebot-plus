# !/bin/bash
# 执行当前目录下所有子目录run.sh脚本
for file in `ls`
do
    if [ -d $file ]; then
        cd $file
        if [ -f run.sh ]; then
            echo "run.sh in $file"
            sh run.sh
        fi
        cd ..
    fi
done