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

const (
	plot1Name = "Array"
	plot2Name = "Hash"
)

func main() {
	// Prepare data
	plot1 := make([]opts.LineData, 0)
	plot2 := make([]opts.LineData, 0)

	file := bytes.Buffer{}
	file.WriteString(`
goos: darwin
goarch: arm64
pkg: github.com/gubernator-io/gubernator-graphs
BenchmarkAccessStructure
BenchmarkAccessStructure/Array_1
BenchmarkAccessStructure/Array_1-10         	918694573	         1.299 ns/op
BenchmarkAccessStructure/Hash_1
BenchmarkAccessStructure/Hash_1-10          	457919066	         2.578 ns/op
BenchmarkAccessStructure/Array_10
BenchmarkAccessStructure/Array_10-10        	920779400	         1.290 ns/op
BenchmarkAccessStructure/Hash_10
BenchmarkAccessStructure/Hash_10-10         	225909036	         5.217 ns/op
BenchmarkAccessStructure/Array_100
BenchmarkAccessStructure/Array_100-10       	946178533	         1.276 ns/op
BenchmarkAccessStructure/Hash_100
BenchmarkAccessStructure/Hash_100-10        	217348762	         5.523 ns/op
BenchmarkAccessStructure/Array_1000
BenchmarkAccessStructure/Array_1000-10      	942012274	         1.269 ns/op
BenchmarkAccessStructure/Hash_1000
BenchmarkAccessStructure/Hash_1000-10       	192828970	         6.076 ns/op
BenchmarkAccessStructure/Array_10000
BenchmarkAccessStructure/Array_10000-10     	921357054	         1.270 ns/op
BenchmarkAccessStructure/Hash_10000
BenchmarkAccessStructure/Hash_10000-10      	42759596	        27.00 ns/op
BenchmarkAccessStructure/Array_100000
BenchmarkAccessStructure/Array_100000-10    	943079271	         1.268 ns/op
BenchmarkAccessStructure/Hash_100000
BenchmarkAccessStructure/Hash_100000-10     	37152421	        32.69 ns/op
BenchmarkAccessStructure/Array_1000000
BenchmarkAccessStructure/Array_1000000-10   	929155744	         1.290 ns/op
BenchmarkAccessStructure/Hash_1000000
BenchmarkAccessStructure/Hash_1000000-10    	23785140	        49.68 ns/op
PASS
	`)

	// Regular expression to match benchmark lines
	re := regexp.MustCompile(`Benchmark.*/(.*)_(\d+)-\d+\s+(\d+)\s+(\d+\.\d+) ns/op`)

	scanner := bufio.NewScanner(&file)
	var xAxis []string
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			perOp, _ := strconv.ParseFloat(matches[4], 64)
			//perOp = perOp / 1000000
			//perOp = perOp / 1000

			fmt.Printf("Match: '%s'\n", matches[1])
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
			Name:         "NanoSeconds",
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
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
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
