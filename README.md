# What Dis?
It's a project to take benchmark output from gubernator and graph it using https://github.com/go-echarts/go-echarts

### Step 1
Edit `cmd/build-graph/main.go` and add the output from the benchmark into the code (you will see the place)

### Step 2
```bash
$ go run ./cmd/build-graph
```

### Step 3
```bash
$ open benchmark_plot.html
```

### Step 4
Profit!
