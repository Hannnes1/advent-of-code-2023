package main

import "testing"

var testInput = `1abc2
fdas2p3u5i2a
22fdas
f8dda3a2
ffda1fdas
kl0y1o
qw0x
rerreqa
fdd3af3a`

var expected = 183

func TestPart1(t *testing.T) {
  result := part1(testInput)

  if result != expected {
    t.Errorf("part1() = %d; want %d", result, expected)
  }
}
