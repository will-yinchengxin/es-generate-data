## ES-Generate-Data
根据模版样式, 快速向 ES 中批量插入数据

编译:
````
go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o WEG main.go
````	
	

使用方式:
```
WEG -e <your_es_addr> -i <your_es_indice_name> -c <the_record_number> -t <the_josn_template_path>
```
