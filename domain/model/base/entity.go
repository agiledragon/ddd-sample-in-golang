package base

type Entity struct {
    Id  string
}

func NewEntity(id string) Entity {
    return Entity{Id: id}
}

func (t *Entity) GetId() string {
    return t.Id
}

func (t *Entity) Equal(other *Entity) bool {
    return t.Id == other.Id
}

func (t *Entity) NotEqual(other *Entity) bool {
    return !t.Equal(other)
}