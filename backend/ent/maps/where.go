// Code generated by ent, DO NOT EDIT.

package maps

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldID, id))
}

// RoomName applies equality check predicate on the "roomName" field. It's identical to RoomNameEQ.
func RoomName(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldRoomName, v))
}

// MapUrl applies equality check predicate on the "mapUrl" field. It's identical to MapUrlEQ.
func MapUrl(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldMapUrl, v))
}

// PolicyNumber applies equality check predicate on the "policyNumber" field. It's identical to PolicyNumberEQ.
func PolicyNumber(v int) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldPolicyNumber, v))
}

// ContactPage applies equality check predicate on the "contactPage" field. It's identical to ContactPageEQ.
func ContactPage(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldContactPage, v))
}

// EnableChat applies equality check predicate on the "enableChat" field. It's identical to EnableChatEQ.
func EnableChat(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChat, v))
}

// EnableChatUpload applies equality check predicate on the "enableChatUpload" field. It's identical to EnableChatUploadEQ.
func EnableChatUpload(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatUpload, v))
}

// EnableChatOnlineList applies equality check predicate on the "enableChatOnlineList" field. It's identical to EnableChatOnlineListEQ.
func EnableChatOnlineList(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatOnlineList, v))
}

// EnableChatDisconnectedList applies equality check predicate on the "enableChatDisconnectedList" field. It's identical to EnableChatDisconnectedListEQ.
func EnableChatDisconnectedList(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatDisconnectedList, v))
}

// CanReport applies equality check predicate on the "canReport" field. It's identical to CanReportEQ.
func CanReport(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldCanReport, v))
}

// ExpireOn applies equality check predicate on the "expireOn" field. It's identical to ExpireOnEQ.
func ExpireOn(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldExpireOn, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldCreatedAt, v))
}

// RoomNameEQ applies the EQ predicate on the "roomName" field.
func RoomNameEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldRoomName, v))
}

// RoomNameNEQ applies the NEQ predicate on the "roomName" field.
func RoomNameNEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldRoomName, v))
}

// RoomNameIn applies the In predicate on the "roomName" field.
func RoomNameIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldRoomName, vs...))
}

// RoomNameNotIn applies the NotIn predicate on the "roomName" field.
func RoomNameNotIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldRoomName, vs...))
}

// RoomNameGT applies the GT predicate on the "roomName" field.
func RoomNameGT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldRoomName, v))
}

// RoomNameGTE applies the GTE predicate on the "roomName" field.
func RoomNameGTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldRoomName, v))
}

// RoomNameLT applies the LT predicate on the "roomName" field.
func RoomNameLT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldRoomName, v))
}

// RoomNameLTE applies the LTE predicate on the "roomName" field.
func RoomNameLTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldRoomName, v))
}

// RoomNameContains applies the Contains predicate on the "roomName" field.
func RoomNameContains(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContains(FieldRoomName, v))
}

// RoomNameHasPrefix applies the HasPrefix predicate on the "roomName" field.
func RoomNameHasPrefix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasPrefix(FieldRoomName, v))
}

// RoomNameHasSuffix applies the HasSuffix predicate on the "roomName" field.
func RoomNameHasSuffix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasSuffix(FieldRoomName, v))
}

// RoomNameEqualFold applies the EqualFold predicate on the "roomName" field.
func RoomNameEqualFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEqualFold(FieldRoomName, v))
}

// RoomNameContainsFold applies the ContainsFold predicate on the "roomName" field.
func RoomNameContainsFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContainsFold(FieldRoomName, v))
}

// MapUrlEQ applies the EQ predicate on the "mapUrl" field.
func MapUrlEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldMapUrl, v))
}

// MapUrlNEQ applies the NEQ predicate on the "mapUrl" field.
func MapUrlNEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldMapUrl, v))
}

// MapUrlIn applies the In predicate on the "mapUrl" field.
func MapUrlIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldMapUrl, vs...))
}

