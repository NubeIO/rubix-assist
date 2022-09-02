#!/bin/bash

# Console colors
DEFAULT="\033[0m"
GREEN="\033[32m"
RED="\033[31m"

PRODUCTION=false

help() {
  echo "Service commands:"
  echo -e "   ${GREEN}--prod | --production: add these suffix to start production"
}

parseCommand() {
  for i in "$@"; do
    case ${i} in
    -h | --help)
      help
      exit 0
      ;;
    --prod | --production)
      PRODUCTION=true
      ;;
    *)
      echo -e "${RED}Unknown options ${i}  (-h, --help for help)${DEFAULT}"
      ;;
    esac
  done
}

parseCommand "$@"

dir=$(pwd)
echo -e "${GREEN}Current working directory is: $dir${DEFAULT}"

if [ ${PRODUCTION} == true ]; then
  echo -e "${GREEN}We are running in production mode!${DEFAULT}"
else
  echo -e "${GREEN}We are running in development mode!${DEFAULT}"
fi

cd $dir

if [ ${PRODUCTION} == true ]; then
  go run app.go -g /data/rubix-assist -d data --prod
else
  go run app.go
fi
