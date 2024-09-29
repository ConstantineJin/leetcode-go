package main

import "math/bits"

func kthCharacter(k int) byte { return byte('a' + bits.OnesCount(uint(k-1))) }
