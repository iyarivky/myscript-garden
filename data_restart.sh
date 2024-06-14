#!/bin/bash

while true; do
    OUTPUT=$(timeout 3 ./toybox wget http://gstatic.com/generate_204 2>&1)

    if [ $? -ne 0 ] || [ -z "$OUTPUT" ]; then
        echo "No response or error detected. Restarting data service..."
        
        svc data disable
        sleep 2
        svc data enable
    fi
    
    sleep 5
done
