package base

type AggregateRoot struct {
	Entity
}

func NewAggregateRoot(id string) AggregateRoot {
	return AggregateRoot{NewEntity(id)}
}
