package main

func checkTwoChessboards(coordinate1, coordinate2 string) bool {
	return (coordinate1[0]^coordinate1[1])&1 == (coordinate2[0]^coordinate2[1])&1
}
