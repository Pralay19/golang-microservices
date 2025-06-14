package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//Every test case these statges-

	//Initialization
	ele := []int{9, 8, 7, 6, 5}

	//Execution
	ele = BubbleSort(ele)

	//Validation
	assert.NotNil(t, ele)
	assert.EqualValues(t, 5, len(ele))
	assert.EqualValues(t, 5, ele[0])
	assert.EqualValues(t, 6, ele[1])
	assert.EqualValues(t, 7, ele[2])
	assert.EqualValues(t, 8, ele[3])
	assert.EqualValues(t, 9, ele[4])
}
func TestBubbleSortBestCase(t *testing.T) {
	//Every test case these statges-

	//Initialization
	ele := []int{5, 6, 7, 8, 9}

	//Execution
	ele = BubbleSort(ele)

	//Validation
	assert.NotNil(t, ele)
	assert.EqualValues(t, 5, len(ele))
	assert.EqualValues(t, 5, ele[0])
	assert.EqualValues(t, 6, ele[1])
	assert.EqualValues(t, 7, ele[2])
	assert.EqualValues(t, 8, ele[3])
	assert.EqualValues(t, 9, ele[4])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	ele := getElements(5)
	assert.NotNil(t, ele)
	assert.EqualValues(t, 5, len(ele))
	assert.EqualValues(t, 4, ele[0])
	assert.EqualValues(t, 3, ele[1])
	assert.EqualValues(t, 2, ele[2])
	assert.EqualValues(t, 1, ele[3])
	assert.EqualValues(t, 0, ele[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	ele := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}
func BenchmarkBubbleSort1000(b *testing.B) {
	ele := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}