// MapUrlNotIn applies the NotIn predicate on the "mapUrl" field.
func MapUrlNotIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldMapUrl, vs...))
}

// MapUrlGT applies the GT predicate on the "mapUrl" field.
func MapUrlGT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldMapUrl, v))
}

// MapUrlGTE applies the GTE predicate on the "mapUrl" field.
func MapUrlGTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldMapUrl, v))
}

// MapUrlLT applies the LT predicate on the "mapUrl" field.
func MapUrlLT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldMapUrl, v))
}

// MapUrlLTE applies the LTE predicate on the "mapUrl" field.
func MapUrlLTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldMapUrl, v))
}

// MapUrlContains applies the Contains predicate on the "mapUrl" field.
func MapUrlContains(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContains(FieldMapUrl, v))
}

// MapUrlHasPrefix applies the HasPrefix predicate on the "mapUrl" field.
func MapUrlHasPrefix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasPrefix(FieldMapUrl, v))
}

// MapUrlHasSuffix applies the HasSuffix predicate on the "mapUrl" field.
func MapUrlHasSuffix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasSuffix(FieldMapUrl, v))
}

// MapUrlEqualFold applies the EqualFold predicate on the "mapUrl" field.
func MapUrlEqualFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEqualFold(FieldMapUrl, v))
}

// MapUrlContainsFold applies the ContainsFold predicate on the "mapUrl" field.
func MapUrlContainsFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContainsFold(FieldMapUrl, v))
}

// PolicyNumberEQ applies the EQ predicate on the "policyNumber" field.
func PolicyNumberEQ(v int) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldPolicyNumber, v))
}

// PolicyNumberNEQ applies the NEQ predicate on the "policyNumber" field.
func PolicyNumberNEQ(v int) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldPolicyNumber, v))
}

// PolicyNumberIn applies the In predicate on the "policyNumber" field.
func PolicyNumberIn(vs ...int) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldPolicyNumber, vs...))
}

// PolicyNumberNotIn applies the NotIn predicate on the "policyNumber" field.
func PolicyNumberNotIn(vs ...int) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldPolicyNumber, vs...))
}

// PolicyNumberGT applies the GT predicate on the "policyNumber" field.
func PolicyNumberGT(v int) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldPolicyNumber, v))
}

// PolicyNumberGTE applies the GTE predicate on the "policyNumber" field.
func PolicyNumberGTE(v int) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldPolicyNumber, v))
}

// PolicyNumberLT applies the LT predicate on the "policyNumber" field.
func PolicyNumberLT(v int) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldPolicyNumber, v))
}

// PolicyNumberLTE applies the LTE predicate on the "policyNumber" field.
func PolicyNumberLTE(v int) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldPolicyNumber, v))
}

// ContactPageEQ applies the EQ predicate on the "contactPage" field.
func ContactPageEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldContactPage, v))
}

// ContactPageNEQ applies the NEQ predicate on the "contactPage" field.
func ContactPageNEQ(v string) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldContactPage, v))
}

// ContactPageIn applies the In predicate on the "contactPage" field.
func ContactPageIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldContactPage, vs...))
}

// ContactPageNotIn applies the NotIn predicate on the "contactPage" field.
func ContactPageNotIn(vs ...string) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldContactPage, vs...))
}

// ContactPageGT applies the GT predicate on the "contactPage" field.
func ContactPageGT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldContactPage, v))
}

// ContactPageGTE applies the GTE predicate on the "contactPage" field.
func ContactPageGTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldContactPage, v))
}

// ContactPageLT applies the LT predicate on the "contactPage" field.
func ContactPageLT(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldContactPage, v))
}

// ContactPageLTE applies the LTE predicate on the "contactPage" field.
func ContactPageLTE(v string) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldContactPage, v))
}

