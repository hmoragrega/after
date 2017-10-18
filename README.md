# After
A Go lang micro library to parse english future or past time events to Go native time durations

## Examples
* `10s`: after ten seconds
* `+1 minute`: after one minute
* `2 hours`: after two hours
* `-1 day`: minus one day
* `-2w`: minus two weeks

## Usage
```go
parser := after.New()

// "Duration" returns a time.Duration object with the equivalent of the input
anHour, err := parser.Duration("1 hour")

// "SinceNow" returns a time.Time object that represents the current point in time plus the specified input
inTenMinutes, err := parser.SinceNow("10 minutes")

// "Since" returns a time.Time object that represents the given point in time plus the specified input
nowAgain, err := parser.Since(inTenMinutes, "-10 minutes")
```

## Available units
The available time units are
* `s`, `second` or `seconds`
* `m`, `minute` or `minutes`
* `h`, `hour`   or `hours`
* `d`, `day`    or `days`
* `w`, `week`   or `weeks`

## Multiplier
It's the number that will multiply the time unit:
 * It **must** start with 1 to 9
 * It can be signed, both positive or negative.
 * Omitting the sign is equivalent to using a plus sign 

### Validation
You can use this regular expression to validate your inputs:
 ```
 ^((\\+|\\-))?([1-9][0-9]*)\\s?(s|seconds?|m|minutes?|h|hours?|d|days?|w|weeks?)$
 ```

## Scope
The scope of the library is small on purpose, if you are looking for a more full-fledged solution check out [olebedev's _when_](https://github.com/olebedev/when) 
