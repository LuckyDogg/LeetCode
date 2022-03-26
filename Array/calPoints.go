package Array

import (
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/baseball-game/

func calPoints(ops []string) int {
	var pointList []int
	for _, v := range ops {
		var point int
		if strings.EqualFold(v, "D") {
			point = pointList[len(pointList)-1] * 2
			pointList = append(pointList, point)
			continue
		} else if strings.EqualFold(v, "C") {
			pointList = pointList[:len(pointList)-1]
			continue
		} else if strings.EqualFold(v, "+") {
			pointList = append(pointList, pointList[len(pointList)-1]+pointList[len(pointList)-2])
		} else {
			// string convert int
			point, _ = strconv.Atoi(v)
			pointList = append(pointList, point)
		}
	}
	var sum int
	for _, v := range pointList {
		sum += v
	}
	return sum
}
