package slicetools

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	numStrings := []string{
		"1",
		"2",
		"3",
	}
	result := Map(numStrings, func(numStr string) int {
		i, err := strconv.Atoi(numStr)
		require.NoError(t, err)
		return i
	})
	require.Equal(t, []int{1, 2, 3}, result)
}
