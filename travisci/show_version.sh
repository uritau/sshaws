# !/usr/bin/env bash

version_file="cmd/version.go"
version=`grep "version=" $version_file | cut -d'=' -f 2`
version=${version//\"}

echo "$version"
exit 0


