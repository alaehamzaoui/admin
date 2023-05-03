// Code generated by ent, DO NOT EDIT.

package textures

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Textures {
	return predicate.Textures(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Textures {
	return predicate.Textures(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Textures {
	return predicate.Textures(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Textures {
	return predicate.Textures(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Textures {
	return predicate.Textures(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Textures {
	return predicate.Textures(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Textures {
	return predicate.Textures(sql.FieldLTE(FieldID, id))
}

// Texture applies equality check predicate on the "texture" field. It's identical to TextureEQ.
func Texture(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldTexture, v))
}

// Layer applies equality check predicate on the "layer" field. It's identical to LayerEQ.
func Layer(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldLayer, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldURL, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldCreatedAt, v))
}

// TextureEQ applies the EQ predicate on the "texture" field.
func TextureEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldTexture, v))
}

// TextureNEQ applies the NEQ predicate on the "texture" field.
func TextureNEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldNEQ(FieldTexture, v))
}

// TextureIn applies the In predicate on the "texture" field.
func TextureIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldIn(FieldTexture, vs...))
}

// TextureNotIn applies the NotIn predicate on the "texture" field.
func TextureNotIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldNotIn(FieldTexture, vs...))
}

// TextureGT applies the GT predicate on the "texture" field.
func TextureGT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGT(FieldTexture, v))
}

// TextureGTE applies the GTE predicate on the "texture" field.
func TextureGTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGTE(FieldTexture, v))
}

// TextureLT applies the LT predicate on the "texture" field.
func TextureLT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLT(FieldTexture, v))
}

// TextureLTE applies the LTE predicate on the "texture" field.
func TextureLTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLTE(FieldTexture, v))
}

// TextureContains applies the Contains predicate on the "texture" field.
func TextureContains(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContains(FieldTexture, v))
}

// TextureHasPrefix applies the HasPrefix predicate on the "texture" field.
func TextureHasPrefix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasPrefix(FieldTexture, v))
}

// TextureHasSuffix applies the HasSuffix predicate on the "texture" field.
func TextureHasSuffix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasSuffix(FieldTexture, v))
}

// TextureEqualFold applies the EqualFold predicate on the "texture" field.
func TextureEqualFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEqualFold(FieldTexture, v))
}

// TextureContainsFold applies the ContainsFold predicate on the "texture" field.
func TextureContainsFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContainsFold(FieldTexture, v))
}

// LayerEQ applies the EQ predicate on the "layer" field.
func LayerEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldLayer, v))
}

// LayerNEQ applies the NEQ predicate on the "layer" field.
func LayerNEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldNEQ(FieldLayer, v))
}

// LayerIn applies the In predicate on the "layer" field.
func LayerIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldIn(FieldLayer, vs...))
}

// LayerNotIn applies the NotIn predicate on the "layer" field.
func LayerNotIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldNotIn(FieldLayer, vs...))
}

// LayerGT applies the GT predicate on the "layer" field.
func LayerGT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGT(FieldLayer, v))
}

// LayerGTE applies the GTE predicate on the "layer" field.
func LayerGTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGTE(FieldLayer, v))
}

// LayerLT applies the LT predicate on the "layer" field.
func LayerLT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLT(FieldLayer, v))
}

// LayerLTE applies the LTE predicate on the "layer" field.
func LayerLTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLTE(FieldLayer, v))
}

// LayerContains applies the Contains predicate on the "layer" field.
func LayerContains(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContains(FieldLayer, v))
}

// LayerHasPrefix applies the HasPrefix predicate on the "layer" field.
func LayerHasPrefix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasPrefix(FieldLayer, v))
}

// LayerHasSuffix applies the HasSuffix predicate on the "layer" field.
func LayerHasSuffix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasSuffix(FieldLayer, v))
}

// LayerEqualFold applies the EqualFold predicate on the "layer" field.
func LayerEqualFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEqualFold(FieldLayer, v))
}

// LayerContainsFold applies the ContainsFold predicate on the "layer" field.
func LayerContainsFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContainsFold(FieldLayer, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Textures {
	return predicate.Textures(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Textures {
	return predicate.Textures(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Textures {
	return predicate.Textures(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Textures {
	return predicate.Textures(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Textures {
	return predicate.Textures(sql.FieldContainsFold(FieldURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Textures {
	return predicate.Textures(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Textures) predicate.Textures {
	return predicate.Textures(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Textures) predicate.Textures {
	return predicate.Textures(func(s *sql.Selector) {
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
func Not(p predicate.Textures) predicate.Textures {
	return predicate.Textures(func(s *sql.Selector) {
		p(s.Not())
	})
}
