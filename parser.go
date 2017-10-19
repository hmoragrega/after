package after

import (
	"time"
	"fmt"
	"strconv"
	"regexp"
)

const (
	Regex = "^((\\+|\\-))?([1-9][0-9]*)\\s?(s|seconds?|m|minutes?|h|hours?|d|days?|w|weeks?)$"
	Day   = time.Hour * 24
	Week  = Day * 7
)

var compiledRegex *regexp.Regexp 

// Parser is the service that will provide the package functionality
type Parser struct {}

// New will return pointer to a new Parser
func New() *Parser {
	return &Parser{}
}

// Duration returns a time.Duration object with the equivalent duration representation of the input
//
// The input string must contain a time unit and a signed multiplier ('+' sign by default if omitted)
//
// The accepted time unit are:
// - `s`, `second` or `seconds`
// - `m`, `minute` or `minutes`
// - `h`, `hour`   or `hours`
// - `d`, `day`    or `days`
// - `w`, `week`   or `weeks`
//
// Examples:
// - `10s`: after ten seconds
// - `+1 minute`: after one minute
// - `2 hours`: after two hours
// - `-1 day`: minus one day
// - `-2w`: minus two weeks
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
	duration, err := p.Duration(duration)
	if err != nil {
		return moment, err
	}

	return moment.Add(duration), nil
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
	switch unit[:1] {
	case "s":
		return time.Second * multiplier, nil
	case "m":
		return time.Minute * multiplier, nil
	case "h":
		return time.Hour * multiplier, nil
	case "d":
		return Day * multiplier, nil
	case "w":
		return Week * multiplier, nil
	}

	return 0, fmt.Errorf("The duration unit '%s' is not supported", unit)
}

func (p *Parser) getRegex() (*regexp.Regexp, error) {
	if compiledRegex != nil {
		return compiledRegex, nil
	}

	regex, err := regexp.Compile(Regex)

	return regex, err
}