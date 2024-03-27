package main

import (
	"bufio"
	"bytes"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile(`BenchmarkAccessStructure/(Array|Hash)_(\d+)-\d+\s+(\d+)\s+(\d+\.\d+) ns/op`)

	scanner := bufio.NewScanner(&file)
	var xAxis []string
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 {
			nsPerOp, _ := strconv.ParseFloat(matches[4], 64)
			//msPerOp := nsPerOp / 1000000

			if matches[1] == "Array" {
				xAxis = append(xAxis, matches[2])
				plot1 = append(plot1, opts.LineData{Value: nsPerOp})
			} else if matches[1] == "Hash" {
				plot2 = append(plot2, opts.LineData{Value: nsPerOp})
			}
		}
	}

	// Create a new line chart instance
	line := charts.NewLine()

	// Set chart options
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Benchmark Results: Array vs Hash",
			Subtitle: "Performance comparison in ns/op",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Nanoseconds",
			NameLocation: "center",
			NameGap:      30,
		}, 0),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Iterations",
			NameLocation: "center",
			NameGap:      30,
		}, 0),
	)

	line.SetXAxis(xAxis).
		AddSeries("Array", plot1).
		AddSeries("Hash", plot2).
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
