package main

import (
	es2 "es_generate_data/es"
	records2 "es_generate_data/records"
	"es_generate_data/resInfo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	esAddress    string
	indexBase    string
	nrGenerate   int
	templateFile string
)

func init() {
	rootCmd.Flags().StringVarP(&esAddress, "es", "e", "http://172.16.27.45:9200", "Elasticsearch address")
	rootCmd.Flags().StringVarP(&indexBase, "index", "i", "will", "Base name for Elasticsearch index")
	rootCmd.Flags().IntVarP(&nrGenerate, "count", "c", 100, "Number of records to insert into Elasticsearch")
	rootCmd.Flags().StringVarP(&templateFile, "template", "t", "demo.json", "Path to JSON template file")

	//rootCmd.MarkFlagRequired("es")
	//rootCmd.MarkFlagRequired("index")
	//rootCmd.MarkFlagRequired("count")
}

var rootCmd = &cobra.Command{
	Use:   "Weg",
	Short: "ES Generate Tool",
	Long: `The ES Generate Tool Made By Will
 {{host}}: eg: 172.16.27.xxx
 {{method}}: eg: GET/POST/PUT/DELETE
 {{port}}: eg: 9740
 {{timestamp}}: eg: 2024-12-06T02:50:06.701499566+08:00
 {{path}}: eg: /example/13gk1x/123
 {{httpStatus}}: eg: 404
 {{domain}}: eg: will.edu.com
 {{uid}}: eg: 333918a4-c7cc-49fe-843a-4dbacac959f1
 {{randStr}}: eg: k1jd232fjs2akl12djh1
 {{time_local}}: eg: 1733473214
`,
	Run: func(cmd *cobra.Command, args []string) {
		records := records2.GetRecord(templateFile, nrGenerate)
		index, err := es2.InsertToES(esAddress, indexBase, records)
		if err != nil {
			fmt.Printf("Error inserting data: %v\n", err)
		} else {
			fmt.Printf("Successfully inserted %d records into Elasticsearch index '%s'.\n", nrGenerate, index)
		}

		resInfo.PrintStats(records2.StatsMap)
		resInfo.PrintTagStats(records2.TagStatsMap)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
