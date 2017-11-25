package main

import (
    "log"
    "github.com/NabAlex/go-stats/aggregate"
    "net/http"
    "time"

    "math/rand"
    "github.com/NabAlex/go-stats"
    "github.com/NabAlex/easyconfig"
)

var graphiteServer = easyconfig.GetString("graphite.server", "95.163.249.148:2003")

func GenerateDigitToken(len int) string {
    str := "1234567890"
    shuffleArray := make([]byte, len)

    rand.Seed(time.Now().Unix())
    for i := 0; i < len; i++ {
        index := rand.Intn(10)
        shuffleArray[i] = str[index]
    }

    return string(shuffleArray)
}

var tokenCounter = aggregate.CreateCounter("example.counter")
var tokenCounterTwice = aggregate.CreateCounter("example.counter_twice")
func generateToken(w http.ResponseWriter, r *http.Request) {
    tokenCounter.Up()

    tokenCounterTwice.Up()
    tokenCounterTwice.Up()

    w.Write([]byte(GenerateDigitToken(32)))
}

func main() {
    aggregate.AddAggregator(tokenCounter, tokenCounterTwice)

    c, err := gostats.CreateGraphiteClient(gostats.UdpProtocol, graphiteServer)
    if err != nil {
        log.Panicln(err)
    }

    err = gostats.InitMetric(c)
    if err != nil {
        log.Panicln(err)
    }

    http.HandleFunc("/token", generateToken) // set router
    err = http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Panicln("Start listena: ", err)
    }
}