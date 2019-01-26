#!/bin/bash

image=eelcoh/service-requests:$1

docker tag service-requests $image && docker push $image



