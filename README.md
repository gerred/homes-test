Small CLI utility to parse valuation data from (potentially duplicated) homes data. Pricing is determined by the last record read from a given date and property id.

No external dependencies outside of the stdlib.

## Install

`go get github.com/gerred/homes-test`

This will install the `homes-test` binary into your `$GOPATH/bin`.

## Run Tests

`go test ./...`

## Run

`homes-test valuations.csv`

Example output:

`
Date,PropertyID,Address,RatesValuation($)
21/01/2015,2,6 Link Road, Wellington,$550,000
21/01/2014,2,6 Link Road, Wellington,$460,000
21/01/2015,3,8 Link Road, Wellington,$540,000
21/01/2015,4,10 Link Road, Wellington,$520,000
21/01/2015,5,12 Link Road, Wellington,$510,000
21/01/2015,1,4 Link Road, Wellington,$500,000
21/01/2014,1,4 Link Road, Wellington,$480,000
`

All property parsing logic is separated into the `properties` package, so it can be used however an implementer wants to use it.

Given this is a code test and not real world, there are some caveats:

* Properties is built out of nested maps. Over a very large file, this would be extremely memory intensive. If working against an actual database or a backend API, this would be chunked out. For printing out records, this doesn't matter as much. In this instance, as well, the backend database would be handling duplicate records.
* Valuation is represented as a string. This saved some code for the purposes of the test. If comparison of values (charting), it'd be easier to have used something like `golang.org/x/text/currency`, which has full CLDR version support and properly handles formatting (and properly handles sub-dollar amounts) beyond just using `int<x>` and manual formatting.