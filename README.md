## ASG Node Roller

The update methodology is simple:

- Find the node in cluster longer than 48 hours (By default TTL is 48 hours).
- Terminate the selected node via ASG (As drain node logic has been covered by AWS ManagedNodeGroups)
