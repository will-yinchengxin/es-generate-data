package records

import (
	"encoding/json"
	"es_generate_data/resInfo"
	"es_generate_data/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	count       = 20
	methods     = []string{"GET", "POST", "PUT", "DELETE"}
	hosts       = utils.GenerateIPs("172.16.27.", count)
	domains     = utils.GenerateDomain("will", utils.GenerateRandomStrings(count, 5))
	randomPaths = utils.GenerateRandomPaths("/example/{path}/{resource}", count) // 限定 10 个随机路径
	statusCodes = []string{"200", "302", "404", "403", "500", "507"}
	randStr     = utils.GenerateRandomStrings(count, 10)
)

var (
	StatsMap    = resInfo.NumStatsMap{}
	TagStatsMap = resInfo.TagStatsMap{}
)

func GetRecord(templateFile string, nrGenerate int) []map[string]interface{} {
	templateData, err := ReadJSONTemplate(templateFile)
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return nil
	}

	return generateDynamicData(templateData, nrGenerate)
}

// ReadJSONTemplate 从文件读取 JSON 模板
func ReadJSONTemplate(filename string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// generateDynamicData 根据模板生成随机数据
func generateDynamicData(template map[string]interface{}, count int) []map[string]interface{} {
	var (
		records    []map[string]interface{}
		now        = time.Now()
		startOfDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	)

	for i := 0; i < count; i++ {
		record := make(map[string]interface{})
		for key, value := range template {
			switch v := value.(type) {
			case string:
				if strings.Contains(v, "{{") {
					record[key] = generateRandomValue(key, v, startOfDay, now)
				} else {
					record[key] = value
				}
			case float64:
				var (
					nr float64
				)
				intPart := int64(v)
				if v != float64(intPart) {
					nr = utils.GetRandFloat() + float64(utils.GetRandomInt(10))
					record[key] = nr

				} else {
					nr = float64(utils.GetRandomInt(100))
					record[key] = nr
				}
				resInfo.UpdateNumericStats(key, nr, StatsMap)
			default:
				record[key] = value
			}
		}
		records = append(records, record)
	}
	return records
}

// generateRandomValue 根据占位符生成随机值
func generateRandomValue(key, placeholder string, start, end time.Time) interface{} {
	timInt, timStr := utils.RandomTimestamp(start, end)
	switch placeholder {
	case "{{host}}":
		t := utils.RandomFromSlice(hosts)
		resInfo.UpdateTagStats(key, t, TagStatsMap)
		return t
	case "{{method}}":
		t := utils.RandomFromSlice(methods)
		resInfo.UpdateTagStats(key, t, TagStatsMap)
		return t
	case "{{port}}":
		t := utils.RandomPort()
		resInfo.UpdateTagStats(key, fmt.Sprintf("%d", t), TagStatsMap)
		return t
	case "{{path}}":
		t := utils.RandomFromSlice(randomPaths)
		resInfo.UpdateTagStats(key+"-{{path}}", t, TagStatsMap)
		return t
	case "{{httpStatus}}":
		t := utils.RandomFromSlice(statusCodes)
		resInfo.UpdateTagStats(key, t, TagStatsMap)
		return t
	case "{{domain}}":
		t := utils.RandomFromSlice(domains)
		resInfo.UpdateTagStats(key, t, TagStatsMap)
		return t
	case "{{randStr}}":
		t := utils.RandomFromSlice(randStr)
		resInfo.UpdateTagStats(key, t, TagStatsMap)
		return t
	case "{{uid}}":
		return utils.GenerateInstanceID()
	case "{{time_local}}":
		return timInt
	case "{{timestamp}}":
		return timStr
	default:
		return "-"
	}
}
