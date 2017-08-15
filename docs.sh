#!/bin/bash

domain=golang.docker.io
gitrepo="https://github.com/tiborvass/docker-sdk-go"
mobycore=docker

cat <<EOF | tee index.html > $mobycore/index.html
<!DOCTYPE html><html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="$domain git $gitrepo">
		<meta name="go-source" content="$domain $gitrepo/ $gitrepo/tree/master{/dir} $gitrepo/blob/master{/dir}/{file}#L{line}">
	</head>
	<body>
EOF

git stash save
commit=$(git rev-parse --abbrev-ref HEAD)
git checkout master
godoc -html $domain/$mobycore > /tmp/index.html
git checkout $commit
git stash pop
cat /tmp/index.html >> $mobycore/index.html

cat <<EOF >> index.html
		<h1>Docker Devkit</h1>
		<ul>
			<li>Go</li>
		</ul>
		<h2>Go</h2>
		<ul>
			<li><a href="$mobycore">Moby Core</a></li>
		</ul>
EOF

cat <<EOF | tee -a index.html >> $mobycore/index.html
	</body>
</html>
EOF
