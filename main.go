package main

import (
	"github.com/DataDog/datadog-go/statsd"
    "log"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type DDMetric struct {
    BackendCount int `json:"backendCount"`
}

func (m DDMetric) toString() string {
    return toJSON(m)
}

func toJSON(m interface{}) string {
    bytes, err := json.Marshal(m)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func getMetrics() []DDMetric {
    raw, err := ioutil.ReadFile("./metric.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []DDMetric
    json.Unmarshal(raw, &c)
    //fmt.Println(c)
    return c
}

func main() {

    c, err := statsd.New("127.0.0.1:8125")
    if err != nil {
        log.Fatal(err)
    }
    // prefix every metric with the app name
    c.Namespace = "allthingscloud."
    // send bananas as a tag with every metric
    c.Tags = append(c.Tags, "bananas")
    
    // read metrics in from json file
    metrics := getMetrics()

    // if there's multiple parameters or metrics
    for _, m := range metrics {
        fmt.Println(m.BackendCount)
    }

    // grab the first metric
    fmt.Println(metrics[0].BackendCount)

    fmt.Println(toJSON(metrics))

    // err = c.Gauge("backend_guage", 4, nil, 1)
    // err = c.Count("pagecount_total", 100, nil, 1)
    // err = c.Incr("pagecount_total", nil, 1)


}