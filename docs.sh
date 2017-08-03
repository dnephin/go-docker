#!/bin/bash

STABLE=17.06

cat <<EOF > go/mobycore/index.html
<!DOCTYPE html><html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="tiborvass.github.io/devkit/go/mobycore git https://github.com/tiborvass/devkit">
		<meta name="go-source" content="tiborvass.github.io/devkit/go/mobycore https://github.com/tiborvass/devkit/ https://github.com/tiborvass/devkit/tree/17.06{/dir} https://github.com/tiborvass/devkit/blob/17.06{/dir}/{file}#L{line}">
	</head>
	<body>
EOF

git stash save
commit=$(git rev-parse HEAD)
git checkout master
godoc ./go/mobycore > /tmp/index.html
git checkout $commit
git stash pop
cat /tmp/index.html >> go/mobycore/index.html

cat <<EOF >> go/mobycore/index.html
	</body>
</html>
EOF
