#!/bin/bash

#image=docker-registry-default.127.0.0.1.nip.io:80/istio-poc/activity:$1
# docker push docker-registry-default.127.0.0.1.nip.io:80/myproject/alpine


image=eelcoh/poc.istio.activities:$1

docker tag poc.istio.activities:latest $image

	if [[ $? -eq 0 ]]; then
		docker push $image
	fi
# docker tag activity $image && docker push $image



