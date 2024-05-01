package main

import "testing"

var testInput1 = `1abc2
fdas2p3u5i2a
22fdas
f8dda3a2
ffda1fdas
kl0y1o
qw0x
rerreqa
fdd3af3a`

var expected1 = 183

var testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eighthree`

var expected2 = 364

func TestPart1(t *testing.T) {
	result := part1(testInput1)

	if result != expected1 {
		t.Errorf("part1() = %d; want %d", result, expected1)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput2)

	if result != expected2 {
		t.Errorf("part2() = %d; want %d", result, expected2)
	}
}
