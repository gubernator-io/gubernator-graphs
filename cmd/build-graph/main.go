package main

import (
	"bufio"
	"bytes"
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
BenchmarkConcurrency
BenchmarkConcurrency/MutexRead_1
BenchmarkConcurrency/MutexRead_1-10         	  866772	      1434 ns/op
BenchmarkConcurrency/MutexRead_10
BenchmarkConcurrency/MutexRead_10-10        	  106581	     12518 ns/op
BenchmarkConcurrency/MutexRead_100
BenchmarkConcurrency/MutexRead_100-10       	    9666	    125122 ns/op
BenchmarkConcurrency/MutexRead_1000
BenchmarkConcurrency/MutexRead_1000-10      	     715	   1593109 ns/op
BenchmarkConcurrency/MutexRead_5000
BenchmarkConcurrency/MutexRead_5000-10      	     127	   8235080 ns/op
BenchmarkConcurrency/MutexRead_10000
BenchmarkConcurrency/MutexRead_10000-10     	      69	  15378351 ns/op
BenchmarkConcurrency/MutexRead_15000
BenchmarkConcurrency/MutexRead_15000-10     	      45	  26539623 ns/op
BenchmarkConcurrency/MutexRead_20000
BenchmarkConcurrency/MutexRead_20000-10     	      33	  33666009 ns/op
BenchmarkConcurrency/MutexWrite_1
BenchmarkConcurrency/MutexWrite_1-10        	 1000000	      1375 ns/op
BenchmarkConcurrency/MutexWrite_10
BenchmarkConcurrency/MutexWrite_10-10       	  102925	     11622 ns/op
BenchmarkConcurrency/MutexWrite_100
BenchmarkConcurrency/MutexWrite_100-10      	   10534	    115577 ns/op
BenchmarkConcurrency/MutexWrite_1000
BenchmarkConcurrency/MutexWrite_1000-10     	     987	   1182021 ns/op
BenchmarkConcurrency/MutexWrite_5000
BenchmarkConcurrency/MutexWrite_5000-10     	     200	   5973197 ns/op
BenchmarkConcurrency/MutexWrite_10000
BenchmarkConcurrency/MutexWrite_10000-10    	      96	  12376698 ns/op
BenchmarkConcurrency/MutexWrite_15000
BenchmarkConcurrency/MutexWrite_15000-10    	      56	  22291136 ns/op
BenchmarkConcurrency/MutexWrite_20000
BenchmarkConcurrency/MutexWrite_20000-10    	      39	  31322339 ns/op
BenchmarkConcurrency/WorkerPoolRead_1
INFO[0125] Starting 10 Gubernator workers...            
INFO[0128] Starting 10 Gubernator workers...            
INFO[0132] Starting 10 Gubernator workers...            
INFO[0136] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_1-10    	  391741	      2572 ns/op
BenchmarkConcurrency/WorkerPoolRead_10
INFO[0140] Starting 10 Gubernator workers...            
INFO[0144] Starting 10 Gubernator workers...            
INFO[0147] Starting 10 Gubernator workers...            
INFO[0152] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_10-10   	   29166	     34597 ns/op
BenchmarkConcurrency/WorkerPoolRead_100
INFO[0156] Starting 10 Gubernator workers...            
INFO[0160] Starting 10 Gubernator workers...            
INFO[0164] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_100-10  	    3332	    373858 ns/op
BenchmarkConcurrency/WorkerPoolRead_1000
INFO[0169] Starting 10 Gubernator workers...            
INFO[0173] Starting 10 Gubernator workers...            
INFO[0177] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_1000-10 	     298	   3839185 ns/op
BenchmarkConcurrency/WorkerPoolRead_5000
INFO[0183] Starting 10 Gubernator workers...            
INFO[0187] Starting 10 Gubernator workers...            
INFO[0191] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_5000-10 	      69	  18828307 ns/op
BenchmarkConcurrency/WorkerPoolRead_10000
INFO[0197] Starting 10 Gubernator workers...            
INFO[0201] Starting 10 Gubernator workers...            
INFO[0206] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_10000-10         	      31	  34572151 ns/op
BenchmarkConcurrency/WorkerPoolRead_15000
INFO[0211] Starting 10 Gubernator workers...            
INFO[0216] Starting 10 Gubernator workers...            
INFO[0221] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_15000-10         	      20	  54786860 ns/op
BenchmarkConcurrency/WorkerPoolRead_20000
INFO[0226] Starting 10 Gubernator workers...            
INFO[0231] Starting 10 Gubernator workers...            
INFO[0236] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolRead_20000-10         	      15	  74307658 ns/op
BenchmarkConcurrency/WorkerPoolWrite_1
INFO[0242] Starting 10 Gubernator workers...            
INFO[0243] Starting 10 Gubernator workers...            
INFO[0244] Starting 10 Gubernator workers...            
INFO[0246] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_1-10            	  505108	      2497 ns/op
BenchmarkConcurrency/WorkerPoolWrite_10
INFO[0249] Starting 10 Gubernator workers...            
INFO[0250] Starting 10 Gubernator workers...            
INFO[0251] Starting 10 Gubernator workers...            
INFO[0253] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_10-10           	   34081	     33313 ns/op
BenchmarkConcurrency/WorkerPoolWrite_100
INFO[0256] Starting 10 Gubernator workers...            
INFO[0257] Starting 10 Gubernator workers...            
INFO[0259] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_100-10          	    2989	    351366 ns/op
BenchmarkConcurrency/WorkerPoolWrite_1000
INFO[0261] Starting 10 Gubernator workers...            
INFO[0263] Starting 10 Gubernator workers...            
INFO[0264] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_1000-10         	     322	   3390951 ns/op
BenchmarkConcurrency/WorkerPoolWrite_5000
INFO[0267] Starting 10 Gubernator workers...            
INFO[0268] Starting 10 Gubernator workers...            
INFO[0271] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_5000-10         	      64	  16556949 ns/op
BenchmarkConcurrency/WorkerPoolWrite_10000
INFO[0273] Starting 10 Gubernator workers...            
INFO[0275] Starting 10 Gubernator workers...            
INFO[0277] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_10000-10        	      32	  38536669 ns/op
BenchmarkConcurrency/WorkerPoolWrite_15000
INFO[0279] Starting 10 Gubernator workers...            
INFO[0281] Starting 10 Gubernator workers...            
INFO[0283] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_15000-10        	      21	  48107173 ns/op
BenchmarkConcurrency/WorkerPoolWrite_20000
INFO[0286] Starting 10 Gubernator workers...            
INFO[0287] Starting 10 Gubernator workers...            
INFO[0289] Starting 10 Gubernator workers...            
BenchmarkConcurrency/WorkerPoolWrite_20000-10        	      16	  77032016 ns/op
PASS
	`)

	plot1 := make([]opts.LineData, 0)
	plot2 := make([]opts.LineData, 0)
	plot3 := make([]opts.LineData, 0)
	plot4 := make([]opts.LineData, 0)

	const (
		plot1Name = "WorkerPoolRead"
		plot2Name = "WorkerPoolWrite"
		plot3Name = "MutexRead"
		plot4Name = "MutexWrite"
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
			}
		}
	}

	// Create a new line chart instance
	line := charts.NewLine()

	// Set chart options
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
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
			Title:    "Concurrency Benchmark",
			Subtitle: "Performance comparison of different concurrences",
			Left:     "center",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Milliseconds",
			NameLocation: "center",
			NameGap:      40,
		}, 0),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Concurrency",
			NameLocation: "center",
			NameGap:      30,
		}, 0),
	)

	line.SetXAxis(xAxis).
		AddSeries(plot1Name, plot1).
		AddSeries(plot2Name, plot2).
		AddSeries(plot3Name, plot3).
		AddSeries(plot4Name, plot4).
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
