package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
)

type Reporter struct {
	Client *statsd.Client
}

func initClient() *Reporter {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	// namespace.metric
	c.Namespace = "unbound."
	c.Tags = append(c.Tags, "unbound")
	c.Tags = append(c.Tags, "ap-northeast-1")
	return &Reporter{
		Client: c,
	}
}

// Todo
func submitDogstatsD(name string, value float64, reporter *Reporter) {
	err := reporter.Client.Gauge(name, value, nil, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func stringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func execUnboundControl(reporter *Reporter) {
	out, err := exec.Command("unbound-control", "stats").Output()
	if err != nil {
		log.Fatal(err)
	}
	array := strings.Split(string(out), "\n")

	for _, v := range array {
		if len(v) == 0 {
			break
		}
		values := strings.Split(v, "=")
		switch values[0] {
		case "total.num.queries":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.queries_ip_ratelimited":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.cachehits":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.cachemiss":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.prefetch":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.zero_ttl":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.num.recursivereplies":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.avg":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.max":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.overwritten":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.exceeded":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.current.all":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.requestlist.current.user":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.recursion.time.avg":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.recursion.time.median":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		case "total.tcpusage":
			submitDogstatsD(values[0], stringToFloat64(values[1]), reporter)
		}
	}
}

func main() {
	reporter := initClient()
	execUnboundControl(reporter)
}
