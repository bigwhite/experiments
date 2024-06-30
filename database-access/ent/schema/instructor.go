package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Instructor holds the schema definition for the Instructor entity.
type Instructor struct {
    ent.Schema
}

// Fields of the Instructor.
func (Instructor) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the Instructor.
func (Instructor) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("department", Department.Type).
            Ref("instructors").
            Unique().
            Required(),
    }
}

