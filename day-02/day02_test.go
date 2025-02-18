package day01

import (
	"reflect"
	"testing"
	"fmt"
)
// Stock Buy and Sell Problem
func TestStockBuySell(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{100, 180, 260, 310, 40, 535, 695}, want: 655},
		{input: []int{4, 2, 2, 2, 4}, want: 2},
		{input: []int{4, 2}, want: 0},
	}

	for _, tc := range tests {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		res := stockSellBuy(arr)
		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("stockSellBuy(%v) = %v; want %v", tc.input, res, tc.want)
		}
	}
}
