package tokenize

import (
	"github.com/stretchr/testify/require"
	"iter"
	"testing"
)

func TestTokenizer(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		w := String("")
		require.Empty(t, gather(w), "must have been empty")
	})

	t.Run("single token", func(t *testing.T) {
		w := String("hello")
		require.Equal(t, []string{"hello"}, gather(w))
	})

	t.Run("excessively spaces polluted", func(t *testing.T) {
		w := String(" hello    world   ")
		require.Equal(t, []string{"hello", "world"}, gather(w))
	})
}

func gather(iterator iter.Seq[string]) (out []string) {
	for str := range iterator {
		out = append(out, str)
	}

	return out
}
