#!/usr/bin/env bash
NAME="web-IDE-back-end"
#编译pb
# for x in **/*.proto; 
# do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done
find ./ -name "*.go" | xargs gofmt -w -s -l
mkdir -p output/bin output/config
cp -r config/ output/config/


name=`uname -o`
if [[ $name =~ "Darwin" ]];then
    GOOS=darwin GOARCH=amd64 go build -o output/bin/${NAME}.out
    chmod +x output/bin/${NAME}.out
elif [[ $name =~ "GNU/Linux" ]];then
    GOOS=linux GOARCH=amd64 go build -o output/bin/${NAME}.out
    chmod +x output/bin/${NAME}.out
else
    GOOS=windows GOARCH=amd64 go build -o output/bin/${NAME}.exe
fi
chmod +x output/start.sh