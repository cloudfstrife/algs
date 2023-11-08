#!/bin/bash

CURRENT_PATH=$(cd `dirname $0`; pwd)
URL=https://algs4.cs.princeton.edu/code/algs4-data.zip
FILE_NAME=${URL##*/}

wget $URL -O $CURRENT_PATH/$FILE_NAME

unzip $CURRENT_PATH/$FILE_NAME -d $CURRENT_PATH

rm $CURRENT_PATH/$FILE_NAME
