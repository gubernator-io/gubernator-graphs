package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func main() {
	file := bytes.Buffer{}
	file.WriteString(`
goos: darwin
goarch: arm64
pkg: github.com/gubernator-io/gubernator/v2/bench
BenchmarkParallel
Current Operating System has '10' CPUs
BenchmarkParallel/NoCache_1
BenchmarkParallel/NoCache_1-10         	 4817940	       238.2 ns/op	   4198712 ops/s
testing: BenchmarkParallel/NoCache_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/NoCache_2
BenchmarkParallel/NoCache_2-10         	 8085074	       150.5 ns/op	   6646397 ops/s
testing: BenchmarkParallel/NoCache_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/NoCache_4
BenchmarkParallel/NoCache_4-10         	13023541	        90.31 ns/op	  11072893 ops/s
testing: BenchmarkParallel/NoCache_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/NoCache_8
BenchmarkParallel/NoCache_8-10         	10760613	       112.1 ns/op	   8923662 ops/s
testing: BenchmarkParallel/NoCache_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/NoCache_12
BenchmarkParallel/NoCache_12-10        	 9760540	       121.0 ns/op	   8266777 ops/s
testing: BenchmarkParallel/NoCache_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/NoCache_16
BenchmarkParallel/NoCache_16-10        	 9564832	       123.0 ns/op	   8132989 ops/s
testing: BenchmarkParallel/NoCache_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/NoCache_20
BenchmarkParallel/NoCache_20-10        	 9639396	       126.0 ns/op	   7935060 ops/s
testing: BenchmarkParallel/NoCache_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/NoCache_24
BenchmarkParallel/NoCache_24-10        	 9446574	       126.0 ns/op	   7938074 ops/s
testing: BenchmarkParallel/NoCache_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/NoCache_28
BenchmarkParallel/NoCache_28-10        	 9678115	       129.3 ns/op	   7735196 ops/s
testing: BenchmarkParallel/NoCache_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/NoCache_32
BenchmarkParallel/NoCache_32-10        	 9232617	       124.2 ns/op	   8053661 ops/s
testing: BenchmarkParallel/NoCache_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/MutexRead_1
BenchmarkParallel/MutexRead_1-10       	 2108930	       563.2 ns/op	   1775389 ops/s
testing: BenchmarkParallel/MutexRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/MutexRead_2
BenchmarkParallel/MutexRead_2-10       	 2139711	       542.8 ns/op	   1842195 ops/s
testing: BenchmarkParallel/MutexRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/MutexRead_4
BenchmarkParallel/MutexRead_4-10       	 2044603	       567.1 ns/op	   1763260 ops/s
testing: BenchmarkParallel/MutexRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/MutexRead_8
BenchmarkParallel/MutexRead_8-10       	 1974379	       605.5 ns/op	   1651569 ops/s
testing: BenchmarkParallel/MutexRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/MutexRead_12
BenchmarkParallel/MutexRead_12-10      	 1911454	       609.7 ns/op	   1640078 ops/s
testing: BenchmarkParallel/MutexRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/MutexRead_16
BenchmarkParallel/MutexRead_16-10      	 1926068	       617.6 ns/op	   1619265 ops/s
testing: BenchmarkParallel/MutexRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/MutexRead_20
BenchmarkParallel/MutexRead_20-10      	 1910180	       620.7 ns/op	   1611067 ops/s
testing: BenchmarkParallel/MutexRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/MutexRead_24
BenchmarkParallel/MutexRead_24-10      	 1896154	       622.8 ns/op	   1605625 ops/s
testing: BenchmarkParallel/MutexRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/MutexRead_28
BenchmarkParallel/MutexRead_28-10      	 1843675	       627.7 ns/op	   1592983 ops/s
testing: BenchmarkParallel/MutexRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/MutexRead_32
BenchmarkParallel/MutexRead_32-10      	 1895296	       630.4 ns/op	   1586326 ops/s
testing: BenchmarkParallel/MutexRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/MutexWrite_1
BenchmarkParallel/MutexWrite_1-10      	 2159419	       561.4 ns/op	   1781144 ops/s
testing: BenchmarkParallel/MutexWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/MutexWrite_2
BenchmarkParallel/MutexWrite_2-10      	 2164585	       538.8 ns/op	   1856053 ops/s
testing: BenchmarkParallel/MutexWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/MutexWrite_4
BenchmarkParallel/MutexWrite_4-10      	 2071174	       571.6 ns/op	   1749567 ops/s
testing: BenchmarkParallel/MutexWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/MutexWrite_8
BenchmarkParallel/MutexWrite_8-10      	 1976018	       602.2 ns/op	   1660422 ops/s
testing: BenchmarkParallel/MutexWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/MutexWrite_12
BenchmarkParallel/MutexWrite_12-10     	 1951068	       607.5 ns/op	   1646076 ops/s
testing: BenchmarkParallel/MutexWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/MutexWrite_16
BenchmarkParallel/MutexWrite_16-10     	 1903671	       623.6 ns/op	   1603474 ops/s
testing: BenchmarkParallel/MutexWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/MutexWrite_20
BenchmarkParallel/MutexWrite_20-10     	 1900621	       621.0 ns/op	   1610405 ops/s
testing: BenchmarkParallel/MutexWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/MutexWrite_24
BenchmarkParallel/MutexWrite_24-10     	 1875744	       616.6 ns/op	   1621721 ops/s
testing: BenchmarkParallel/MutexWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/MutexWrite_28
BenchmarkParallel/MutexWrite_28-10     	 1899164	       623.2 ns/op	   1604494 ops/s
testing: BenchmarkParallel/MutexWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/MutexWrite_32
BenchmarkParallel/MutexWrite_32-10     	 1886140	       635.5 ns/op	   1573633 ops/s
testing: BenchmarkParallel/MutexWrite_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/OtterRead_1
BenchmarkParallel/OtterRead_1-10       	 1708082	       695.2 ns/op	   1438294 ops/s
testing: BenchmarkParallel/OtterRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/OtterRead_2
BenchmarkParallel/OtterRead_2-10       	 2260218	       522.4 ns/op	   1914099 ops/s
testing: BenchmarkParallel/OtterRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/OtterRead_4
BenchmarkParallel/OtterRead_4-10       	 2447698	       487.5 ns/op	   2051058 ops/s
testing: BenchmarkParallel/OtterRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/OtterRead_8
BenchmarkParallel/OtterRead_8-10       	 1248458	       979.4 ns/op	   1021022 ops/s
testing: BenchmarkParallel/OtterRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/OtterRead_12
BenchmarkParallel/OtterRead_12-10      	 1000000	      1204 ns/op	    830660 ops/s
testing: BenchmarkParallel/OtterRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/OtterRead_16
BenchmarkParallel/OtterRead_16-10      	  921481	      1268 ns/op	    788476 ops/s
testing: BenchmarkParallel/OtterRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/OtterRead_20
BenchmarkParallel/OtterRead_20-10      	 1000000	      1411 ns/op	    708606 ops/s
testing: BenchmarkParallel/OtterRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/OtterRead_24
BenchmarkParallel/OtterRead_24-10      	 1000000	      1383 ns/op	    722901 ops/s
testing: BenchmarkParallel/OtterRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/OtterRead_28
BenchmarkParallel/OtterRead_28-10      	 1000000	      1475 ns/op	    677903 ops/s
testing: BenchmarkParallel/OtterRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/OtterRead_32
BenchmarkParallel/OtterRead_32-10      	  865990	      1508 ns/op	    663294 ops/s
testing: BenchmarkParallel/OtterRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/OtterWrite_1
BenchmarkParallel/OtterWrite_1-10      	 1813698	       747.4 ns/op	   1337915 ops/s
testing: BenchmarkParallel/OtterWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/OtterWrite_2
BenchmarkParallel/OtterWrite_2-10      	 2432829	       528.5 ns/op	   1892215 ops/s
testing: BenchmarkParallel/OtterWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/OtterWrite_4
BenchmarkParallel/OtterWrite_4-10      	 2490639	       483.8 ns/op	   2066775 ops/s
testing: BenchmarkParallel/OtterWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/OtterWrite_8
BenchmarkParallel/OtterWrite_8-10      	 1227604	       978.9 ns/op	   1021519 ops/s
testing: BenchmarkParallel/OtterWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/OtterWrite_12
BenchmarkParallel/OtterWrite_12-10     	 1000000	      1224 ns/op	    817108 ops/s
testing: BenchmarkParallel/OtterWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/OtterWrite_16
BenchmarkParallel/OtterWrite_16-10     	 1000000	      1358 ns/op	    736153 ops/s
testing: BenchmarkParallel/OtterWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/OtterWrite_20
BenchmarkParallel/OtterWrite_20-10     	 1000000	      1342 ns/op	    745296 ops/s
testing: BenchmarkParallel/OtterWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/OtterWrite_24
BenchmarkParallel/OtterWrite_24-10     	 1000000	      1341 ns/op	    745956 ops/s
testing: BenchmarkParallel/OtterWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/OtterWrite_28
BenchmarkParallel/OtterWrite_28-10     	 1000000	      1507 ns/op	    663633 ops/s
testing: BenchmarkParallel/OtterWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/OtterWrite_32
BenchmarkParallel/OtterWrite_32-10     	 1000000	      1682 ns/op	    594502 ops/s
testing: BenchmarkParallel/OtterWrite_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/WorkerPoolRead_1
INFO[0092] Starting 1 Gubernator workers...             
INFO[0092] Starting 1 Gubernator workers...             
INFO[0092] Starting 1 Gubernator workers...             
INFO[0093] Starting 1 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_1-10  	  881487	      1610 ns/op	    620939 ops/s
testing: BenchmarkParallel/WorkerPoolRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/WorkerPoolRead_2
INFO[0094] Starting 2 Gubernator workers...             
INFO[0095] Starting 2 Gubernator workers...             
INFO[0095] Starting 2 Gubernator workers...             
INFO[0095] Starting 2 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_2-10  	  730093	      1736 ns/op	    575964 ops/s
testing: BenchmarkParallel/WorkerPoolRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/WorkerPoolRead_4
INFO[0097] Starting 4 Gubernator workers...             
INFO[0097] Starting 4 Gubernator workers...             
INFO[0097] Starting 4 Gubernator workers...             
INFO[0097] Starting 4 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_4-10  	  825300	      1530 ns/op	    653783 ops/s
testing: BenchmarkParallel/WorkerPoolRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/WorkerPoolRead_8
INFO[0098] Starting 8 Gubernator workers...             
INFO[0099] Starting 8 Gubernator workers...             
INFO[0099] Starting 8 Gubernator workers...             
INFO[0099] Starting 8 Gubernator workers...             
INFO[0100] Starting 8 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_8-10  	  789648	      1559 ns/op	    641613 ops/s
testing: BenchmarkParallel/WorkerPoolRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/WorkerPoolRead_12
INFO[0101] Starting 12 Gubernator workers...            
INFO[0101] Starting 12 Gubernator workers...            
INFO[0101] Starting 12 Gubernator workers...            
INFO[0102] Starting 12 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_12-10 	  798327	      1872 ns/op	    534163 ops/s
testing: BenchmarkParallel/WorkerPoolRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/WorkerPoolRead_16
INFO[0103] Starting 16 Gubernator workers...            
INFO[0103] Starting 16 Gubernator workers...            
INFO[0104] Starting 16 Gubernator workers...            
INFO[0104] Starting 16 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_16-10 	  691045	      1591 ns/op	    628379 ops/s
testing: BenchmarkParallel/WorkerPoolRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/WorkerPoolRead_20
INFO[0105] Starting 20 Gubernator workers...            
INFO[0105] Starting 20 Gubernator workers...            
INFO[0105] Starting 20 Gubernator workers...            
INFO[0106] Starting 20 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_20-10 	  860224	      1798 ns/op	    556264 ops/s
testing: BenchmarkParallel/WorkerPoolRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/WorkerPoolRead_24
INFO[0107] Starting 24 Gubernator workers...            
INFO[0107] Starting 24 Gubernator workers...            
INFO[0108] Starting 24 Gubernator workers...            
INFO[0108] Starting 24 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_24-10 	  756060	      1536 ns/op	    651161 ops/s
testing: BenchmarkParallel/WorkerPoolRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/WorkerPoolRead_28
INFO[0109] Starting 28 Gubernator workers...            
INFO[0109] Starting 28 Gubernator workers...            
INFO[0110] Starting 28 Gubernator workers...            
INFO[0110] Starting 28 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_28-10 	  722101	      2111 ns/op	    473770 ops/s
testing: BenchmarkParallel/WorkerPoolRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/WorkerPoolRead_32
INFO[0111] Starting 32 Gubernator workers...            
INFO[0112] Starting 32 Gubernator workers...            
INFO[0112] Starting 32 Gubernator workers...            
INFO[0112] Starting 32 Gubernator workers...            
INFO[0113] Starting 32 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_32-10 	  825039	      1380 ns/op	    724842 ops/s
testing: BenchmarkParallel/WorkerPoolRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/WorkerPoolWrite_1
INFO[0114] Starting 1 Gubernator workers...             
INFO[0115] Starting 1 Gubernator workers...             
INFO[0115] Starting 1 Gubernator workers...             
INFO[0116] Starting 1 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_1-10 	  812012	      1375 ns/op	    727156 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/WorkerPoolWrite_2
INFO[0117] Starting 2 Gubernator workers...             
INFO[0117] Starting 2 Gubernator workers...             
INFO[0118] Starting 2 Gubernator workers...             
INFO[0118] Starting 2 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_2-10 	  729654	      1625 ns/op	    615555 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/WorkerPoolWrite_4
INFO[0119] Starting 4 Gubernator workers...             
INFO[0120] Starting 4 Gubernator workers...             
INFO[0120] Starting 4 Gubernator workers...             
INFO[0120] Starting 4 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_4-10 	  782449	      1484 ns/op	    673895 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/WorkerPoolWrite_8
INFO[0121] Starting 8 Gubernator workers...             
INFO[0121] Starting 8 Gubernator workers...             
INFO[0121] Starting 8 Gubernator workers...             
INFO[0122] Starting 8 Gubernator workers...             
INFO[0122] Starting 8 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_8-10 	  707136	      1546 ns/op	    646928 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/WorkerPoolWrite_12
INFO[0124] Starting 12 Gubernator workers...            
INFO[0124] Starting 12 Gubernator workers...            
INFO[0124] Starting 12 Gubernator workers...            
INFO[0124] Starting 12 Gubernator workers...            
INFO[0125] Starting 12 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_12-10         	  820294	      1785 ns/op	    560186 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/WorkerPoolWrite_16
INFO[0126] Starting 16 Gubernator workers...            
INFO[0126] Starting 16 Gubernator workers...            
INFO[0126] Starting 16 Gubernator workers...            
INFO[0127] Starting 16 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_16-10         	  276675	      3729 ns/op	    268182 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/WorkerPoolWrite_20
INFO[0128] Starting 20 Gubernator workers...            
INFO[0128] Starting 20 Gubernator workers...            
INFO[0128] Starting 20 Gubernator workers...            
INFO[0128] Starting 20 Gubernator workers...            
INFO[0129] Starting 20 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_20-10         	  662256	      4345 ns/op	    230138 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/WorkerPoolWrite_24
INFO[0132] Starting 24 Gubernator workers...            
INFO[0132] Starting 24 Gubernator workers...            
INFO[0132] Starting 24 Gubernator workers...            
INFO[0132] Starting 24 Gubernator workers...            
INFO[0133] Starting 24 Gubernator workers...            
INFO[0134] Starting 24 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_24-10         	  639648	      2621 ns/op	    381467 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/WorkerPoolWrite_28
INFO[0136] Starting 28 Gubernator workers...            
INFO[0136] Starting 28 Gubernator workers...            
INFO[0136] Starting 28 Gubernator workers...            
INFO[0136] Starting 28 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_28-10         	  266378	      4230 ns/op	    236412 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/WorkerPoolWrite_32
INFO[0138] Starting 32 Gubernator workers...            
INFO[0138] Starting 32 Gubernator workers...            
INFO[0138] Starting 32 Gubernator workers...            
INFO[0138] Starting 32 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_32-10         	  315619	      4937 ns/op	    202537 ops/s
testing: BenchmarkParallel/WorkerPoolWrite_32-10 left GOMAXPROCS set to 32
PASS


	`)

	plot1 := make([]opts.LineData, 0)
	plot2 := make([]opts.LineData, 0)
	plot3 := make([]opts.LineData, 0)
	plot4 := make([]opts.LineData, 0)
	plot5 := make([]opts.LineData, 0)
	plot6 := make([]opts.LineData, 0)
	plot7 := make([]opts.LineData, 0)

	const (
		plot1Name = "MutexRead"
		plot2Name = "MutexWrite"
		plot3Name = "NoCache"
		plot4Name = "WorkerPoolRead"
		plot5Name = "WorkerPoolWrite"
		plot6Name = "OtterRead"
		plot7Name = "OtterWrite"
	)

	// Regular expression to match benchmark lines
	// Pattern to match benchmark that reports only nanoseconds
	// re := regexp.MustCompile(`Benchmark\w+/(\w+)_(\d+)-\d+\s+(\d+)\s+(\d+(\.\d+|)) ns/op`)
	// IE: BenchmarkWorkerPool/Write_100-10      	    373  380416 ns/op

	// Expression to match benchmark output that includes the operations per second
	// IE: BenchmarkParallel/NoCache_1-10         	 4955510       246.4 ns/op	   4058316 ops/s
	re := regexp.MustCompile(`Benchmark\w+/(\w+)_(\d+)-\d+\s+(\d+)\s+(\d+(\.\d+|))\s+ns/op\s*(\d+)\s+ops/s`)

	scanner := bufio.NewScanner(&file)
	var xAxis []string
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Get the ops/s
			yValue, _ := strconv.ParseInt(matches[6], 10, 64)

			fmt.Printf("Match: %s [%s] Y: %d\n", matches[1], matches[6], yValue)
			switch matches[1] {
			case plot1Name:
				fmt.Printf("")
				xAxis = append(xAxis, matches[2])
				plot1 = append(plot1, opts.LineData{Value: yValue})
			case plot2Name:
				plot2 = append(plot2, opts.LineData{Value: yValue})
			case plot3Name:
				plot3 = append(plot3, opts.LineData{Value: yValue})
			case plot4Name:
				plot4 = append(plot4, opts.LineData{Value: yValue})
			case plot5Name:
				plot5 = append(plot5, opts.LineData{Value: yValue})
			case plot6Name:
				plot6 = append(plot6, opts.LineData{Value: yValue})
			case plot7Name:
				plot7 = append(plot7, opts.LineData{Value: yValue})
			}
		}
	}

	fmt.Printf("X: %v\n", xAxis)
	// Create a new line chart instance
	line := charts.NewLine()

	// Set chart options
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeMacarons}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
			Top:  "96%",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Operations Per Second",
			Subtitle: "Comparison with different CPU counts",
			Left:     "center",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Operations Per Second",
			NameLocation: "center",
			NameGap:      70,
		}, 0),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "CPU's",
			NameLocation: "center",
			NameGap:      30,
		}, 0),
	)

	line.SetXAxis(xAxis).
		AddSeries(plot1Name, plot1).
		AddSeries(plot2Name, plot2).
		AddSeries(plot3Name, plot3).
		AddSeries(plot4Name, plot4).
		AddSeries(plot5Name, plot5).
		AddSeries(plot6Name, plot6).
		AddSeries(plot7Name, plot7).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				//Show: true,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth:     true,
				ShowSymbol: true,
			}),
		)

	// Save the chart
	f, err := os.Create("operations_per_second.html")
	if err != nil {
		panic(err)
	}
	if err = line.Render(f); err != nil {
		panic(err)
	}
}
