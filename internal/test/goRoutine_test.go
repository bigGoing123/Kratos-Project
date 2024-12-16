package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestGoRoutine(t *testing.T) {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	n := len(points)
	ret := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	flag := make([]int, n)
	for i := 0; i < n; i++ {
		flag[i] = 0 //标记数组，是否已经加入某一个区间
	}
	for i := 0; i < n; i++ {
		if flag[i] == 1 {
			continue
		}
		raw := points[i]
		flag[i] = 1
		for j := i + 1; j < n && flag[j] == 0; j++ {
			room := points[j]
			if raw[1] < room[0] {
				break
			}
			if room[0] >= raw[0] && room[0] <= raw[1] {
				flag[j] = 1
			}
		}
		ret++
	}
	fmt.Println(ret)
}

func Test123(t *testing.T) {

}
