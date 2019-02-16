#!/bin/bash

image=eelcoh/poc.istio.messages:$1

docker tag poc.istio.messages $image && docker push $image



