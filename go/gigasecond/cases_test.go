package gigasecond

// Source: exercism/problem-specifications
// Commit: 5506bac gigasecond: Apply new "input" policy
// Problem Specifications Version: 1.1.0

// Add one gigasecond to the input.
var addCases = []struct {
	description string
	in          string
	want        string
}{
	{
		"date only specification of time",
		"2043-01-01T01:46:40",
		"2043-01-01T01:46:40",
	},
	{
		"second test for date only specification of time",
		"2009-02-19T01:46:40",
		"2009-02-19T01:46:40",
	},
	{
		"third test for date only specification of time",
		"1991-03-27T01:46:40",
		"1991-03-27T01:46:40",
	},
	{
		"full time specified",
		"2046-10-02T23:46:40",
		"2046-10-02T23:46:40",
	},
	{
		"full time with day roll-over",
		"2046-10-03 01:46:39",
		"2046-10-03T01:46:39",
	},
}
