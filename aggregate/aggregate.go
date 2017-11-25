package aggregate

const StartSize = 100

var (
    aggregates = make([]Aggregator, 0, StartSize)
)

type Aggregator interface {
    GetVal() int
    GetName() string
    Up()

    Refresh()
}

func AddAggregator(aggregators ...Aggregator) {
    aggregates = append(aggregates, aggregators...)
}

func GetAllAggregators() []Aggregator {
    return aggregates
}