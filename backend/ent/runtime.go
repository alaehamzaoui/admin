// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/bo-mathadventure/admin/ent/announcement"
	"github.com/bo-mathadventure/admin/ent/ban"
	"github.com/bo-mathadventure/admin/ent/maps"
	"github.com/bo-mathadventure/admin/ent/report"
	"github.com/bo-mathadventure/admin/ent/schema"
	"github.com/bo-mathadventure/admin/ent/textures"
	"github.com/bo-mathadventure/admin/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	announcementFields := schema.Announcement{}.Fields()
	_ = announcementFields
	// announcementDescType is the schema descriptor for type field.
	announcementDescType := announcementFields[0].Descriptor()
	// announcement.DefaultType holds the default value on creation for the type field.
	announcement.DefaultType = announcementDescType.Default.(string)
	// announcementDescCreatedAt is the schema descriptor for createdAt field.
	announcementDescCreatedAt := announcementFields[2].Descriptor()
	// announcement.DefaultCreatedAt holds the default value on creation for the createdAt field.
	announcement.DefaultCreatedAt = announcementDescCreatedAt.Default.(func() time.Time)
	banFields := schema.Ban{}.Fields()
	_ = banFields
	// banDescCreatedAt is the schema descriptor for createdAt field.
	banDescCreatedAt := banFields[3].Descriptor()
	// ban.DefaultCreatedAt holds the default value on creation for the createdAt field.
	ban.DefaultCreatedAt = banDescCreatedAt.Default.(func() time.Time)
	mapsFields := schema.Maps{}.Fields()
	_ = mapsFields
	// mapsDescTags is the schema descriptor for tags field.
	mapsDescTags := mapsFields[4].Descriptor()
	// maps.DefaultTags holds the default value on creation for the tags field.
	maps.DefaultTags = mapsDescTags.Default.([]string)
	// mapsDescCreatedAt is the schema descriptor for createdAt field.
	mapsDescCreatedAt := mapsFields[11].Descriptor()
	// maps.DefaultCreatedAt holds the default value on creation for the createdAt field.
	maps.DefaultCreatedAt = mapsDescCreatedAt.Default.(func() time.Time)
	reportFields := schema.Report{}.Fields()
	_ = reportFields
	// reportDescHide is the schema descriptor for hide field.
	reportDescHide := reportFields[2].Descriptor()
	// report.DefaultHide holds the default value on creation for the hide field.
	report.DefaultHide = reportDescHide.Default.(bool)
	// reportDescCreatedAt is the schema descriptor for createdAt field.
	reportDescCreatedAt := reportFields[3].Descriptor()
	// report.DefaultCreatedAt holds the default value on creation for the createdAt field.
	report.DefaultCreatedAt = reportDescCreatedAt.Default.(func() time.Time)
	texturesFields := schema.Textures{}.Fields()
	_ = texturesFields
	// texturesDescTags is the schema descriptor for tags field.
	texturesDescTags := texturesFields[3].Descriptor()
	// textures.DefaultTags holds the default value on creation for the tags field.
	textures.DefaultTags = texturesDescTags.Default.([]string)
	// texturesDescCreatedAt is the schema descriptor for createdAt field.
	texturesDescCreatedAt := texturesFields[4].Descriptor()
	// textures.DefaultCreatedAt holds the default value on creation for the createdAt field.
	textures.DefaultCreatedAt = texturesDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPermissions is the schema descriptor for permissions field.
	userDescPermissions := userFields[6].Descriptor()
	// user.DefaultPermissions holds the default value on creation for the permissions field.
	user.DefaultPermissions = userDescPermissions.Default.([]string)
	// userDescTags is the schema descriptor for tags field.
	userDescTags := userFields[7].Descriptor()
	// user.DefaultTags holds the default value on creation for the tags field.
	user.DefaultTags = userDescTags.Default.([]string)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[8].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
