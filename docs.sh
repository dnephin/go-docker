#!/bin/bash

cat <<EOF > go/mobycore/index.html
<!DOCTYPE html><html>
	<head>
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
