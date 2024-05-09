package chessboard

// File stores if a square is occupied by a piece
type File []bool

// Chessboard Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	total := 0
	for _, present := range cb[file] {
		if present {
			total++
		}
	}
	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	total := 0

	if rank > 8 || rank < 1 {
		return 0
	}

	for _, file := range cb {
		if file[rank-1] {
			total++
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total := 0
	for _, file := range cb {
		for _, _ = range file {
			total++
		}
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.m
func CountOccupied(cb Chessboard) int {
	total := 0
	for _, file := range cb {
		for _, present := range file {
			if present {
				total++
			}
		}
	}
	return total
}
