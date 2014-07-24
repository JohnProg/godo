#/bin/bash
. ./config/deploy
scp ./godo $serverName:$remoteFolder
scp ./config/config.json $serverName:$remoteFolder/config