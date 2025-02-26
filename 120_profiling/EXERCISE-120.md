# Exercises for chapter Profiling

## Intro 

Examine the given source code in this directory.

There are three implementations of the Fibonacci sequence.

Also have a look how exactly the benchmarks in `fib_test.go` have been implemented:

- what is `testing.TB`?
- why do we pass a function to `performTest`?

## Run profiling on test

Now collecting CPU profiling data for the **Test** (not the benchmarks):

```bash
go test -cpuprofile=cpu.out 
```

Then investigate the results with `go tool pprof`.

Investigate and experiment with the following commands:

- `top`
- `list`
- `web`
- `peek`
- `png`

What other commands exist? What do you find useful and want to share
with your colleagues later?

## Run profiling on benchmarks

Now collecting cpu and memory profiling data for a specific **Benchmark** 
without running any tests (remember that the "Rec" part below is a regular expression to
match functions beginning with "Benchmark"):

```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench Rec -run ^#  
```

Again, investigate the results with `go tool pprof`. 

Also, you can use your IDE to run the benchmark and investigate it.

Feel free to run other benchmarks then (e.g. `Fast` or `Memwaste`).




