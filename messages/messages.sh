#!/bin/bash
case $1 in
    start)
        echo "Starting messages web service."
        /root/go/bin/hello &
        ;;
    stop)
        echo "Stopping messages web service."
        sudo kill $(sudo lsof -t -i:8001)
        ;;
    *)
        echo "Messages Web Service."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0