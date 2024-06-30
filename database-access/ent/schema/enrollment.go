package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Enrollment holds the schema definition for the Enrollment entity.
type Enrollment struct {
    ent.Schema
}

// Fields of the Enrollment.
func (Enrollment) Fields() []ent.Field {
    return []ent.Field{
        field.String("semester").NotEmpty(),
        field.Int("year").Positive(),
    }
}

// Edges of the Enrollment.
func (Enrollment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("student", Student.Type).
            Ref("enrollments").
            Unique().
            Required(),
        edge.From("course", Course.Type).
            Ref("enrollments").
            Unique().
            Required(),
    }
}

