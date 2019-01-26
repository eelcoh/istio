#!/bin/bash

image=eelcoh/token:$1

docker tag token $image && docker push $image



