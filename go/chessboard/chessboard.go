package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    occupiedCount := 0

    fileSquares, exists := cb[file]
    if !exists {
        return 0
    }

    for _, occupied := range fileSquares {
        if occupied {
            occupiedCount++
        }
    }
    return occupiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }

    occupiedCount := 0

    for _, file := range cb {
        index := rank - 1
        if file[index] {
            occupiedCount++
        }
    }
    return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    totalCount := 0

    for _, file := range cb {
        totalCount += len(file)
    }
    return totalCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    occupiedCount := 0

    for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                occupiedCount++
            }
        }
    }
    return occupiedCount
}
