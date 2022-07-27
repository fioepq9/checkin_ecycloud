#!/bin/bash
docker run --mount type=bind,source=$PWD/log,target=/workspace/log -d checkin_ecycloud
