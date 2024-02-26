#!/bin/bash

# Navigate to the directory where your application binary is located
# cd /app

# Start the Go service
nohup $(dirname "$0")/app/hotel-rates-api &