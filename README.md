# challenge-2000
A tech challenge for Mercado Libre

## Risk Criteria
The risk is ranked in a scale from `0` (No Risk) to `100` (Inmediate action needed)

### Option 1
Initially I thought about this criteria:

* If the user has access to any sensitive DB: return 100
* If the user has `total` access to any critical app: return 75
* If the user access is other than `total`: return 50
* If the user is active: return 25
* If the user is inactive: return 0

The problem with this criteria is that it wouldn't properly separate the users into smaller, more handy groups but it would return a lot of users as urgent to address.

### Option 2
Another aproach to solve this problem is a progressive criteria where:

* Calculate the DBs access:
  * The sensitive DBs access is the 60% of the total weight of the criteria.
  * Calculate what percentage of the sensitive DBs the user has access to.
  * Add to the total risk value the weighted DBs: `(DB Access * 60) / Sensitive DBs`
* Calculate Critical Apps access:
  * The critical apps access is the 40% of the total weight of the criteria.
  * Calculate what percentage of critical apps the user has access to.
  * Calculate what percentage of those critical apps have `total` access.
  * Add to the total risk value the weighted accesses: `() / `