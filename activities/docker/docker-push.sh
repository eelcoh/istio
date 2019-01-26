#!/bin/bash

image=eelcoh/activity-log:$1

docker tag token $image && docker push $image



