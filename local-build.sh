#!/bin/bash

mkdir -p build
cd build

curl -L https://github.com/OpenListTeam/OpenList-Frontend/releases/latest/download/openlist-frontend-dist-$(cat ../upstream_version).tar.gz -o dist.tar.gz
mkdir -p dist
tar -zxvf dist.tar.gz -C dist
curl -L https://cdn.oplist.org/gh/OpenListTeam/Logo@main/logo.svg -o dist/images/logo.svg
rm -rf ../public/dist
mv -f dist ../public
rm -rf dist.tar.gz

cd ..

appName="openlist"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor=$(git show -s --format='format:%aN <%ae>' HEAD)
gitCommit=$(git log --pretty=format:"%h" -1)
#version=$(git describe --long --tags --dirty --always)
version="asogii-$(cat upstream_version)-$(git log -1 --pretty=format:%h)"
webVersion=$(wget -qO- -t1 -T2 "https://api.github.com/repos/OpenListTeam/OpenList-Frontend/releases/latest" | grep "tag_name" | head -n 1 | awk -F ":" '{print $2}' | sed 's/\"//g;s/,//g;s/ //g')
ldflags="\
-w -s \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.BuiltAt=$builtAt' \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.GoVersion=$goVersion' \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.GitAuthor=$gitAuthor' \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.GitCommit=$gitCommit' \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.Version=$version' \
-X 'github.com/OpenListTeam/OpenList/v4/internal/conf.WebVersion=$webVersion' \
"
go build -ldflags="$ldflags" -x .

mv openlist build/.
