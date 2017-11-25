package gostats

import (
    "github.com/NabAlex/go-stats/aggregate"
    "time"
    "errors"
)

var (
    client MetricClient
)

type MetricClient interface {
    Close()
    IsClose() bool

    sendAggregator(now time.Time, aggregator []aggregate.Aggregator)
}

func runUpdateMetric() {
    ticker := time.NewTicker(1 * time.Second)

    defer ticker.Stop()
    defer client.Close()

    for now := range ticker.C {
        allAggregators := aggregate.GetAllAggregators()
        client.sendAggregator(now, allAggregators)
        for _, a := range allAggregators {
            a.Refresh()
        }
    }
}

func InitMetric(clientMetric MetricClient) error {
    if clientMetric.IsClose() {
        return errors.New("close metric")
    }

    client = clientMetric

    go runUpdateMetric()
    return nil
}
