#!/bin/bash

export DB_USER=root2
export DB_PASS=root2
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=PLAYLISTBOOK

OPTIONS="-config=dbconfig.yml -env development"

case "$1" in
    "new")
    sql-migrate new $OPTIONS $2
    ;;
    "up")
    sql-migrate up $OPTIONS
    ;;
    "redo")
    sql-migrate redo $OPTIONS
    ;;
    "status")
    sql-migrate status $OPTIONS
    ;;
    "down")
    sql-migrate down $OPTIONS
    ;;
    *)
    echo "Usage: $(basename "$0") new {name}/up/status/down"
    exit 1
esac