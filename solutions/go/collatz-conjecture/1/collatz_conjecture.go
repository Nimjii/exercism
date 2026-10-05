package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, fmt.Errorf("Input has to be 1 or higher (%d given)", n)
    }

    var steps int
    
    for n > 1 {
        steps += 1
        
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
    }

    return steps, nil
}
