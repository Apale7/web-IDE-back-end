#编译pb
for x in **/*.proto; 
do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done