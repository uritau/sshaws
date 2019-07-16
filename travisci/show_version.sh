# !/usr/bin/env bash

version_file="version.go"
version=`grep "version=" $version_file | cut -d'=' -f 2`
old_version=${version//\"}
new_version=`echo $old_version | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g'`

echo "$new_version"
exit 0


