#!/usr/bin/env bash
echo "Hello, I am inside 'script.sh' file."

now=$(date)
echo "Current date: $now"
echo "Current date: $(date)"

echo "View the current default target using the systemctl: $systemctl get-default"

