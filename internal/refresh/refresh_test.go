package refresh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tsvsheet/tsvsheet.go/internal/refresh"
)

// mon1005 is a fixed instant: Monday 2026-01-05 10:05 UTC.
var mon1005 = time.Date(2026, 1, 5, 10, 5, 0, 0, time.UTC)

func TestEvery(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 5*time.Second, refresh.Every(5*time.Second)(mon1005))
}

func TestParse_DurationAndDisabled(t *testing.T) {
	t.Parallel()

	off, err := refresh.Parse("")
	require.NoError(t, err)
	assert.Nil(t, off) // empty → disabled

	zero, err := refresh.Parse("0s")
	require.NoError(t, err)
	assert.Nil(t, zero) // non-positive → disabled

	fixed, err := refresh.Parse("30s")
	require.NoError(t, err)
	assert.Equal(t, 30*time.Second, fixed(mon1005))
}

func TestParse_IsnowPattern(t *testing.T) {
	t.Parallel()

	// "+[30mn]" fires every 30 minutes; the next from 10:05 is 10:30 (25m).
	every30, err := refresh.Parse("+[30mn]")
	require.NoError(t, err)
	assert.Equal(t, 25*time.Minute, every30(mon1005))

	// A pattern whose bound is in the past has no next match → disabled tick.
	past, err := refresh.Parse("12 <2016")
	require.NoError(t, err)
	assert.Equal(t, time.Duration(0), past(mon1005))
}

func TestParse_MalformedPattern(t *testing.T) {
	t.Parallel()

	_, err := refresh.Parse("garbage!!!")
	require.Error(t, err)
}

func TestUnion(t *testing.T) {
	t.Parallel()

	def := refresh.Every(time.Second)

	// No specs → nil (no cadence).
	none, err := refresh.Union(nil, def)
	require.NoError(t, err)
	assert.Nil(t, none)

	// Empty specs use the default cadence.
	byDefault, err := refresh.Union([]refresh.Spec{""}, def)
	require.NoError(t, err)
	assert.Equal(t, time.Second, byDefault(mon1005))

	// The soonest across a fast and a slow explicit cadence wins.
	mixed, err := refresh.Union([]refresh.Spec{"1m", "5s"}, def)
	require.NoError(t, err)
	assert.Equal(t, 5*time.Second, mixed(mon1005))

	// A default (1s) beats a slower explicit cadence.
	withDefault, err := refresh.Union([]refresh.Spec{"1m", ""}, def)
	require.NoError(t, err)
	assert.Equal(t, time.Second, withDefault(mon1005))

	// A disabled explicit spec ("0s") contributes no cadence; the default remains.
	withDisabled, err := refresh.Union([]refresh.Spec{"0s", ""}, def)
	require.NoError(t, err)
	assert.Equal(t, time.Second, withDisabled(mon1005))

	// Every entry disabled → nil.
	allOff, err := refresh.Union([]refresh.Spec{"0s"}, nil)
	require.NoError(t, err)
	assert.Nil(t, allOff)

	// A malformed pattern is rejected.
	_, err = refresh.Union([]refresh.Spec{"garbage!!!"}, def)
	require.Error(t, err)
}

func TestUnion_EdgeBranches(t *testing.T) {
	t.Parallel()

	def := refresh.Every(time.Second)

	// A slower cadence after a faster one keeps the faster (earlier "keep best").
	ordered, err := refresh.Union([]refresh.Spec{"5s", "1m"}, def)
	require.NoError(t, err)
	assert.Equal(t, 5*time.Second, ordered(mon1005))

	// A pattern with no future instant yields 0 (stop), not a bogus wait.
	past, err := refresh.Union([]refresh.Spec{"12 <2020"}, nil)
	require.NoError(t, err)
	require.NotNil(t, past)
	assert.Zero(t, past(mon1005))
}
