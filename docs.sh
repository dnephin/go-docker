#!/bin/bash

cat <<EOF > go/mobycore/index.html
<!DOCTYPE html><html>
	<head>
	</head>
	<body>
EOF

godoc ./go/mobycore >> index.html

cat <<EOF >> index.html
	</body>
</html>
EOF
