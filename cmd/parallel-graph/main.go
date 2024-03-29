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
BenchmarkParallel/NoCache_1-10         	  540829	      2476 ns/op
testing: BenchmarkParallel/NoCache_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/NoCache_2
BenchmarkParallel/NoCache_2-10         	  935966	      1521 ns/op
testing: BenchmarkParallel/NoCache_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/NoCache_4
BenchmarkParallel/NoCache_4-10         	 1000000	      1381 ns/op
testing: BenchmarkParallel/NoCache_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/NoCache_8
BenchmarkParallel/NoCache_8-10         	  988514	      1263 ns/op
testing: BenchmarkParallel/NoCache_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/NoCache_10
BenchmarkParallel/NoCache_10-10        	  992695	      1373 ns/op
BenchmarkParallel/NoCache_12
BenchmarkParallel/NoCache_12-10        	  882828	      1369 ns/op
testing: BenchmarkParallel/NoCache_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/NoCache_16
BenchmarkParallel/NoCache_16-10        	  876135	      1367 ns/op
testing: BenchmarkParallel/NoCache_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/NoCache_20
BenchmarkParallel/NoCache_20-10        	  871648	      1377 ns/op
testing: BenchmarkParallel/NoCache_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/NoCache_24
BenchmarkParallel/NoCache_24-10        	  883058	      1407 ns/op
testing: BenchmarkParallel/NoCache_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/NoCache_28
BenchmarkParallel/NoCache_28-10        	  840752	      1454 ns/op
testing: BenchmarkParallel/NoCache_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/NoCache_32
BenchmarkParallel/NoCache_32-10        	  803805	      1529 ns/op
testing: BenchmarkParallel/NoCache_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/MutexRead_1
BenchmarkParallel/MutexRead_1-10       	  578823	      2770 ns/op
testing: BenchmarkParallel/MutexRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/MutexRead_2
BenchmarkParallel/MutexRead_2-10       	  782298	      1802 ns/op
testing: BenchmarkParallel/MutexRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/MutexRead_4
BenchmarkParallel/MutexRead_4-10       	 1000000	      1580 ns/op
testing: BenchmarkParallel/MutexRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/MutexRead_8
BenchmarkParallel/MutexRead_8-10       	  959151	      1633 ns/op
testing: BenchmarkParallel/MutexRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/MutexRead_10
BenchmarkParallel/MutexRead_10-10      	  915333	      1827 ns/op
BenchmarkParallel/MutexRead_12
BenchmarkParallel/MutexRead_12-10      	  792325	      1753 ns/op
testing: BenchmarkParallel/MutexRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/MutexRead_16
BenchmarkParallel/MutexRead_16-10      	  867234	      1732 ns/op
testing: BenchmarkParallel/MutexRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/MutexRead_20
BenchmarkParallel/MutexRead_20-10      	  923047	      1671 ns/op
testing: BenchmarkParallel/MutexRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/MutexRead_24
BenchmarkParallel/MutexRead_24-10      	  999052	      1652 ns/op
testing: BenchmarkParallel/MutexRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/MutexRead_28
BenchmarkParallel/MutexRead_28-10      	  977300	      1519 ns/op
testing: BenchmarkParallel/MutexRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/MutexRead_32
BenchmarkParallel/MutexRead_32-10      	 1000000	      1503 ns/op
testing: BenchmarkParallel/MutexRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/MutexWrite_1
BenchmarkParallel/MutexWrite_1-10      	  323600	      3496 ns/op
testing: BenchmarkParallel/MutexWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/MutexWrite_2
BenchmarkParallel/MutexWrite_2-10      	  632142	      1808 ns/op
testing: BenchmarkParallel/MutexWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/MutexWrite_4
BenchmarkParallel/MutexWrite_4-10      	  722974	      1553 ns/op
testing: BenchmarkParallel/MutexWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/MutexWrite_8
BenchmarkParallel/MutexWrite_8-10      	  784740	      1494 ns/op
testing: BenchmarkParallel/MutexWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/MutexWrite_10
BenchmarkParallel/MutexWrite_10-10     	  800350	      1472 ns/op
BenchmarkParallel/MutexWrite_12
BenchmarkParallel/MutexWrite_12-10     	  758217	      1371 ns/op
testing: BenchmarkParallel/MutexWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/MutexWrite_16
BenchmarkParallel/MutexWrite_16-10     	  779378	      1390 ns/op
testing: BenchmarkParallel/MutexWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/MutexWrite_20
BenchmarkParallel/MutexWrite_20-10     	  731805	      1431 ns/op
testing: BenchmarkParallel/MutexWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/MutexWrite_24
BenchmarkParallel/MutexWrite_24-10     	  731206	      1426 ns/op
testing: BenchmarkParallel/MutexWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/MutexWrite_28
BenchmarkParallel/MutexWrite_28-10     	  753448	      1387 ns/op
testing: BenchmarkParallel/MutexWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/MutexWrite_32
BenchmarkParallel/MutexWrite_32-10     	  746794	      1393 ns/op
testing: BenchmarkParallel/MutexWrite_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/OtterRead_1
BenchmarkParallel/OtterRead_1-10       	  434704	      3323 ns/op
testing: BenchmarkParallel/OtterRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/OtterRead_2
BenchmarkParallel/OtterRead_2-10       	  877093	      1808 ns/op
testing: BenchmarkParallel/OtterRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/OtterRead_4
BenchmarkParallel/OtterRead_4-10       	 1000000	      1138 ns/op
testing: BenchmarkParallel/OtterRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/OtterRead_8
BenchmarkParallel/OtterRead_8-10       	  990319	      1262 ns/op
testing: BenchmarkParallel/OtterRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/OtterRead_10
BenchmarkParallel/OtterRead_10-10      	  972058	      1368 ns/op
BenchmarkParallel/OtterRead_12
BenchmarkParallel/OtterRead_12-10      	  840284	      1529 ns/op
testing: BenchmarkParallel/OtterRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/OtterRead_16
BenchmarkParallel/OtterRead_16-10      	  799975	      1604 ns/op
testing: BenchmarkParallel/OtterRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/OtterRead_20
BenchmarkParallel/OtterRead_20-10      	  716205	      1870 ns/op
testing: BenchmarkParallel/OtterRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/OtterRead_24
BenchmarkParallel/OtterRead_24-10      	  841615	      1945 ns/op
testing: BenchmarkParallel/OtterRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/OtterRead_28
BenchmarkParallel/OtterRead_28-10      	  631198	      2111 ns/op
testing: BenchmarkParallel/OtterRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/OtterRead_32
BenchmarkParallel/OtterRead_32-10      	  563864	      2197 ns/op
testing: BenchmarkParallel/OtterRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/OtterWrite_1
BenchmarkParallel/OtterWrite_1-10      	  448948	      2610 ns/op
testing: BenchmarkParallel/OtterWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/OtterWrite_2
BenchmarkParallel/OtterWrite_2-10      	  759397	      1733 ns/op
testing: BenchmarkParallel/OtterWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/OtterWrite_4
BenchmarkParallel/OtterWrite_4-10      	  985896	      1292 ns/op
testing: BenchmarkParallel/OtterWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/OtterWrite_8
BenchmarkParallel/OtterWrite_8-10      	 1000000	      1245 ns/op
testing: BenchmarkParallel/OtterWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/OtterWrite_10
BenchmarkParallel/OtterWrite_10-10     	 1000000	      1380 ns/op
BenchmarkParallel/OtterWrite_12
BenchmarkParallel/OtterWrite_12-10     	 1000000	      1495 ns/op
testing: BenchmarkParallel/OtterWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/OtterWrite_16
BenchmarkParallel/OtterWrite_16-10     	 1000000	      1672 ns/op
testing: BenchmarkParallel/OtterWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/OtterWrite_20
BenchmarkParallel/OtterWrite_20-10     	 1000000	      1821 ns/op
testing: BenchmarkParallel/OtterWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/OtterWrite_24
BenchmarkParallel/OtterWrite_24-10     	 1000000	      1977 ns/op
testing: BenchmarkParallel/OtterWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/OtterWrite_28
BenchmarkParallel/OtterWrite_28-10     	  978598	      2130 ns/op
testing: BenchmarkParallel/OtterWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/OtterWrite_32
BenchmarkParallel/OtterWrite_32-10     	  929874	      2209 ns/op
testing: BenchmarkParallel/OtterWrite_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/WorkerPoolRead_1
INFO[0095] Starting 1 Gubernator workers...             
INFO[0095] Starting 1 Gubernator workers...             
INFO[0095] Starting 1 Gubernator workers...             
INFO[0096] Starting 1 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_1-10  	  336408	      3797 ns/op
testing: BenchmarkParallel/WorkerPoolRead_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/WorkerPoolRead_2
INFO[0097] Starting 2 Gubernator workers...             
INFO[0097] Starting 2 Gubernator workers...             
INFO[0098] Starting 2 Gubernator workers...             
INFO[0098] Starting 2 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_2-10  	  585507	      1925 ns/op
testing: BenchmarkParallel/WorkerPoolRead_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/WorkerPoolRead_4
INFO[0099] Starting 4 Gubernator workers...             
INFO[0099] Starting 4 Gubernator workers...             
INFO[0100] Starting 4 Gubernator workers...             
INFO[0100] Starting 4 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_4-10  	  812762	      1660 ns/op
testing: BenchmarkParallel/WorkerPoolRead_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/WorkerPoolRead_8
INFO[0101] Starting 8 Gubernator workers...             
INFO[0101] Starting 8 Gubernator workers...             
INFO[0102] Starting 8 Gubernator workers...             
INFO[0102] Starting 8 Gubernator workers...             
BenchmarkParallel/WorkerPoolRead_8-10  	  543258	      2963 ns/op
testing: BenchmarkParallel/WorkerPoolRead_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/WorkerPoolRead_10
INFO[0104] Starting 10 Gubernator workers...            
INFO[0104] Starting 10 Gubernator workers...            
INFO[0104] Starting 10 Gubernator workers...            
INFO[0104] Starting 10 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_10-10 	  718182	      3986 ns/op
BenchmarkParallel/WorkerPoolRead_12
INFO[0107] Starting 12 Gubernator workers...            
INFO[0107] Starting 12 Gubernator workers...            
INFO[0108] Starting 12 Gubernator workers...            
INFO[0108] Starting 12 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_12-10 	  813004	      4587 ns/op
testing: BenchmarkParallel/WorkerPoolRead_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/WorkerPoolRead_16
INFO[0112] Starting 16 Gubernator workers...            
INFO[0112] Starting 16 Gubernator workers...            
INFO[0112] Starting 16 Gubernator workers...            
INFO[0112] Starting 16 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_16-10 	  437653	      5253 ns/op
testing: BenchmarkParallel/WorkerPoolRead_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/WorkerPoolRead_20
INFO[0115] Starting 20 Gubernator workers...            
INFO[0115] Starting 20 Gubernator workers...            
INFO[0115] Starting 20 Gubernator workers...            
INFO[0115] Starting 20 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_20-10 	  554270	      4380 ns/op
testing: BenchmarkParallel/WorkerPoolRead_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/WorkerPoolRead_24
INFO[0118] Starting 24 Gubernator workers...            
INFO[0118] Starting 24 Gubernator workers...            
INFO[0118] Starting 24 Gubernator workers...            
INFO[0119] Starting 24 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_24-10 	  374275	      4183 ns/op
testing: BenchmarkParallel/WorkerPoolRead_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/WorkerPoolRead_28
INFO[0121] Starting 28 Gubernator workers...            
INFO[0121] Starting 28 Gubernator workers...            
INFO[0121] Starting 28 Gubernator workers...            
INFO[0121] Starting 28 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_28-10 	  438096	      5309 ns/op
testing: BenchmarkParallel/WorkerPoolRead_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/WorkerPoolRead_32
INFO[0124] Starting 32 Gubernator workers...            
INFO[0124] Starting 32 Gubernator workers...            
INFO[0124] Starting 32 Gubernator workers...            
INFO[0125] Starting 32 Gubernator workers...            
BenchmarkParallel/WorkerPoolRead_32-10 	  245386	      4342 ns/op
testing: BenchmarkParallel/WorkerPoolRead_32-10 left GOMAXPROCS set to 32
BenchmarkParallel/WorkerPoolWrite_1
INFO[0126] Starting 1 Gubernator workers...             
INFO[0126] Starting 1 Gubernator workers...             
INFO[0126] Starting 1 Gubernator workers...             
INFO[0127] Starting 1 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_1-10 	  348675	      3212 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_1-10 left GOMAXPROCS set to 1
BenchmarkParallel/WorkerPoolWrite_2
INFO[0129] Starting 2 Gubernator workers...             
INFO[0129] Starting 2 Gubernator workers...             
INFO[0129] Starting 2 Gubernator workers...             
INFO[0129] Starting 2 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_2-10 	  609136	      2085 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_2-10 left GOMAXPROCS set to 2
BenchmarkParallel/WorkerPoolWrite_4
INFO[0131] Starting 4 Gubernator workers...             
INFO[0131] Starting 4 Gubernator workers...             
INFO[0131] Starting 4 Gubernator workers...             
INFO[0132] Starting 4 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_4-10 	  772257	      1566 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_4-10 left GOMAXPROCS set to 4
BenchmarkParallel/WorkerPoolWrite_8
INFO[0133] Starting 8 Gubernator workers...             
INFO[0133] Starting 8 Gubernator workers...             
INFO[0133] Starting 8 Gubernator workers...             
INFO[0133] Starting 8 Gubernator workers...             
BenchmarkParallel/WorkerPoolWrite_8-10 	  348687	      2935 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_8-10 left GOMAXPROCS set to 8
BenchmarkParallel/WorkerPoolWrite_10
INFO[0135] Starting 10 Gubernator workers...            
INFO[0135] Starting 10 Gubernator workers...            
INFO[0135] Starting 10 Gubernator workers...            
INFO[0135] Starting 10 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_10-10         	  319131	      3641 ns/op
BenchmarkParallel/WorkerPoolWrite_12
INFO[0136] Starting 12 Gubernator workers...            
INFO[0137] Starting 12 Gubernator workers...            
INFO[0137] Starting 12 Gubernator workers...            
INFO[0137] Starting 12 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_12-10         	  286860	      4122 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_12-10 left GOMAXPROCS set to 12
BenchmarkParallel/WorkerPoolWrite_16
INFO[0138] Starting 16 Gubernator workers...            
INFO[0138] Starting 16 Gubernator workers...            
INFO[0138] Starting 16 Gubernator workers...            
INFO[0139] Starting 16 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_16-10         	  238358	      4208 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_16-10 left GOMAXPROCS set to 16
BenchmarkParallel/WorkerPoolWrite_20
INFO[0140] Starting 20 Gubernator workers...            
INFO[0140] Starting 20 Gubernator workers...            
INFO[0140] Starting 20 Gubernator workers...            
INFO[0140] Starting 20 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_20-10         	  219445	      4596 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_20-10 left GOMAXPROCS set to 20
BenchmarkParallel/WorkerPoolWrite_24
INFO[0141] Starting 24 Gubernator workers...            
INFO[0141] Starting 24 Gubernator workers...            
INFO[0142] Starting 24 Gubernator workers...            
INFO[0142] Starting 24 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_24-10         	  293625	      5828 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_24-10 left GOMAXPROCS set to 24
BenchmarkParallel/WorkerPoolWrite_28
INFO[0144] Starting 28 Gubernator workers...            
INFO[0144] Starting 28 Gubernator workers...            
INFO[0144] Starting 28 Gubernator workers...            
INFO[0144] Starting 28 Gubernator workers...            
INFO[0145] Starting 28 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_28-10         	  288228	      4877 ns/op
testing: BenchmarkParallel/WorkerPoolWrite_28-10 left GOMAXPROCS set to 28
BenchmarkParallel/WorkerPoolWrite_32
INFO[0147] Starting 32 Gubernator workers...            
INFO[0147] Starting 32 Gubernator workers...            
INFO[0147] Starting 32 Gubernator workers...            
INFO[0147] Starting 32 Gubernator workers...            
BenchmarkParallel/WorkerPoolWrite_32-10         	  228300	      5163 ns/op
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
	re := regexp.MustCompile(`Benchmark\w+/(\w+)_(\d+)-\d+\s+(\d+)\s+(\d+) ns/op`)

	scanner := bufio.NewScanner(&file)
	var xAxis []string
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			perOp, _ := strconv.ParseFloat(matches[4], 64)
			perOp = perOp / 1000000
			//perOp = perOp / 1000

			fmt.Printf("Match: %s\n", matches[1])
			switch matches[1] {
			case plot1Name:
				xAxis = append(xAxis, matches[2])
				plot1 = append(plot1, opts.LineData{Value: perOp})
			case plot2Name:
				plot2 = append(plot2, opts.LineData{Value: perOp})
			case plot3Name:
				plot3 = append(plot3, opts.LineData{Value: perOp})
			case plot4Name:
				plot4 = append(plot4, opts.LineData{Value: perOp})
			case plot5Name:
				plot5 = append(plot5, opts.LineData{Value: perOp})
			case plot6Name:
				plot6 = append(plot6, opts.LineData{Value: perOp})
			case plot7Name:
				plot7 = append(plot7, opts.LineData{Value: perOp})
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
			//Right:         "",
			//Bottom:        "",
			//Data:          nil,
			//Orient:        "",
			//InactiveColor: "",
			//Selected:      nil,
			//SelectedMode:  "",
			//Padding:       nil,
			//ItemWidth:     0,
			//ItemHeight:    0,
			//X:             "",
			//Y:             "",
			//Width:         "",
			//Height:        "",
			//Align:         "",
			//TextStyle:     nil,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Parallel Benchmark",
			Subtitle: "Performance comparison with different CPU counts",
			Left:     "center",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Milliseconds",
			NameLocation: "center",
			NameGap:      40,
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
	f, err := os.Create("benchmark_plot.html")
	if err != nil {
		panic(err)
	}
	if err = line.Render(f); err != nil {
		panic(err)
	}
}
