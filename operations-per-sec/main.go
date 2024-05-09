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
Current Operating System has '32' CPUs
goos: linux
goarch: amd64
pkg: github.com/gubernator-io/gubernator/v2/bench
cpu: AMD EPYC 7B13
BenchmarkParallel/Hashi_1-32         	 2485050	       490.0 ns/op	   2040969 ops/s
testing: BenchmarkParallel/Hashi_1-32 left GOMAXPROCS set to 1
BenchmarkParallel/Hashi_2-32         	 2693011	       483.6 ns/op	   2067917 ops/s
testing: BenchmarkParallel/Hashi_2-32 left GOMAXPROCS set to 2
BenchmarkParallel/Hashi_4-32         	 1801766	       645.0 ns/op	   1550184 ops/s
testing: BenchmarkParallel/Hashi_4-32 left GOMAXPROCS set to 4
BenchmarkParallel/Hashi_8-32         	 1504105	       800.6 ns/op	   1248989 ops/s
testing: BenchmarkParallel/Hashi_8-32 left GOMAXPROCS set to 8
BenchmarkParallel/Hashi_12-32        	 1414710	       920.3 ns/op	   1086567 ops/s
testing: BenchmarkParallel/Hashi_12-32 left GOMAXPROCS set to 12
BenchmarkParallel/Hashi_16-32        	 1283340	       937.1 ns/op	   1067084 ops/s
testing: BenchmarkParallel/Hashi_16-32 left GOMAXPROCS set to 16
BenchmarkParallel/Hashi_20-32        	 1000000	      1022 ns/op	    978151 ops/s
testing: BenchmarkParallel/Hashi_20-32 left GOMAXPROCS set to 20
BenchmarkParallel/Hashi_24-32        	 1000000	      1006 ns/op	    994300 ops/s
testing: BenchmarkParallel/Hashi_24-32 left GOMAXPROCS set to 24
BenchmarkParallel/Hashi_28-32        	 1000000	      1008 ns/op	    991721 ops/s
testing: BenchmarkParallel/Hashi_28-32 left GOMAXPROCS set to 28
BenchmarkParallel/Hashi_32-32        	 1000000	      1019 ns/op	    981560 ops/s
BenchmarkParallel/Mutex_1-32         	 2444936	       487.7 ns/op	   2050474 ops/s
testing: BenchmarkParallel/Mutex_1-32 left GOMAXPROCS set to 1
BenchmarkParallel/Mutex_2-32         	 2254138	       496.5 ns/op	   2014016 ops/s
testing: BenchmarkParallel/Mutex_2-32 left GOMAXPROCS set to 2
BenchmarkParallel/Mutex_4-32         	 2306122	       524.6 ns/op	   1906233 ops/s
testing: BenchmarkParallel/Mutex_4-32 left GOMAXPROCS set to 4
BenchmarkParallel/Mutex_8-32         	 2294396	       525.0 ns/op	   1904735 ops/s
testing: BenchmarkParallel/Mutex_8-32 left GOMAXPROCS set to 8
BenchmarkParallel/Mutex_12-32        	 2217972	       527.3 ns/op	   1896342 ops/s
testing: BenchmarkParallel/Mutex_12-32 left GOMAXPROCS set to 12
BenchmarkParallel/Mutex_16-32        	 2195224	       541.1 ns/op	   1847877 ops/s
testing: BenchmarkParallel/Mutex_16-32 left GOMAXPROCS set to 16
BenchmarkParallel/Mutex_20-32        	 2191591	       542.2 ns/op	   1844286 ops/s
testing: BenchmarkParallel/Mutex_20-32 left GOMAXPROCS set to 20
BenchmarkParallel/Mutex_24-32        	 2229174	       544.6 ns/op	   1835989 ops/s
testing: BenchmarkParallel/Mutex_24-32 left GOMAXPROCS set to 24
BenchmarkParallel/Mutex_28-32        	 2162965	       549.6 ns/op	   1819528 ops/s
testing: BenchmarkParallel/Mutex_28-32 left GOMAXPROCS set to 28
BenchmarkParallel/Mutex_32-32        	 2157846	       553.2 ns/op	   1807525 ops/s
BenchmarkParallel/OtterRead_1-32     	 2441359	       515.9 ns/op	   1938281 ops/s
testing: BenchmarkParallel/OtterRead_1-32 left GOMAXPROCS set to 1
BenchmarkParallel/OtterRead_2-32     	 4589662	       250.7 ns/op	   3988275 ops/s
testing: BenchmarkParallel/OtterRead_2-32 left GOMAXPROCS set to 2
BenchmarkParallel/OtterRead_4-32     	 6500647	       159.2 ns/op	   6280530 ops/s
testing: BenchmarkParallel/OtterRead_4-32 left GOMAXPROCS set to 4
BenchmarkParallel/OtterRead_8-32     	11865475	       100.7 ns/op	   9929738 ops/s
testing: BenchmarkParallel/OtterRead_8-32 left GOMAXPROCS set to 8
BenchmarkParallel/OtterRead_12-32    	14165311	        85.06 ns/op	  11756120 ops/s
testing: BenchmarkParallel/OtterRead_12-32 left GOMAXPROCS set to 12
BenchmarkParallel/OtterRead_16-32    	15390075	        77.08 ns/op	  12973182 ops/s
testing: BenchmarkParallel/OtterRead_16-32 left GOMAXPROCS set to 16
BenchmarkParallel/OtterRead_20-32    	15842086	        75.33 ns/op	  13275598 ops/s
testing: BenchmarkParallel/OtterRead_20-32 left GOMAXPROCS set to 20
BenchmarkParallel/OtterRead_24-32    	16291057	        75.03 ns/op	  13327300 ops/s
testing: BenchmarkParallel/OtterRead_24-32 left GOMAXPROCS set to 24
BenchmarkParallel/OtterRead_28-32    	16135291	        73.02 ns/op	  13694685 ops/s
testing: BenchmarkParallel/OtterRead_28-32 left GOMAXPROCS set to 28
BenchmarkParallel/OtterRead_32-32    	16027861	        74.40 ns/op	  13440457 ops/s
BenchmarkParallel/OtterWrite_1-32    	 2171982	       520.2 ns/op	   1922382 ops/s
testing: BenchmarkParallel/OtterWrite_1-32 left GOMAXPROCS set to 1
BenchmarkParallel/OtterWrite_2-32    	 3306013	       324.7 ns/op	   3079558 ops/s
testing: BenchmarkParallel/OtterWrite_2-32 left GOMAXPROCS set to 2
BenchmarkParallel/OtterWrite_4-32    	 7373560	       156.8 ns/op	   6376510 ops/s
testing: BenchmarkParallel/OtterWrite_4-32 left GOMAXPROCS set to 4
BenchmarkParallel/OtterWrite_8-32    	10692710	        96.15 ns/op	  10400615 ops/s
testing: BenchmarkParallel/OtterWrite_8-32 left GOMAXPROCS set to 8
BenchmarkParallel/OtterWrite_12-32   	16016918	        76.52 ns/op	  13068634 ops/s
testing: BenchmarkParallel/OtterWrite_12-32 left GOMAXPROCS set to 12
BenchmarkParallel/OtterWrite_16-32   	15529743	        67.39 ns/op	  14839551 ops/s
testing: BenchmarkParallel/OtterWrite_16-32 left GOMAXPROCS set to 16
BenchmarkParallel/OtterWrite_20-32   	19350819	        61.86 ns/op	  16165578 ops/s
testing: BenchmarkParallel/OtterWrite_20-32 left GOMAXPROCS set to 20
BenchmarkParallel/OtterWrite_24-32   	16762803	        59.72 ns/op	  16744224 ops/s
testing: BenchmarkParallel/OtterWrite_24-32 left GOMAXPROCS set to 24
BenchmarkParallel/OtterWrite_28-32   	18298399	        56.57 ns/op	  17677667 ops/s
testing: BenchmarkParallel/OtterWrite_28-32 left GOMAXPROCS set to 28
BenchmarkParallel/OtterWrite_32-32   	19195123	        55.31 ns/op	  18079519 ops/s
BenchmarkParallel/WorkerPool_1-32    	  603769	      2455 ns/op	    407327 ops/s
testing: BenchmarkParallel/WorkerPool_1-32 left GOMAXPROCS set to 1
BenchmarkParallel/WorkerPool_2-32    	  630951	      2582 ns/op	    387298 ops/s
testing: BenchmarkParallel/WorkerPool_2-32 left GOMAXPROCS set to 2
BenchmarkParallel/WorkerPool_4-32    	  580464	      2448 ns/op	    408512 ops/s
testing: BenchmarkParallel/WorkerPool_4-32 left GOMAXPROCS set to 4
BenchmarkParallel/WorkerPool_8-32    	  589813	      2013 ns/op	    496650 ops/s
testing: BenchmarkParallel/WorkerPool_8-32 left GOMAXPROCS set to 8
BenchmarkParallel/WorkerPool_12-32   	  594933	      1980 ns/op	    504964 ops/s
testing: BenchmarkParallel/WorkerPool_12-32 left GOMAXPROCS set to 12
BenchmarkParallel/WorkerPool_16-32   	  582597	      2154 ns/op	    464229 ops/s
testing: BenchmarkParallel/WorkerPool_16-32 left GOMAXPROCS set to 16
BenchmarkParallel/WorkerPool_20-32   	  559522	      2327 ns/op	    429637 ops/s
testing: BenchmarkParallel/WorkerPool_20-32 left GOMAXPROCS set to 20
BenchmarkParallel/WorkerPool_24-32   	  537469	      2196 ns/op	    455392 ops/s
testing: BenchmarkParallel/WorkerPool_24-32 left GOMAXPROCS set to 24
BenchmarkParallel/WorkerPool_28-32   	  541045	      2165 ns/op	    461917 ops/s
testing: BenchmarkParallel/WorkerPool_28-32 left GOMAXPROCS set to 28
BenchmarkParallel/WorkerPool_32-32   	  396483	      2629 ns/op	    380292 ops/s
PASS
ok  	github.com/gubernator-io/gubernator/v2/bench	102.331s

	`)

	plot1 := make([]opts.LineData, 0)
	//plot2 := make([]opts.LineData, 0)
	//plot3 := make([]opts.LineData, 0)
	plot4 := make([]opts.LineData, 0)
	//plot5 := make([]opts.LineData, 0)
	plot6 := make([]opts.LineData, 0)
	plot7 := make([]opts.LineData, 0)
	plot8 := make([]opts.LineData, 0)

	const (
		plot1Name = "Mutex"
		//plot2Name = "Mutex_Write"
		//plot3Name = "NoCache"
		plot4Name = "WorkerPool"
		//plot5Name = "WorkerPool_Write"
		plot6Name = "OtterRead"
		plot7Name = "OtterWrite"
		plot8Name = "Hashi"
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
			//case plot2Name:
			//	plot2 = append(plot2, opts.LineData{Value: yValue})
			//case plot3Name:
			//	plot3 = append(plot3, opts.LineData{Value: yValue})
			case plot4Name:
				plot4 = append(plot4, opts.LineData{Value: yValue})
			//case plot5Name:
			//	plot5 = append(plot5, opts.LineData{Value: yValue})
			case plot6Name:
				plot6 = append(plot6, opts.LineData{Value: yValue})
			case plot7Name:
				plot7 = append(plot7, opts.LineData{Value: yValue})
			case plot8Name:
				plot8 = append(plot8, opts.LineData{Value: yValue})
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
		//AddSeries(plot2Name, plot2).
		//AddSeries(plot3Name, plot3).
		AddSeries(plot4Name, plot4).
		//AddSeries(plot5Name, plot5).
		AddSeries(plot6Name, plot6).
		AddSeries(plot7Name, plot7).
		AddSeries(plot8Name+" LRU", plot8).
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
