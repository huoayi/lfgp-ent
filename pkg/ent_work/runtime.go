// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"time"

	"github.com/huoayi/lfgp-ent/pkg/ent_work/creation"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/schema"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	creationMixin := schema.Creation{}.Mixin()
	creationMixinFields0 := creationMixin[0].Fields()
	_ = creationMixinFields0
	creationFields := schema.Creation{}.Fields()
	_ = creationFields
	// creationDescCreatedBy is the schema descriptor for created_by field.
	creationDescCreatedBy := creationMixinFields0[1].Descriptor()
	// creation.DefaultCreatedBy holds the default value on creation for the created_by field.
	creation.DefaultCreatedBy = creationDescCreatedBy.Default.(int64)
	// creationDescUpdatedBy is the schema descriptor for updated_by field.
	creationDescUpdatedBy := creationMixinFields0[2].Descriptor()
	// creation.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	creation.DefaultUpdatedBy = creationDescUpdatedBy.Default.(int64)
	// creationDescCreatedAt is the schema descriptor for created_at field.
	creationDescCreatedAt := creationMixinFields0[3].Descriptor()
	// creation.DefaultCreatedAt holds the default value on creation for the created_at field.
	creation.DefaultCreatedAt = creationDescCreatedAt.Default.(func() time.Time)
	// creationDescUpdatedAt is the schema descriptor for updated_at field.
	creationDescUpdatedAt := creationMixinFields0[4].Descriptor()
	// creation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	creation.DefaultUpdatedAt = creationDescUpdatedAt.Default.(func() time.Time)
	// creation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	creation.UpdateDefaultUpdatedAt = creationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// creationDescDeletedAt is the schema descriptor for deleted_at field.
	creationDescDeletedAt := creationMixinFields0[5].Descriptor()
	// creation.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	creation.DefaultDeletedAt = creationDescDeletedAt.Default.(time.Time)
	// creationDescParameter is the schema descriptor for parameter field.
	creationDescParameter := creationFields[1].Descriptor()
	// creation.DefaultParameter holds the default value on creation for the parameter field.
	creation.DefaultParameter = creationDescParameter.Default.(string)
	// creationDescURL is the schema descriptor for url field.
	creationDescURL := creationFields[2].Descriptor()
	// creation.DefaultURL holds the default value on creation for the url field.
	creation.DefaultURL = creationDescURL.Default.(string)
	// creationDescUserID is the schema descriptor for user_id field.
	creationDescUserID := creationFields[3].Descriptor()
	// creation.DefaultUserID holds the default value on creation for the user_id field.
	creation.DefaultUserID = creationDescUserID.Default.(int64)
	// creationDescID is the schema descriptor for id field.
	creationDescID := creationMixinFields0[0].Descriptor()
	// creation.DefaultID holds the default value on creation for the id field.
	creation.DefaultID = creationDescID.Default.(func() int64)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(int64)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(int64)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[5].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[1].Descriptor()
	// user.DefaultNickName holds the default value on creation for the nick_name field.
	user.DefaultNickName = userDescNickName.Default.(string)
	// userDescJpgURL is the schema descriptor for jpg_url field.
	userDescJpgURL := userFields[2].Descriptor()
	// user.DefaultJpgURL holds the default value on creation for the jpg_url field.
	user.DefaultJpgURL = userDescJpgURL.Default.(string)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[4].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
}
