# !/usr/bin/env bash

version_file="pkg/cmd/version.go"
version=`grep "version = " $version_file | cut -d'=' -f 2`
version=`grep "version = " $version_file | cut -d'=' -f 2 | sed 's/"//g'`

old_version=version
new_version=`echo $old_version | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g'`

echo "New version is $new_version"

git remote set-url origin git@github.com:uritau/sshaws.git
git checkout master
sed -i "s/version=\"$old_version\"/version=\"$new_version\"/g" $version_file
git add .
git commit -m "Bump version"
git tag $NEW_TAG &> /dev/null
git push && git push --tags  &> /dev/null
exit 0
