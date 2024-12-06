package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"time"
)

// GenerateRandomPaths 根据路径模板生成随机路径
func GenerateRandomPaths(template string, count int) []string {
	var paths []string
	re := regexp.MustCompile(`\{([^{}]+)\}`)

	for i := 0; i < count; i++ {
		path := template
		matches := re.FindAllStringSubmatch(template, -1)
		for _, match := range matches {
			randomValue := fmt.Sprintf("%d", RandomInt(100, 999))
			path = regexp.MustCompile(regexp.QuoteMeta(match[0])).ReplaceAllString(path, randomValue)
		}
		paths = append(paths, path)
	}

	return paths
}

// GenerateIPs 生成指定范围内的 IP 列表
func GenerateIPs(prefix string, count int) []string {
	var ips []string
	for i := 1; i <= count; i++ {
		ips = append(ips, fmt.Sprintf("%s%d", prefix, i))
	}
	return ips
}

// RandomFromSlice 从切片中随机取一个值
func RandomFromSlice(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func RandomFromIntSlice(i []int) int {
	return i[rand.Intn(len(i))]
}

// RandomPort 生成随机端口号
func RandomPort() int {
	return rand.Intn(65535-1024) + 1024
}

// RandomInt 生成指定范围内的随机整数
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomFloat 生成指定范围内的随机浮点数
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomTimestamp 生成随机时间戳
func RandomTimestamp(start, end time.Time) (int64, string) {
	diff := end.Sub(start)
	randomDuration := time.Duration(rand.Int63n(int64(diff)))
	s := start.Add(randomDuration)

	return s.Unix(), s.Format(time.RFC3339Nano)
}

// GenerateInstanceID 生成随机的 instance_id
func GenerateInstanceID() string {
	return uuid.New().String()
}

// GenerateDomain 生成随机的 domain
func GenerateDomain(baseDomain string, subdomains []string) []string {
	d := make([]string, 0)
	for _, val := range subdomains {
		d = append(d, fmt.Sprintf("%s.%s.com", val, baseDomain))
	}

	return d
}

func GenerateRandomStrings(n int, length int) []string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, RandomString(length, charset))
	}

	return result
}

// RandomString 生成指定长度的随机字符串
func RandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func GetRandomInt(num int) int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	return r.Intn(num) + 1
}

func GetRandFloat() float64 {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	return r.Float64()
}
