package luhn

func Valid(id string) bool {
    var result int
    var pos int

	for i := len(id) - 1; i >= 0; i-- {        
        if id[i] == ' ' {
            continue
        }
        
        if (id[i] < '0' || id[i] > '9') {
            return false
        }

        pos++
        digit := int(id[i] - '0')
        
        if pos % 2 == 0 {
            result += (digit * 2 - 1) % 9 + 1
        } else {
            result += digit
        }
    }

    if pos <= 1 {
        return false
    }

    return result % 10 == 0
}
