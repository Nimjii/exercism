package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("hamming: strings are not of same length")
    }

    charsA := []rune(a)
    charsB := []rune(b)
    distance := 0
    
	for i := 0; i < len(charsA); i++ {
        if charsA[i] != charsB[i] {
            distance += 1
        }
    }

    return distance, nil
}
