package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var count int
    for i := 0; i < len(birdsPerDay); i++ {
        count += birdsPerDay[i]
    }
    return count
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	if len(birdsPerDay) / 7 < week {
        return 0
    }

    count := 0
    start := week * 7 - 7
    birdsPerDay = birdsPerDay[start:start + 7]
    
    for i := 0; i < len(birdsPerDay); i++ {
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
