package engine

//Entity is the base for all objects containing it's id
type Entity struct {
	id string
}

//NewEntity creates a new Entity
func NewEntity(id string) Entity {
	ent := Entity{id: id}
	return ent
}

//ID returns the entity id
func (ent *Entity) ID() string {
	return ent.id
}