package base

type Entity struct {
	id string
}

func NewEntity(id string) Entity {
	return Entity{id: id}
}

func (this *Entity) Id() string {
	return this.id
}

func (this *Entity) Equal(other *Entity) bool {
	return this.id == other.id
}

func (this *Entity) NotEqual(other *Entity) bool {
	return !this.Equal(other)
}
