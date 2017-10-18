package after

import (
	"time"
	"fmt"
	"strconv"
	"regexp"
)

const (
	Regex      = "^((\\+|\\-))?([1-9][0-9]*)\\s?(s|seconds?|m|minutes?|h|hours?|d|days?|w|weeks?)$"
	Day        = time.Hour * 24
	Week       = Day * 7
	Sign       = 2
	Multiplier = 3
	Unit       = 4
	Minus      = "-"
)

var compiledRegex *regexp.Regexp 

type Parser struct {}

func New() *Parser {
	return &Parser{}
}

// Duration returns the time duration calculated from the given input string 
func (p *Parser) Duration(input string) (time.Duration, error) {
	regex, err := p.getRegex()
	if err != nil {
		return 0, err
	}

	matches := regex.FindStringSubmatch(input)
	if matches == nil || len(matches) == 0 {
		return 0, fmt.Errorf("The provided string '%s' is not a valid duration indicator", input)
	}

	multiplier, err := p.getMultiplier(matches[Sign], matches[Multiplier])
	if err != nil {
		return 0, err
	}

	return p.getDuration(matches[Unit], multiplier)
}

// Since returns a point in time after adding (or subtracting) the duration of the input string to a given moment
func (p *Parser) Since(moment time.Time, input string) (time.Time, error) {
	duration, err := p.Duration(input)
	if err != nil {
		return moment, err
	}

	return moment.Add(duration), nil
}

// SinceNow returns a point in time after adding (or subtracting) the duration of the input string from now
func (p *Parser) SinceNow(input string) (time.Time, error) {
	return p.Since(time.Now(), input)
}

func (p *Parser) getMultiplier(sign string, multiplier string) (time.Duration, error) {
	result, err := strconv.Atoi(multiplier)
	if err != nil {
		return 0, err
	}

	if sign == Minus {
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