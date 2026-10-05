package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var count int
    for i := 0; i < len(birdsPerDay); i++ {
        count += birdsPerDay[i]
    }
    return count
}

func BirdsInWeek(birdsPerDay []int, week int) int {
    count := 0
    for i := week * 7 - 7; i < week * 7; i++ {
        count += birdsPerDay[i]
    }
    return count
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
