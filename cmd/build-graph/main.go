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
	// Prepare data
	plot1 := make([]opts.LineData, 0)
	plot2 := make([]opts.LineData, 0)

	file := bytes.Buffer{}
	file.WriteString(`
goos: darwin
goarch: arm64
pkg: github.com/gubernator-io/gubernator/v2/bench
BenchmarkWorkerPool
BenchmarkWorkerPool/Read_1
INFO[0000] Starting 10 Gubernator workers...            
INFO[0003] Starting 10 Gubernator workers...            
INFO[0007] Starting 10 Gubernator workers...            
INFO[0010] Starting 10 Gubernator workers...            
INFO[0015] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_1-10         	  471453	      2545 ns/op
BenchmarkWorkerPool/Read_10
INFO[0020] Starting 10 Gubernator workers...            
INFO[0023] Starting 10 Gubernator workers...            
INFO[0027] Starting 10 Gubernator workers...            
INFO[0031] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_10-10        	   44229	     31061 ns/op
BenchmarkWorkerPool/Read_100
INFO[0036] Starting 10 Gubernator workers...            
INFO[0040] Starting 10 Gubernator workers...            
INFO[0044] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_100-10       	    4228	    338240 ns/op
BenchmarkWorkerPool/Read_1000
INFO[0049] Starting 10 Gubernator workers...            
INFO[0053] Starting 10 Gubernator workers...            
INFO[0058] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_1000-10      	     314	   3407452 ns/op
BenchmarkWorkerPool/Read_10000
INFO[0063] Starting 10 Gubernator workers...            
INFO[0067] Starting 10 Gubernator workers...            
INFO[0072] Starting 10 Gubernator workers...            
INFO[0077] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_10000-10     	      34	  42071989 ns/op
BenchmarkWorkerPool/Read_100000
INFO[0083] Starting 10 Gubernator workers...            
INFO[0088] Starting 10 Gubernator workers...            
INFO[0093] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Read_100000-10    	       3	 411317972 ns/op
BenchmarkWorkerPool/Write_1
INFO[0099] Starting 10 Gubernator workers...            
INFO[0100] Starting 10 Gubernator workers...            
INFO[0101] Starting 10 Gubernator workers...            
INFO[0103] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_1-10        	  504410	      2529 ns/op
BenchmarkWorkerPool/Write_10
INFO[0105] Starting 10 Gubernator workers...            
INFO[0107] Starting 10 Gubernator workers...            
INFO[0108] Starting 10 Gubernator workers...            
INFO[0109] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_10-10       	   50311	     30179 ns/op
BenchmarkWorkerPool/Write_100
INFO[0112] Starting 10 Gubernator workers...            
INFO[0114] Starting 10 Gubernator workers...            
INFO[0115] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_100-10      	    2990	    337452 ns/op
BenchmarkWorkerPool/Write_1000
INFO[0117] Starting 10 Gubernator workers...            
INFO[0119] Starting 10 Gubernator workers...            
INFO[0120] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_1000-10     	     307	   3521005 ns/op
BenchmarkWorkerPool/Write_10000
INFO[0123] Starting 10 Gubernator workers...            
INFO[0124] Starting 10 Gubernator workers...            
INFO[0126] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_10000-10    	      38	  32510420 ns/op
BenchmarkWorkerPool/Write_100000
INFO[0129] Starting 10 Gubernator workers...            
INFO[0130] Starting 10 Gubernator workers...            
INFO[0133] Starting 10 Gubernator workers...            
BenchmarkWorkerPool/Write_100000-10   	       3	 417362694 ns/op
PASS
	`)

	const (
		plot1Name = "Read"
		plot2Name = "Write"
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

			if matches[1] == plot1Name {
				xAxis = append(xAxis, matches[2])
				plot1 = append(plot1, opts.LineData{Value: perOp})
			} else if matches[1] == plot2Name {
				plot2 = append(plot2, opts.LineData{Value: perOp})
			}
		}
	}

	// Create a new line chart instance
	line := charts.NewLine()

	// Set chart options
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Concurrency Benchmark",
			Subtitle: "Performance comparison at different concurrency",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Milliseconds",
			NameLocation: "center",
			NameGap:      40,
		}, 0),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Iterations",
			NameLocation: "center",
			NameGap:      30,
		}, 0),
	)

	line.SetXAxis(xAxis).
		AddSeries(plot1Name, plot1).
		AddSeries(plot2Name, plot2).
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
