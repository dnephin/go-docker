#!/bin/bash

set -e

ENGINE_BRANCH=master
NOTARY_BRANCH=9ae66476d611af5df9b7efb09f830e6fc16e8a65
# Distribution is taken from engine's vendor.conf

# branch that needs docker/docker be rewritten to moby/moby-core. Following was used on docker-ce:
# 
# git grep -l "github.com/docker/docker" | xargs sed -i '' -E 's,github.com/docker/docker(["/ ]|$),github.com/moby/moby-core\1,g' && git status -s | cut -d' ' -f3- | grep '.*\.go$' | xargs gofmt -w -s

# Hint: to get list of exported identifiers:
# go doc | tail -n +3 | sed -E 's,^ *[^ ]+ ([^ (]+).*,\1,g'

domain=golang.docker.com
urlpath=go-docker
importpath="$domain/$urlpath"
package=docker

sed=$(which gsed) || sed=$(which sed)
dir=$(pwd)
rm -rf *.go api registry notary
tmp=/tmp/testing

set -x
cd "$tmp"

[ ! -d docker ] && git clone --depth 1 -b "$ENGINE_BRANCH" https://github.com/docker/docker
[ ! -d notary ] && git clone https://github.com/docker/notary
pushd notary
git checkout $NOTARY_BRANCH
popd


pushd docker
distribution_commit=$(grep 'github.com/docker/distribution' vendor.conf | head -1 | cut -d' ' -f2)
popd

[ ! -d distribution ] && git clone https://github.com/docker/distribution
pushd distribution
git checkout $distribution_commit
popd

pushd docker
for folder in api client; do
	find "$folder" -name '*.go' -type f -exec sed -i'' -E 's#github.com/docker/docker/api(/?)#'"${importpath}"'/api\1#g' {} \;
	find "$folder" -name '*.go' -type f -exec sed -i'' -E 's#github.com/docker/docker/client(/?)#'"${importpath}"'\1#g' {} \;
done
cp client/*.go "$dir/"
cp -rf api "$dir/"
rm -rf "$dir/api/server" "$dir/swarm_get_unlock_key_test.go" "$dir/api/errdefs" "$dir/api/templates" "$dir/api/types/backend"
popd

pushd "$dir"
find . -name '*.go' -depth 1 -print | xargs $sed -i'' -E 's,^package client\b,package '"${package}"' // import "'${importpath}'",g'
find . -name '*.go' -depth 1 -print | xargs $sed -i'' -E 's,^Package client\b,Package '"${package}"',g'
sed -i'' -E 's#client(\.NewEnvClient\(\))#docker\1#g' client.go
popd



# notary
pushd notary
find client -name '*.go' -type f -exec sed -i'' -E 's#github.com/docker/notary/client(/?)#'"${importpath}"'/notary\1#g' {} \;
find client -name '*.go' -depth 1 -print | xargs $sed -i'' -E 's,^package client\b,package notary // import "'${importpath}'/notary",g'
mkdir -p "$dir/notary"
cp -rf client/{changelist,*.go} "$dir/notary/"
popd



# registry
pushd distribution
find registry/client -name '*.go' -type f -exec sed -i'' -E 's#github.com/docker/distribution/registry/client(/?)#'"${importpath}"'/registry\1#g' {} \;
find registry/client -name '*.go' -depth 1 -print | xargs $sed -i'' -E 's,^package client\b,package registry // import "'${importpath}'/registry",g'
mkdir -p "$dir/registry"
cp -rf registry/client/* "$dir/registry/"
popd

cd "$dir"
# reset README.md
git checkout README.md

function strip_doc() {
	tail -n +$(grep -n '^package ' "$1" | cut -d: -f1) "$1" > "$1".new
	mv "$1".new "$1"
}

# replace documentation
strip_doc client.go
cp scripts/files/root_doc.go doc.go
cp scripts/files/notary_doc.go notary/doc.go
cp scripts/files/registry_doc.go registry/doc.go
