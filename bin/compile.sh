#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo
  echo "usage: ./bin/compile.sh <cluster> <version>"
  echo "  ie. ./bin/compile.sh alpha-apse2-v1 v0.1.2"
  exit 255
fi


cluster=${1}
version=${2}
configfile=envs/${cluster}.yaml

function prepareBuildFolder() {
  echo "preparing $cluster template"
  rm -rf _build
  mkdir -p _build/${cluster}
}

function castTemplate() {
  echo -n "casting $cluster templates"
  echo "{version: $version}" | docker-compose run --rm gomplate \
    --left-delim='<<' --right-delim='>>' \
    --input-dir ./templates/ \
    --output-dir=_build/${cluster}/ \
    --context config=${configfile} \
    -d data=stdin:///${version}-config.json
}

if [[ ! -f "${configfile}" ]]
then
  echo "error: configfile does not exist: ${configfile}"
  exit 1
fi

prepareBuildFolder
castTemplate

echo "node!"
