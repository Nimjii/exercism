package sorting

import (
    "fmt"
    "strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	nFancy, _ := fnb.(FancyNumber)
    n, _ := strconv.Atoi(nFancy.Value())
    return n
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	nFancy, _ := fnb.(FancyNumber)
    float, _ := strconv.ParseFloat(nFancy.Value(), 64)
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float)
}

func DescribeAnything(i any) string {
	switch v := i.(type) {
    case int:
        return DescribeNumber(float64(v))
    case float64:
        return DescribeNumber(v)
    case NumberBox:
        return DescribeNumberBox(v)
    case FancyNumberBox:
        return DescribeFancyNumberBox(v)
    }
    return "Return to sender"
}
