package sheet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsVolatile(t *testing.T) {
	t.Parallel()

	assert.True(t, parse(t, "=now()\n").IsVolatile())
	assert.True(t, parse(t, "=today()\n").IsVolatile())
	assert.True(t, parse(t, "=isnow(\"noon\")\n").IsVolatile())
	assert.False(t, parse(t, "=sum(A1:A2)\n").IsVolatile()) // a call, but not clock-dependent
	assert.False(t, parse(t, "5\t=A1 + 1\n").IsVolatile())  // literals and a call-free formula
}
