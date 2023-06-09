package stats_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/montanaflynn/stats"
)

func ExampleLoadRawData() {
	data := stats.LoadRawData([]interface{}{1.1, "2", 3})
	fmt.Println(data)
	// Output: [1.1 2 3]
}

var allTestData = []struct {
	actual   interface{}
	expected stats.Float64Data
}{
	{
		[]interface{}{1.0, "2", 3.0, uint(4), "4.0", 5, time.Duration(6), time.Duration(-7)},
		stats.Float64Data{1.0, 2.0, 3.0, 4.0, 4.0, 5.0, 6.0, -7.0},
	},
	{
		[]interface{}{"-345", "223", "-654.4", "194", "898.3"},
		stats.Float64Data{-345.0, 223.0, -654.4, 194.0, 898.3},
	},
	{
		[]interface{}{7862, 4234, 9872.1, 8794},
		stats.Float64Data{7862.0, 4234.0, 9872.1, 8794.0},
	},
	{
		[]interface{}{true, false, true, false, false},
		stats.Float64Data{1.0, 0.0, 1.0, 0.0, 0.0},
	},
	{
		[]interface{}{14.3, 26, 17.7, "shoe"},
		stats.Float64Data{14.3, 26.0, 17.7},
	},
	{
		[]bool{true, false, true, true, false},
		stats.Float64Data{1.0, 0.0, 1.0, 1.0, 0.0},
	},
	{
		[]float64{10230.9823, 93432.9384, 23443.945, 12374.945},
		stats.Float64Data{10230.9823, 93432.9384, 23443.945, 12374.945},
	},
	{
		[]time.Duration{-843, 923, -398, 1000},
		stats.Float64Data{-843.0, 923.0, -398.0, 1000.0},
	},
	{
		[]string{"-843.2", "923", "hello", "-398", "1000.5"},
		stats.Float64Data{-843.2, 923.0, -398.0, 1000.5},
	},
	{
		[]uint{34, 12, 65, 230, 30},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 30.0},
	},
	{
		[]uint8{34, 12, 65, 23, 255},
		stats.Float64Data{34.0, 12.0, 65.0, 23.0, 255.0},
	},
	{
		[]uint16{34, 12, 65, 230, 65535},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 65535.0},
	},
	{
		[]uint32{34, 12, 65, 230, 4294967295},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 4294967295.0},
	},
	{
		[]uint64{34, 12, 65, 230, 18446744073709551615},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 18446744073709552000.0},
	},
	{
		[]int{-843, 923, -398, 1000},
		stats.Float64Data{-843.0, 923.0, -398.0, 1000.0},
	},
	{
		[]int8{-43, 23, -128, 127},
		stats.Float64Data{-43.0, 23.0, -128.0, 127.0},
	},
	{
		[]int16{-843, 923, -32768, 32767},
		stats.Float64Data{-843.0, 923.0, -32768.0, 32767.0},
	},
	{
		[]int32{-843, 923, -2147483648, 2147483647},
		stats.Float64Data{-843.0, 923.0, -2147483648.0, 2147483647.0},
	},
	{
		[]int64{-843, 923, -9223372036854775808, 9223372036854775807, 9223372036854775800},
		stats.Float64Data{-843.0, 923.0, -9223372036854776000.0, 9223372036854776000.0, 9223372036854776000.0},
	},
	{
		map[int]bool{0: true, 1: true, 2: false, 3: true, 4: false},
		stats.Float64Data{1.0, 1.0, 0.0, 1.0, 0.0},
	},
	{
		map[int]float64{0: 68.6, 1: 72.1, 2: -33.3, 3: -99.2},
		stats.Float64Data{68.6, 72.1, -33.3, -99.2},
	},
	{
		map[int]time.Duration{0: -843, 1: 923, 2: -398, 3: 1000},
		stats.Float64Data{-843.0, 923.0, -398.0, 1000.0},
	},
	{
		map[int]string{0: "456", 1: "758", 2: "-9874", 3: "-1981", 4: "68.6", 5: "72.1", 6: "-33.3", 7: "-99.2"},
		stats.Float64Data{456.0, 758.0, -9874.0, -1981.0, 68.6, 72.1, -33.3, -99.2},
	},
	{
		map[int]uint{0: 4567, 1: 7580, 2: 98742, 3: 19817},
		stats.Float64Data{4567.0, 7580.0, 98742.0, 19817.0},
	},
	{
		map[int]uint8{0: 34, 1: 12, 2: 65, 3: 23, 4: 255},
		stats.Float64Data{34.0, 12.0, 65.0, 23.0, 255.0},
	},
	{
		map[int]uint16{0: 34, 1: 12, 2: 65, 3: 230, 4: 65535},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 65535.0},
	},
	{
		map[int]uint32{0: 34, 1: 12, 2: 65, 3: 230, 4: 4294967295},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 4294967295.0},
	},
	{
		map[int]uint64{0: 34, 1: 12, 2: 65, 3: 230, 4: 18446744073709551615},
		stats.Float64Data{34.0, 12.0, 65.0, 230.0, 18446744073709552000.0},
	},
	{
		map[int]int{0: 456, 1: 758, 2: -9874, 3: -1981},
		stats.Float64Data{456.0, 758.0, -9874.0, -1981.0},
	},
	{
		map[int]int8{0: -43, 1: 23, 2: -128, 3: 127},
		stats.Float64Data{-43.0, 23.0, -128.0, 127.0},
	},
	{
		map[int]int16{0: -843, 1: 923, 2: -32768, 3: 32767},
		stats.Float64Data{-843.0, 923.0, -32768.0, 32767.0},
	},
	{
		map[int]int32{0: -843, 1: 923, 2: -2147483648, 3: 2147483647},
		stats.Float64Data{-843.0, 923.0, -2147483648.0, 2147483647.0},
	},
	{
		map[int]int64{0: -843, 1: 923, 2: -9223372036854775808, 3: 9223372036854775807, 4: 9223372036854775800},
		stats.Float64Data{-843.0, 923.0, -9223372036854776000.0, 9223372036854776000.0, 9223372036854776000.0},
	},
	{
		"1\n\n2 3.3\n  4.4",
		stats.Float64Data{1.0, 2, 3.3, 4.4},
	},
	{
		strings.NewReader("1\n\n2 3.3\n  4.4"),
		stats.Float64Data{1.0, 2, 3.3, 4.4},
	},
}

func equal(actual, expected stats.Float64Data) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, actualVal := range actual {
		if actualVal != expected[k] {
			return false
		}
	}

	return true
}

func TestLoadRawData(t *testing.T) {
	for _, data := range allTestData {
		actual := stats.LoadRawData(data.actual)
		if !equal(actual, data.expected) {
			t.Fatalf("Transform(%v). Expected [%v], Actual [%v]", data.actual, data.expected, actual)
		}
	}
}
