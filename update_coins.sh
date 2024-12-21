#!/bin/bash

FOLDER="/home/admin/mm2-client"
rm ${FOLDER}/coins_config.json
rm ${FOLDER}/coins
wget -P ${FOLDER} https://raw.githubusercontent.com/KomodoPlatform/coins/master/utils/coins_config.json
wget -P ${FOLDER} https://raw.githubusercontent.com/KomodoPlatform/coins/master/coins
