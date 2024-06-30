package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
    ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("instructors", Instructor.Type),
        edge.To("courses", Course.Type),
        edge.To("students", Student.Type),
    }
}

