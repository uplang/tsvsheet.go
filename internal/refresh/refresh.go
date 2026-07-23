// Package refresh models an auto-refresh cadence for clock-dependent
// (TODAY/NOW/ISNOW) cells. A cadence is either a fixed interval or an isnow
// date/time pattern (tsvsheet/isnow) whose next matching instant drives the next
// refresh — so the view updates on a schedule ("every 30 min, 9-to-5, on
// weekdays") rather than a dumb fixed tick. Both frontends (serve, tui) consume
// a Next; the CLI builds one from the --refresh-interval flag.
package refresh

import (
	"time"

	isnow "github.com/tsvsheet/go-isnow"
)

// Next returns the delay from now until the next refresh; a non-positive result
// means no further refresh (the caller stops polling/ticking). A nil Next is a
// disabled cadence.
type Next func(now time.Time) time.Duration

// Every returns a Next with a fixed cadence: it always waits interval.
func Every(interval time.Duration) Next {
	return func(time.Time) time.Duration { return interval }
}

// pattern returns a Next driven by an isnow pattern: it waits until the
// pattern's next matching instant, or stops when there is none.
func pattern(p isnow.Pattern) Next {
	return func(now time.Time) time.Duration {
		at, ok := p.Next(now)
		if !ok {
			return 0
		}
		return at.Sub(now)
	}
}

// Spec is a refresh-cadence specification: a Go duration ("30s", "0s") or an
// isnow pattern ("M-F +[30mn] >=9 <=17").
type Spec string

// Union combines the cadence specs of every volatile(…) cell into one Next that
// fires at the soonest next refresh across them. An empty spec — a volatile()
// with no schedule of its own — uses def, the frontend's default cadence. The
// result is nil when no spec yields a cadence. A malformed pattern is an error.
func Union(specs []Spec, def Next) (Next, error) {
	nexts := make([]Next, 0, len(specs))
	for _, spec := range specs {
		next, err := resolveSpec(spec, def)
		if err != nil {
			return nil, err
		}
		if next != nil {
			nexts = append(nexts, next)
		}
	}
	return soonest(nexts), nil
}

// resolveSpec parses one spec, substituting def for the empty (default) spec.
func resolveSpec(spec Spec, def Next) (Next, error) {
	if spec == "" {
		return def, nil
	}
	return Parse(spec)
}

// soonest returns a Next yielding the smallest positive delay across nexts, or 0
// (stop) when none is due; nil when there are no cadences.
func soonest(nexts []Next) Next {
	if len(nexts) == 0 {
		return nil
	}
	return func(now time.Time) time.Duration {
		best := time.Duration(0)
		for _, next := range nexts {
			best = earlier(best, next(now))
		}
		return best
	}
}

// earlier returns the sooner of two delays, treating 0 as "none yet" and
// ignoring non-positive (stopped) cadences.
func earlier(best, d time.Duration) time.Duration {
	if d > 0 && (best == 0 || d < best) {
		return d
	}
	return best
}

// Parse builds a Next from spec. An empty spec, or a non-positive duration, is a
// disabled cadence (nil). A malformed pattern is an error.
func Parse(spec Spec) (Next, error) {
	if spec == "" {
		return nil, nil
	}
	if d, err := time.ParseDuration(string(spec)); err == nil {
		if d <= 0 {
			return nil, nil
		}
		return Every(d), nil
	}
	p, err := isnow.Parse(isnow.PatternText(spec))
	if err != nil {
		return nil, err
	}
	return pattern(p), nil
}
