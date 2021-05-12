#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo
  echo "usage: ./bin/deploy.sh <cluster> <version>"
  echo "  ie. ./bin/deploy.sh alpha-apse2-v1 v0.1.30"
  exit 255
fi

cluster=${1}
version=${2}
overlayComponent="_build/${cluster}/asg-node-roller"

./bin/compile.sh $cluster $version

# stackup cfn for component
if [[ -f "${overlayComponent}/cfn/template.yaml" ]]
then
  echo ":cloudformation: deploying cfn for asg-node-roller"
  docker-compose run --rm stackup "asg-node-roller-${cluster}" up -t ${overlayComponent}/cfn/template.yaml
fi

docker-compose run --rm kubectl apply -f ${overlayComponent}
