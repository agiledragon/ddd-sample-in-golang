package base

type Entity struct {
	id string
}

func NewEntity(id string) Entity {
	return Entity{id: id}
}

func (t *Entity) Id() string {
	return t.id
}

func (t *Entity) Equal(other *Entity) bool {
	return t.id == other.id
}

func (t *Entity) NotEqual(other *Entity) bool {
	return !t.Equal(other)
}
