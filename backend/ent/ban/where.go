// Code generated by ent, DO NOT EDIT.

package ban

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ban {
	return predicate.Ban(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ban {
	return predicate.Ban(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ban {
	return predicate.Ban(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ban {
	return predicate.Ban(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ban {
	return predicate.Ban(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ban {
	return predicate.Ban(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ban {
	return predicate.Ban(sql.FieldLTE(FieldID, id))
}

// Check applies equality check predicate on the "check" field. It's identical to CheckEQ.
func Check(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldCheck, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldMessage, v))
}

// ValidUntil applies equality check predicate on the "validUntil" field. It's identical to ValidUntilEQ.
func ValidUntil(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldValidUntil, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldCreatedAt, v))
}

// CheckEQ applies the EQ predicate on the "check" field.
func CheckEQ(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldCheck, v))
}

// CheckNEQ applies the NEQ predicate on the "check" field.
func CheckNEQ(v string) predicate.Ban {
	return predicate.Ban(sql.FieldNEQ(FieldCheck, v))
}

// CheckIn applies the In predicate on the "check" field.
func CheckIn(vs ...string) predicate.Ban {
	return predicate.Ban(sql.FieldIn(FieldCheck, vs...))
}

// CheckNotIn applies the NotIn predicate on the "check" field.
func CheckNotIn(vs ...string) predicate.Ban {
	return predicate.Ban(sql.FieldNotIn(FieldCheck, vs...))
}

// CheckGT applies the GT predicate on the "check" field.
func CheckGT(v string) predicate.Ban {
	return predicate.Ban(sql.FieldGT(FieldCheck, v))
}

// CheckGTE applies the GTE predicate on the "check" field.
func CheckGTE(v string) predicate.Ban {
	return predicate.Ban(sql.FieldGTE(FieldCheck, v))
}

// CheckLT applies the LT predicate on the "check" field.
func CheckLT(v string) predicate.Ban {
	return predicate.Ban(sql.FieldLT(FieldCheck, v))
}

// CheckLTE applies the LTE predicate on the "check" field.
func CheckLTE(v string) predicate.Ban {
	return predicate.Ban(sql.FieldLTE(FieldCheck, v))
}

// CheckContains applies the Contains predicate on the "check" field.
func CheckContains(v string) predicate.Ban {
	return predicate.Ban(sql.FieldContains(FieldCheck, v))
}

// CheckHasPrefix applies the HasPrefix predicate on the "check" field.
func CheckHasPrefix(v string) predicate.Ban {
	return predicate.Ban(sql.FieldHasPrefix(FieldCheck, v))
}

// CheckHasSuffix applies the HasSuffix predicate on the "check" field.
func CheckHasSuffix(v string) predicate.Ban {
	return predicate.Ban(sql.FieldHasSuffix(FieldCheck, v))
}

// CheckEqualFold applies the EqualFold predicate on the "check" field.
func CheckEqualFold(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEqualFold(FieldCheck, v))
}

// CheckContainsFold applies the ContainsFold predicate on the "check" field.
func CheckContainsFold(v string) predicate.Ban {
	return predicate.Ban(sql.FieldContainsFold(FieldCheck, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Ban {
	return predicate.Ban(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Ban {
	return predicate.Ban(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Ban {
	return predicate.Ban(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Ban {
	return predicate.Ban(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Ban {
	return predicate.Ban(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Ban {
	return predicate.Ban(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Ban {
	return predicate.Ban(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Ban {
	return predicate.Ban(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Ban {
	return predicate.Ban(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Ban {
	return predicate.Ban(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Ban {
	return predicate.Ban(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Ban {
	return predicate.Ban(sql.FieldContainsFold(FieldMessage, v))
}

// ValidUntilEQ applies the EQ predicate on the "validUntil" field.
func ValidUntilEQ(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldValidUntil, v))
}

// ValidUntilNEQ applies the NEQ predicate on the "validUntil" field.
func ValidUntilNEQ(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldNEQ(FieldValidUntil, v))
}

// ValidUntilIn applies the In predicate on the "validUntil" field.
func ValidUntilIn(vs ...time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldIn(FieldValidUntil, vs...))
}

// ValidUntilNotIn applies the NotIn predicate on the "validUntil" field.
func ValidUntilNotIn(vs ...time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldNotIn(FieldValidUntil, vs...))
}

// ValidUntilGT applies the GT predicate on the "validUntil" field.
func ValidUntilGT(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldGT(FieldValidUntil, v))
}

// ValidUntilGTE applies the GTE predicate on the "validUntil" field.
func ValidUntilGTE(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldGTE(FieldValidUntil, v))
}

// ValidUntilLT applies the LT predicate on the "validUntil" field.
func ValidUntilLT(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldLT(FieldValidUntil, v))
}

// ValidUntilLTE applies the LTE predicate on the "validUntil" field.
func ValidUntilLTE(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldLTE(FieldValidUntil, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Ban {
	return predicate.Ban(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ban) predicate.Ban {
	return predicate.Ban(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ban) predicate.Ban {
	return predicate.Ban(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ban) predicate.Ban {
	return predicate.Ban(func(s *sql.Selector) {
		p(s.Not())
	})
}
