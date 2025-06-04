package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	ele := []int{9, 8, 7, 6, 5}

	ele = BubbleSort(ele)

	assert.NotNil(t, ele)
	assert.EqualValues(t, 5, len(ele))
	assert.EqualValues(t, 5, ele[0])
	assert.EqualValues(t, 6, ele[1])
	assert.EqualValues(t, 7, ele[2])
	assert.EqualValues(t, 8, ele[3])
	assert.EqualValues(t, 9, ele[4])
}
