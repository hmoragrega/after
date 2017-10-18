package after_test

import (
	"testing"
	"time"
	"github.com/hmoragrega/after"
)

func TestItParsePositiveDurationsCorrectly(t *testing.T) {
	parser := after.New()

	valid := map[string]time.Duration{
		"1s":          time.Second,
		"+1 second":   time.Second,
		"20seconds":   time.Second * 20,
		"1m":          time.Minute,
		"+1 minute":   time.Minute,
		"20 minutes":  time.Minute * 20,
		"1h":          time.Hour,
		"+1 hour":     time.Hour,
		"20hours":     time.Hour * 20,
		"1d":          time.Hour * 24,
		"+1 day":      time.Hour * 24,
		"20 days":     time.Hour * 24 * 20,
		"1w":          time.Hour * 24 * 7,
		"+1 week":     time.Hour * 24 * 7,
		"20 weeks":    time.Hour * 24 * 7 * 20,
	}

	for input, expected := range valid {
		duration, err := parser.Duration(input);
		if err != nil {
			t.Errorf("The input '%s' should be valid, error found: %s", input, err)
		}

		if expected != duration {
			t.Errorf("The parsed input '%s' does not macth the expected duration. got %s - expected %s", input, duration, expected)
		}
	}
}

func TestItParseNegativeDurationsCorrectly(t *testing.T) {
	parser := after.New()

	valid := map[string]time.Duration{
		"-1s":         -time.Second,
		"-1 second":   -time.Second,
		"-20seconds":  -time.Second * 20,
		"-1m":         -time.Minute,
		"-1 minute":   -time.Minute,
		"-20 minutes": -time.Minute * 20,
		"-1h":         -time.Hour,
		"-1 hour":     -time.Hour,
		"-20hours":    -time.Hour * 20,
		"-1d":         -time.Hour * 24,
		"-1 day":      -time.Hour * 24,
		"-20 days":    -time.Hour * 24 * 20,
		"-1w":         -time.Hour * 24 * 7,
		"-1 week":     -time.Hour * 24 * 7,
		"-20 weeks":   -time.Hour * 24 * 7 * 20,
	}

	for input, expected := range valid {
		duration, err := parser.Duration(input);
		if err != nil {
			t.Errorf("The input '%s' should be valid, error found: %s", input, err)
		}

		if expected != duration {
			t.Errorf("The parsed input '%s' does not macth the expected duration. got %s - expected %s", input, duration, expected)
		}
	}
}

func TestItFailsOnNonValidStrings(t *testing.T) {
	parser := after.New()

	invalid := []string{
		"1f",  "+1 foo",    // invalid unit  'f - foo'
		"0m",  "+0 minute", // invalid value '0'
		"h",   "hour",      // missing value
		" 1w", "+1   week", // extra whitespace
	}

	for _, input := range invalid {
		if _, err := parser.Duration(input); err == nil {
			t.Errorf("The string '%s' should NOT be valid", input)
		}
	}
}

func TestItAddTheDurationToAGivenMoment(t *testing.T) {
	parser := after.New()

	moment   := time.Now()
	input    := "1m"
	expected := moment.Add(time.Minute)

	result, err := parser.Since(moment, input)
	if err != nil {
		t.Errorf("There was an unexpected error while calculting a point in time: %s", err)
	}

	if result != expected {
		t.Errorf("There result obtained dows not match the expected time. Got %s - expected %s", result, expected)
	}
}

func TestItAddTheDurationToNow(t *testing.T) {
	parser := after.New()

	input    := "1m"
	_, err := parser.SinceNow(input)
	if err != nil {
		t.Errorf("There was an unexpected error while calculting a point in time since now: %s", err)
	}
}