## ASG Node Roller

The update methodology is simple:

- Find the node longer than (By default TTL 48 hours) in ASG (every half an hour).
- Terminate the selected node via ASG (As drain node logic has been covered by AWS ManagedNodeGroups)
