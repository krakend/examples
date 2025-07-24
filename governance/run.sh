#!/bin/bash

CURDIR=$(pwd)
cd redis
. ./start.sh
cd $CURDIR
docker compose up -d
