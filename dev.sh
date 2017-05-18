#!/usr/bin/env bash
# Attempt to set APP_HOME
# Resolve links: $0 may be a link
PRG="$0"
# Need this for relative symlinks.
while [ -h "$PRG" ] ; do
    ls=`ls -ld "$PRG"`
    link=`expr "$ls" : '.*-> \(.*\)$'`
    if expr "$link" : '/.*' > /dev/null; then
        PRG="$link"
    else
        PRG=`dirname "$PRG"`"/$link"
    fi
done
SAVED="`pwd`"
cd "`dirname \"$PRG\"`" >/dev/null
APP_HOME="`pwd -P`"
cd "$SAVED" >/dev/null

if [ "$1" == "gobuild" ];then
    ps -ef |grep "/tmp/go-build" |grep -v grep |awk '{print $2}'|xargs kill -15 
    sleep 2
    go run $APP_HOME/server.go
else
    nohup watcher -cmd "babel watch_file -s -o watch_parent/../react_module/watch_name " $APP_HOME/resource/react_src  >$APP_HOME/dev.log 2>&1 &
    sleep 1
    nohup watcher -cmd "bash "$APP_HOME"/dev.sh gobuild" -ext go $APP_HOME  >$APP_HOME/dev.log 2>&1 &    
    sleep 1
    tail -100f $APP_HOME/dev.log
    exit
fi


