package gostats

import (
    "github.com/NabAlex/go-stats/aggregate"
    "time"
    "net"
    "errors"
    "fmt"
    "log"
)

type GraphiteProtocol string
const (
    TcpProtocol GraphiteProtocol = "tcp"
    UdpProtocol GraphiteProtocol = "udp"
)

type GraphiteClient struct {
    protocol GraphiteProtocol
    conn net.Conn

    closed bool
}

var bucket = make([]byte, 0, 100 * 1024)
func (gc *GraphiteClient) sendAggregator(now time.Time, aggregators []aggregate.Aggregator) {
    /* simple */


    for _, agg := range aggregators {
        res := fmt.Sprintf("%s %d %d\n", agg.GetName(), agg.GetVal(), now.Unix())
        bucket = append(bucket, res...)

        log.Println(res)
    }

    _, err := gc.conn.Write(bucket)
    if err != nil {
        /* try reconnect */
        log.Println("cannot send to server")
    }
}

func (gc *GraphiteClient) Close() {
    gc.conn.Close()
    gc.closed = true
}

func (gc *GraphiteClient) IsClose() bool {
    return gc.closed
}

func CreateGraphiteClient(proto GraphiteProtocol, servername string) (MetricClient, error) {
    if proto == TcpProtocol {
        return nil, errors.New("cannot tcp protocol")
    }


    conn, err := net.DialTimeout(string(proto), servername, 5 * time.Second)
    if err != nil {
        return nil, err
    }

    c := new(GraphiteClient)
    c.conn = conn
    c.protocol = proto
    c.closed = false
    return c, nil
}