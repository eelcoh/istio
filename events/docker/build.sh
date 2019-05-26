#!/bin/sh

echo "compiling ... "
CGO_ENABLED=0 GOOS=linux go build -o ./bin/events -a -tags netgo -ldflags '-w' ..


if [[ $? -eq 0 ]]; then

	echo "build docker file"
	docker build . -t poc.istio.events:latest
	
	if [[ $? -eq 0 ]]; then
		echo "done"
	fi

fi

