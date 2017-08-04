#!/bin/bash

mobycore=go/docker/mobycore

cat <<EOF | tee index.html > $mobycore/index.html
<!DOCTYPE html><html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="tiborvass.github.io/devkit git https://github.com/tiborvass/devkit">
		<meta name="go-source" content="tiborvass.github.io/devkit https://github.com/tiborvass/devkit/ https://github.com/tiborvass/devkit/tree/master{/dir} https://github.com/tiborvass/devkit/blob/master{/dir}/{file}#L{line}">
	</head>
	<body>
EOF

git stash save
commit=$(git rev-parse --abbrev-ref HEAD)
git checkout master
godoc -html tiborvass.github.io/devkit/$mobycore > /tmp/index.html
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
			<li><a href="go/mobycore">Moby Core</a></li>
		</ul>
EOF

cat <<EOF | tee -a index.html >> $mobycore/index.html
	</body>
</html>
EOF
