package resInfo

import "fmt"

type TagStatsMap map[string]map[string]*TagStats

type TagStats struct {
	Sum float64
}

func UpdateTagStats(key, subKey string, stats TagStatsMap) {
	if stats[key] == nil {
		stats[key] = make(map[string]*TagStats)
	}
	if stats[key][subKey] == nil {
		stats[key][subKey] = &TagStats{}
	}
	stats[key][subKey].Sum += 1
}

func PrintTagStats(stats TagStatsMap) {
	for field, stat := range stats {
		fmt.Printf("Tag: %s\n", field)
		for key, val := range stat {
			fmt.Printf("  Field: %s, Sum: %v\n", key, val.Sum)
		}
	}
}
