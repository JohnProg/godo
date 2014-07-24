#/bin/bash
. ./config/deploy.cfg
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o dist/linux/godo main.go
scp dist/linux/godo ${serverName}:${remoteFolder}
scp config/config.json ${serverName}:${remoteFolder}/config