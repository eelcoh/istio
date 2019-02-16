#!/bin/bash

image=eelcoh/poc.istio.tokens:$1

docker tag poc.istio.tokens $image && docker push $image



