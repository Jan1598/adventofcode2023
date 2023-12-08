package day08

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day08"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")
var lines2 = adventofcode2023.ReadData("input_test2.txt")
var lines3 = adventofcode2023.ReadData("input_test3.txt")

func TestCheckKinds(t *testing.T) {
	result := day08.CreateNetworkMap(lines)

	adventofcode2023.AssertEquals(t, 7, len(result))
}

func TestGetInstructionByInput(t *testing.T) {
	result := day08.GetInstructionByInput(lines)

	adventofcode2023.AssertEquals(t, "RL", result)
}

func TestReachValueInMap(t *testing.T) {
	result := day08.ReachValueInMap(day08.CreateNetworkMap(lines), day08.GetInstructionByInput(lines))

	adventofcode2023.AssertEquals(t, 2, result)
}

func TestReachValueInMap2(t *testing.T) {
	result := day08.ReachValueInMap(day08.CreateNetworkMap(lines2), day08.GetInstructionByInput(lines2))

	adventofcode2023.AssertEquals(t, 6, result)
}

func TestCollectAllStartPoints(t *testing.T) {
	result := day08.CollectAllStartPoints(day08.CreateNetworkMap(lines3))

	adventofcode2023.AssertEquals(t, 2, len(result))
}

func TestReachValueInMapSimultaneously(t *testing.T) {
	result := day08.ReachValueInMapSimultaneously(day08.CreateNetworkMap(lines3), day08.GetInstructionByInput(lines3), 10)

	adventofcode2023.AssertEquals(t, 6, result)
}
