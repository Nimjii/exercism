package isbnverifier

import (
    "strings"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")

    if len(isbn) != 10 {
        return false
    }

    var result int

    for i, j := 0, 10; i < len(isbn); i, j = i+1, j-1 {
        if isbn[i] == 'X' {
            if i != len(isbn) - 1 {
                return false
            }
            
            result += 10 * j
        } else {
         	digit := int(isbn[i] - '0')
            result += digit * j
        }
    }
    
    return result % 11 == 0
}
