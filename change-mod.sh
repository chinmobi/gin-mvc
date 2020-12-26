#!/usr/bin/env bash

cd $(dirname $0)

newmod=$1

if [ "x$newmod" = "x" ]; then
  echo '  Usage:'
  echo '    ./change-mod.sh <your-new-module-name>'
  exit 1
fi

new_proj_name=$(echo "$newmod" | awk -F '/' '{print $3}')

if [ "x$new_proj_name" = "x" ]; then
  echo '  Invalid module name pattern, using like:'
  echo '    ./change-mod.sh <github.com>/<account-name>/<project-name>'
  exit 1
fi

oldmod=$(sed -n '1p' go.mod | awk -F ' ' '{print $2}')

if [ "$newmod" = "$oldmod" ]; then
  exit 0
fi

old_proj_name=$(echo "$oldmod" | awk -F '/' '{print $3}')


### Begin replacing
echo "  Change module name from: ${oldmod} to ${newmod} ..."

### For verifying
#old_mod=$(echo "${oldmod}" | sed 's|/|\\/|g')
#new_mod=$(echo "${newmod}" | sed 's|/|\\/|g')

### Replace the *.go files' imported package name
find . -depth -name "*.go" | xargs -I {} sed -i "s|$oldmod|$newmod|g" {}

### Verify
#find . -depth -name "*.go" | xargs -I {} sed -n "/$old_mod/p" {}
#find . -depth -name "*.go" | xargs -I {} sed -n "/$new_mod/p" {}

### Replace the go.mod file's module name
sed -i "s|$oldmod|$newmod|" ./go.mod

### End
