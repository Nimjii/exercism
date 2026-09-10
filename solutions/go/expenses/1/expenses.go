package expenses

import "fmt"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var records []Record
    for _, record := range(in) {
        if predicate(record) {
            records = append(records, record)
        }
    }
    return records
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
        return record.Day >= p.From && record.Day <= p.To
    }
}

func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
        return record.Category == c
    }
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    var total float64
    for _, record := range(Filter(in, ByDaysPeriod(p))) {
        total += record.Amount
    }
    return total
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var total float64
	categoryRecords := Filter(in, ByCategory(c))
    if len(categoryRecords) == 0 {
        return 0.0, fmt.Errorf("unknown category %s", c)
    }
    for _, record := range(Filter(categoryRecords, ByDaysPeriod(p))) {
        total += record.Amount
    }
    return total, nil
}
