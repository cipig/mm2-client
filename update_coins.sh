#!/bin/bash

# Use this as a daily cron job

cd /home/admin/mm2-client/mm2_tools_server
# backup existing coins_config.json
cp /home/admin/mm2-client/coins_config.json /home/admin/mm2-client/coins_config.json.bk
# Get new coins_config.json from source
/usr/bin/curl https://raw.githubusercontent.com/KomodoPlatform/coins/master/utils/coins_config.json -o /home/admin/mm2-client/coins_config.json
# Check new coins_config.json file is valid
data=$(cat /home/admin/mm2-client/coins_config.json | wc -l)
# Revert to old file if new file is invalid
if [[ "$data" -eq 1 ]]; then
    /home/admin/mm2-client/coins_config.json.bk /home/admin/mm2-client/coins_config.json
fi
# Rebuild service
/usr/local/go/bin/go build -o /home/admin/mm2-client/prices_komodo_earth /home/admin/mm2-client/cmd/mm2_tools_server/mm2_tools_server.go
# restart service
/bin/systemctl restart prices-komodo-earth
# Print date for logs
now=$(date +%m-%d-%Y)
echo "prices updated at ${now}"
