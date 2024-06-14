#!/bin/bash

while true; do
    OUTPUT=$(./toybox wget http://gstatic.com/generate_204 2>&1)

    if echo "$OUTPUT" | grep -q "Connection reset by peer"; then
        echo "Connection reset by peer detected. Restarting data service..."
        
        adb shell svc data disable
        sleep 2
        adb shell svc data enable
    fi
    
    sleep 5
done
