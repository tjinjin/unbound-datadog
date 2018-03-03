package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
)

type UnboundThread struct {
	NumQueries              int
	NumQueriesIpRatelimited int
	NumCachehits            int
	NumCachemiss            int
	NumPrefetch             int
	NumZeroTtl              int
	NumRecursivereplies     int
	RequestlistAvg          int
	RequestlistMax          int
	RequestlistOverwritten  int
	RequestlistExceeded     int
	RequestlistCurrentAll   int
	RequestlistCurrentUser  int
	RecursionTimeAvg        float64
	RecursionTimeMedian     int
	Tcpusage                int
}

type Unbound struct {
	threadN                 []UnboundThread
	NumQueries              int
	NumQueriesIpRatelimited int
	NumCachehits            int
	NumCachemiss            int
	NumPrefetch             int
	NumZeroTtl              int
	NumRecursivereplies     int
	RequestlistAvg          int
	RequestlistMax          int
	RequestlistOverwritten  int
	RequestlistExceeded     int
	RequestlistCurrentAll   int
	RequestlistCurrentUser  int
	RecursionTimeAvg        float64
	RecursionTimeMedian     int
	Tcpusage                int
	timeNow                 float64
	timeUp                  float64
	timeElapsed             float64
}

// Todo
func datadogSample() {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	c.Namespace = "unbound."
	c.Tags = append(c.Tags, "test")
	err = c.Gauge("request.duration", 1.2, nil, 1)
	if err != nil {
		log.Fatal(err)
	}
}

// Todo
func sendDatadog(name string, value float64) {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	c.Namespace = "unbound"
	c.Tags = append(c.Tags, "")
	err = c.Gauge(name, value, nil, 1)
	if err != nil {
		log.Fatal(err)
	}
}

// Todo
func convertDotNotation(str string) string {
	fmt.Println("test")
	return "test"
}

func stringToFloat64(str string) float64 {
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

func stringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func execUnboundControl() {
	out, err := exec.Command("unbound-control", "stats").Output()
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(string(out), "\n")

	metrics := Unbound{}
	for _, v := range strs {
		if len(v) == 0 {
			break
		}
		values := strings.Split(v, "=")
		switch values[0] {
		case "total.num.queries_ip_ratelimited":
			metrics.NumQueriesIpRatelimited = stringToInt(values[1])
		case "total.num.cachehits":
			metrics.NumCachehits = stringToInt(values[1])
		case "total.num.cachemiss ":
			metrics.NumCachemiss = stringToInt(values[1])
		case "total.num.prefetch ":
			metrics.NumPrefetch = stringToInt(values[1])
		case "total.num.zero_ttl ":
			metrics.NumZeroTtl = stringToInt(values[1])
		case "total.num.recursivereplies ":
			metrics.NumRecursivereplies = stringToInt(values[1])
		case "total.requestlist.avg ":
			metrics.RequestlistAvg = stringToInt(values[1])
		case "total.requestlist.max ":
			metrics.RequestlistMax = stringToInt(values[1])
		case "total.requestlist.overwritten ":
			metrics.RequestlistOverwritten = stringToInt(values[1])
		case "total.requestlist.exceeded ":
			metrics.RequestlistExceeded = stringToInt(values[1])
		case "total.requestlist.current.all ":
			metrics.RequestlistCurrentAll = stringToInt(values[1])
		case "total.requestlist.current.user ":
			metrics.RequestlistCurrentUser = stringToInt(values[1])
		case "total.recursion.time.avg":
			metrics.RecursionTimeAvg = stringToFloat64(values[1])
		case "total.recursion.time.median":
			metrics.RecursionTimeMedian = stringToInt(values[1])
		case "total.tcpusage":
			metrics.Tcpusage = stringToInt(values[1])
		case "total.num.queries":
			metrics.NumQueries = stringToInt(values[1])
		case "time.now":
			metrics.timeNow = stringToFloat64(values[1])
		case "time.up":
			metrics.timeUp = stringToFloat64(values[1])
		case "time.elapsed":
			metrics.timeElapsed = stringToFloat64(values[1])
		}
	}
	fmt.Println(metrics)
}

func main() {
	execUnboundControl()
}
