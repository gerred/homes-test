Test #2.
Instead of inserting the last record, make sure that no duplicates are entered at all. ie, if there is a duplicate record, do not insert any rows.

Test #3.
Modify the codebase to split the validated (non-duplicate) records into chunks, and 'process' each chunk via its own goroutine. Process is defined as: 
 1. Filter out cheap properties (anything under 400k)
 2. Filter out properties that are avenues, crescents, or places (AVE, CRES, PL) cos those guys are just pretentious...
 3. Filter out every 10th property (to keep our users on their toes!)

Combine the results of the 'processes' back into a single list, and print the list.

Small CLI utility to parse valuation data from (potentially duplicated) homes data. Pricing is determined by the last record read from a given date and property id.

No external dependencies outside of the stdlib.

## Install

`go get github.com/gerred/homes-test`

NOTE: If you don't have it already, add this to your `~/.gitconfig`:

```
[url "git@github.com:"]
    insteadOf = https://github.com/
```

This will enable you to `go get` private repositories.

This will install the `homes-test` binary into your `$GOPATH/bin`.

## Run Tests

`go test ./...`

## Run

`homes-test valuations.csv`

Example output:

```
Date,PropertyID,Address,RatesValuation($) 
21/01/2015,2,6 Link Road, Wellington,$550,000 
21/01/2014,2,6 Link Road, Wellington,$460,000 
21/01/2015,3,8 Link Road, Wellington,$540,000 
21/01/2015,4,10 Link Road, Wellington,$520,000 
21/01/2015,5,12 Link Road, Wellington,$510,000 
21/01/2015,1,4 Link Road, Wellington,$500,000 
21/01/2014,1,4 Link Road, Wellington,$480,000
```

## Data cardinality / architecture

In this exercise, the data structure was chosen based on the primary and secondary fields: propertyID and date, respectively. This is following the assumption that the "home" itself is the top level data structure, and it's changes over time are representative of that. Obviously, a relational database would be better in production because there may be questions that go past a single-home case (pricing trends of an area, for example), but we're cleaning CSVs for this test, not building a data warehouse.

## Notes

All property parsing logic is separated into the `properties` package, so it can be used however an implementer wants to use it.

Given this is a code test and not real world, there are some caveats:

* This library is not concurrency safe at this time, mostly because it wasn't necessary for this exercise. Easiest way would be to add a sync.Mutex on Properties. However, that's slow, and I don't particularly think CRDT's are ready to go mainstream, so I'd rely on whatever DB I was using to handle transactions and do connection pooling on that.
* Properties is built out of nested maps. Over a very large file, this would be extremely memory intensive. If working against an actual database or a backend API, this would be chunked out. For printing out records, this doesn't matter as much. In this instance, as well, the backend database would be handling duplicate records.
* Valuation is represented as a string. This saved some code for the purposes of the test. If comparison of values is required (charting, analytics), it'd be easier to have used something like `golang.org/x/text/currency`, which has full CLDR version support and properly handles formatting (and properly handles sub-dollar amounts) beyond just using `int<x>` and manual formatting.
