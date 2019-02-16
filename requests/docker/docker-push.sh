#!/bin/bash

image=eelcoh/poc.istio.requests:$1

docker tag poc.istio.requests $image && docker push $image