// ContactPageContains applies the Contains predicate on the "contactPage" field.
func ContactPageContains(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContains(FieldContactPage, v))
}

// ContactPageHasPrefix applies the HasPrefix predicate on the "contactPage" field.
func ContactPageHasPrefix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasPrefix(FieldContactPage, v))
}

// ContactPageHasSuffix applies the HasSuffix predicate on the "contactPage" field.
func ContactPageHasSuffix(v string) predicate.Maps {
	return predicate.Maps(sql.FieldHasSuffix(FieldContactPage, v))
}

// ContactPageEqualFold applies the EqualFold predicate on the "contactPage" field.
func ContactPageEqualFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldEqualFold(FieldContactPage, v))
}

// ContactPageContainsFold applies the ContainsFold predicate on the "contactPage" field.
func ContactPageContainsFold(v string) predicate.Maps {
	return predicate.Maps(sql.FieldContainsFold(FieldContactPage, v))
}

// EnableChatEQ applies the EQ predicate on the "enableChat" field.
func EnableChatEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChat, v))
}

// EnableChatNEQ applies the NEQ predicate on the "enableChat" field.
func EnableChatNEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldEnableChat, v))
}

// EnableChatUploadEQ applies the EQ predicate on the "enableChatUpload" field.
func EnableChatUploadEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatUpload, v))
}

// EnableChatUploadNEQ applies the NEQ predicate on the "enableChatUpload" field.
func EnableChatUploadNEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldEnableChatUpload, v))
}

// EnableChatOnlineListEQ applies the EQ predicate on the "enableChatOnlineList" field.
func EnableChatOnlineListEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatOnlineList, v))
}

// EnableChatOnlineListNEQ applies the NEQ predicate on the "enableChatOnlineList" field.
func EnableChatOnlineListNEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldEnableChatOnlineList, v))
}

// EnableChatDisconnectedListEQ applies the EQ predicate on the "enableChatDisconnectedList" field.
func EnableChatDisconnectedListEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldEnableChatDisconnectedList, v))
}

// EnableChatDisconnectedListNEQ applies the NEQ predicate on the "enableChatDisconnectedList" field.
func EnableChatDisconnectedListNEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldEnableChatDisconnectedList, v))
}

// CanReportEQ applies the EQ predicate on the "canReport" field.
func CanReportEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldCanReport, v))
}

// CanReportNEQ applies the NEQ predicate on the "canReport" field.
func CanReportNEQ(v bool) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldCanReport, v))
}

// ExpireOnEQ applies the EQ predicate on the "expireOn" field.
func ExpireOnEQ(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldExpireOn, v))
}

// ExpireOnNEQ applies the NEQ predicate on the "expireOn" field.
func ExpireOnNEQ(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldExpireOn, v))
}

// ExpireOnIn applies the In predicate on the "expireOn" field.
func ExpireOnIn(vs ...time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldExpireOn, vs...))
}

// ExpireOnNotIn applies the NotIn predicate on the "expireOn" field.
func ExpireOnNotIn(vs ...time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldExpireOn, vs...))
}

// ExpireOnGT applies the GT predicate on the "expireOn" field.
func ExpireOnGT(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldExpireOn, v))
}

// ExpireOnGTE applies the GTE predicate on the "expireOn" field.
func ExpireOnGTE(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldExpireOn, v))
}

// ExpireOnLT applies the LT predicate on the "expireOn" field.
func ExpireOnLT(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldExpireOn, v))
}

// ExpireOnLTE applies the LTE predicate on the "expireOn" field.
func ExpireOnLTE(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldExpireOn, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Maps {
	return predicate.Maps(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Maps) predicate.Maps {
	return predicate.Maps(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Maps) predicate.Maps {
	return predicate.Maps(func(s *sql.Selector) {
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
func Not(p predicate.Maps) predicate.Maps {
	return predicate.Maps(func(s *sql.Selector) {
		p(s.Not())
	})
}
