# !/usr/bin/env bash

version_file="pkg/cmd/version.go"
version=`grep "version = " $version_file | cut -d'=' -f 2 | sed 's/"//g'`

echo "$version"
exit 0
