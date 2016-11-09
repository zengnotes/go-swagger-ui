#!/bin/sh
ProjectPath=`pwd`
export GOPATH=$ProjectPath/vendor:$ProjectPath:$GOPATH
if [ -z "$1" ]
then
	echo "place select a version to build...\n\n"
	exit 1
fi

if [ "X$1" = "Xtest" ]
then
	go test -v $2
	exit 0
fi

if [ "X$1" = "Xbench" ]
then
	go test -test.bench=".*"
	exit 0
fi

rm -rf $ProjectPath/bin/*

cd $ProjectPath
echo  $1 "version building..."
##-gcflags "-N -l"
go build -ldflags "-X config.VERSION=$1 "  -o ./bin/swagger-ui ./src/cmd
echo  $1 "bin start..."
if [ "X$1" = "XPro" ]
then
cd $ProjectPath/bin
cp ../src/cmd/api.json  ./swagger.json
./swagger-ui
else
cd $ProjectPath/bin
cp ../src/cmd/api.json ./swagger.json
./swagger-ui
fi

#调试的时候使用
echo  $1 "version build finish"