#!/bin/bash

docker build -t dougefresh/presentation-serverless:latest .
docker push  dougefresh/presentation-serverless:latest
