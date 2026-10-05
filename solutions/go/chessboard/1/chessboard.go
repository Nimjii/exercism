package chessboard

type (
    File []bool
    Chessboard map[string]File
)

func CountInFile(cb Chessboard, file string) int {
    var count int
	for _, occupied := range cb[file] {
		if occupied {
            count++
        }
    }
    return count
}

func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
	var count int
    for _, file := range(cb) {
        if file[rank - 1] {
            count++
        }
    }
    return count
}

func CountAll(cb Chessboard) int {
    var count int
    for _, file := range(cb) {
		count += len(file)
    }
    return count
}

func CountOccupied(cb Chessboard) int {
	var count int
    for _, file := range(cb) {
        for _, occupied := range(file) {
            if occupied {
                count++
            }
        }
    }
    return count
}
