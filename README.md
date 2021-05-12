## ASG Node Roller

The update methodology is simple:

- Find the node in cluster longer than 48 hours (By default TTL is 48 hours).
- Terminate the selected node via ASG (As drain node logic has been covered by AWS ManagedNodeGroups)

## How to deploy

- ./bin/release.sh v0.2.1
- ./bin/deploy.sh alpha-apse2-v1 v0.2.1


## How to compile
- ./bin/compile.sh alpha-apse2-v1 v0.2.1
