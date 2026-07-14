package session_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/session"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// sample is a small worksheet: 3 data rows, a body formula in column E.
var (
	sampleData     = sheet.Grid{{"1", "2", "3", "4"}, {"2", "3", "4", "5"}, {"3", "4", "5", "6"}}
	sampleTemplate = []byte("=body\nE=C + D\n")
)

func newSession(t *testing.T) *session.Session {
	t.Helper()
	s, err := session.New(sampleTemplate, sampleData)
	require.NoError(t, err)
	return s
}

func TestNew_ComputesEagerly(t *testing.T) {
	t.Parallel()

	state := newSession(t).Snapshot()
	assert.Equal(t, "7", state.Computed[0][4]) // C+D at row 0 = 3+4
	assert.Equal(t, "=body\nE=C + D\n", state.Template)
	assert.False(t, state.Dirty)
	assert.Empty(t, state.Diagnostics)
}

func TestNew_SyntaxError(t *testing.T) {
	t.Parallel()

	_, err := session.New([]byte("=sum("), sampleData)
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}

func TestNew_RejectedTemplate(t *testing.T) {
	t.Parallel()

	_, err := session.New([]byte("=final\n=A:C<"), sampleData)
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrUnsupported)
}

func TestSetTemplate_RecomputesAndDirties(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	require.NoError(t, s.SetTemplate([]byte("=body\nE=C * D\n")))
	state := s.Snapshot()
	assert.Equal(t, "12", state.Computed[0][4]) // C*D at row 0 = 3*4
	assert.True(t, state.Dirty)
}

func TestSetTemplate_AtomicOnSyntaxError(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	before := s.Snapshot()

	err := s.SetTemplate([]byte("=sum("))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)

	after := s.Snapshot()
	assert.Equal(t, before.Computed, after.Computed)
	assert.Equal(t, before.Template, after.Template)
	assert.False(t, after.Dirty) // unchanged: still clean
}

func TestSetTemplate_AtomicOnRejectedTemplate(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	before := s.Snapshot()

	err := s.SetTemplate([]byte("=final\n=A:C<"))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrUnsupported)
	assert.Equal(t, before.Template, s.Snapshot().Template)
}

func TestSetDataCell_EditsAndRecomputes(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	require.NoError(t, s.SetDataCell(sheet.Address{Row: 0, Col: 2}, "10")) // C row 0 = 10
	state := s.Snapshot()
	assert.Equal(t, "10", state.Data[0][2])
	assert.Equal(t, "14", state.Computed[0][4]) // C+D = 10+4
	assert.True(t, state.Dirty)
}

func TestSetDataCell_RejectsNegativeAddress(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	err := s.SetDataCell(sheet.Address{Row: -1, Col: 0}, "x")
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrInvalidValue)
	assert.False(t, s.Snapshot().Dirty) // rejected before any mutation
}

func TestSetDataCell_GrowsGridOnAppend(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	require.NoError(t, s.SetDataCell(sheet.Address{Row: 3, Col: 0}, "new")) // one past last row
	state := s.Snapshot()
	require.Len(t, state.Data, 4)
	assert.Equal(t, "new", state.Data[3][0])
}

func TestDirtyLifecycle(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	assert.False(t, s.Snapshot().Dirty)

	require.NoError(t, s.SetDataCell(sheet.Address{Row: 0, Col: 0}, "9"))
	assert.True(t, s.Snapshot().Dirty)

	s.MarkSaved()
	assert.False(t, s.Snapshot().Dirty)
}

func TestTemplateText(t *testing.T) {
	t.Parallel()

	assert.Equal(t, sampleTemplate, newSession(t).TemplateText())
}

func TestDataTSV(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1\t2\t3\t4\n2\t3\t4\t5\n3\t4\t5\t6\n", string(newSession(t).DataTSV()))
}

func TestSnapshot_IsIsolatedCopy(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	state := s.Snapshot()
	state.Computed[0][0] = "mutated"                  // mutate the snapshot
	assert.Equal(t, "1", s.Snapshot().Computed[0][0]) // session unaffected
}

func TestExplain(t *testing.T) {
	t.Parallel()

	trace, err := newSession(t).Explain(sheet.Address{Row: 0, Col: 4}) // E1 = C+D
	require.NoError(t, err)
	assert.Equal(t, "7", trace.Value)
	assert.Equal(t, "C + D", trace.Formula)
}

func TestExplain_OutOfGrid(t *testing.T) {
	t.Parallel()

	_, err := newSession(t).Explain(sheet.Address{Row: 99, Col: 0})
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrNotFound)
}

func TestConcurrentAccess(t *testing.T) {
	t.Parallel()

	s := newSession(t)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			_ = s.SetDataCell(sheet.Address{Row: 0, Col: 0}, "x")
			_ = s.Snapshot()
			_ = s.TemplateText()
			s.MarkSaved()
		}(i)
	}
	wg.Wait()
}
