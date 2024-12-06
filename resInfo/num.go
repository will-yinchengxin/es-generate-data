package resInfo

import "fmt"

type NumStatsMap map[string]*FieldStats

type FieldStats struct {
	Count int
	Sum   float64
	Max   float64
	Min   float64
	Avg   float64
}

func PrintStats(stats NumStatsMap) {
	for field, stat := range stats {
		fmt.Printf("Field: %s\n", field)
		fmt.Printf("  Count: %d\n", stat.Count)
		if stat.Sum != 0 || stat.Max != 0 || stat.Min != 0 || stat.Avg != 0 {
			fmt.Printf("  Sum: %.4f\n", stat.Sum)
			fmt.Printf("  Max: %.4f\n", stat.Max)
			fmt.Printf("  Min: %.4f\n", stat.Min)
			fmt.Printf("  Avg: %.4f\n", stat.Avg)
		}
	}
}

func UpdateNumericStats(key string, value float64, stats NumStatsMap) {
	if stats[key] == nil {
		stats[key] = &FieldStats{}
	}
	stat := stats[key]
	stat.Count++
	stat.Sum += value
	if stat.Max < value || stat.Count == 1 {
		stat.Max = value
	}
	if stat.Min > value || stat.Count == 1 {
		stat.Min = value
	}
	stat.Avg = stat.Sum / float64(stat.Count)
}
