package after

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	// Regex used to match the duration
	Regex = "^((\\+|\\-))?([1-9][0-9]*)\\s?(ms|milliseconds?|s|seconds?|m|minutes?|h|hours?|d|days?|w|weeks?)$"
	// Day represents 24 hours durations
	Day = time.Hour * 24
	// Week represents 7 days
	Week = Day * 7
)

var compiledRegex *regexp.Regexp

var unitMap = map[string]time.Duration{
	"ms":           time.Millisecond,
	"millisecond":  time.Millisecond,
	"milliseconds": time.Millisecond,
	"s":            time.Second,
	"second":       time.Second,
	"seconds":      time.Second,
	"m":            time.Minute,
	"minute":       time.Minute,
	"minutes":      time.Minute,
	"h":            time.Hour,
	"hour":         time.Hour,
	"hours":        time.Hour,
	"d":            Day,
	"day":          Day,
	"days":         Day,
	"w":            Week,
	"week":         Week,
	"weeks":        Week,
}

// Parser is the service that will provide the package functionality
type Parser struct{}

// New will return pointer to a new Parser
func New() *Parser {
	return &Parser{}
}

// Duration returns a time.Duration object with the equivalent duration representation of the input
//
// The input string must contain a time unit and a signed multiplier ('+' sign by default if omitted)
//
// The accepted time unit are:
//  - `ms`, `millisecond` or `milliseconds`
//  - `s`, `second` or `seconds`
//  - `m`, `minute` or `minutes`
//  - `h`, `hour`   or `hours`
//  - `d`, `day`    or `days`
//  - `w`, `week`   or `weeks`
//
// Examples:
//  - `10s`: after ten seconds
//  - `+1 minute`: after one minute
//  - `2 hours`: after two hours
//  - `-1 day`: minus one day
//  - `-2w`: minus two weeks
func (p *Parser) Duration(duration string) (time.Duration, error) {
	const (
		Sign       = 2
		Multiplier = 3
		Unit       = 4
	)

	regex, err := p.getRegex()
	if err != nil {
		return 0, err
	}

	matches := regex.FindStringSubmatch(duration)
	if matches == nil || len(matches) == 0 {
		return 0, fmt.Errorf("The provided string '%s' is not a valid duration indicator", duration)
	}

	multiplier, err := p.getMultiplier(matches[Sign], matches[Multiplier])
	if err != nil {
		return 0, err
	}

	return p.getDuration(matches[Unit], multiplier)
}

// SinceNow returns a time.Time object that represents the current point in time plus (or minus) the specified duration
func (p *Parser) SinceNow(duration string) (time.Time, error) {
	return p.Since(time.Now(), duration)
}

// Since returns a time.Time object that represents the given point in time plus (or minus) the specified duration
func (p *Parser) Since(moment time.Time, duration string) (time.Time, error) {
	result, err := p.Duration(duration)
	if err != nil {
		return moment, err
	}

	return moment.Add(result), nil
}

func (p *Parser) getMultiplier(sign string, multiplier string) (time.Duration, error) {
	result, err := strconv.Atoi(multiplier)
	if err != nil {
		return 0, err
	}

	if sign == "-" {
		result *= -1
	}

	return time.Duration(result), nil
}

func (p *Parser) getDuration(unit string, multiplier time.Duration) (time.Duration, error) {
	duration, found := unitMap[unit]
	if !found {
		return 0, fmt.Errorf("The duration unit '%s' is not supported", unit)
	}

	return duration * multiplier, nil
}

func (p *Parser) getRegex() (*regexp.Regexp, error) {
	if compiledRegex != nil {
		return compiledRegex, nil
	}

	regex, err := regexp.Compile(Regex)

	return regex, err
}
