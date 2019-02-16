#!/bin/sh

echo "compiling ... "
CGO_ENABLED=0 GOOS=linux go build -o ./bin/activity -a -tags netgo -ldflags '-w' ..


if [[ $? -eq 0 ]]; then

	echo "build docker file"
	docker build . -t activity:latest
	
	if [[ $? -eq 0 ]]; then
		echo "done"
	fi

fi

