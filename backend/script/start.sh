#!/bin/sh
set -e  # Exit immediately if a command exits with a non-zero status
    /root/database.sh 
    /root/migrate.sh
exec /root/main