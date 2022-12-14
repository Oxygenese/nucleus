// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/base"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/node"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	baseFields := schema.Base{}.Fields()
	_ = baseFields
	// baseDescCreatedAt is the schema descriptor for created_at field.
	baseDescCreatedAt := baseFields[1].Descriptor()
	// base.DefaultCreatedAt holds the default value on creation for the created_at field.
	base.DefaultCreatedAt = baseDescCreatedAt.Default.(func() time.Time)
	// baseDescUpdatedAt is the schema descriptor for updated_at field.
	baseDescUpdatedAt := baseFields[2].Descriptor()
	// base.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	base.UpdateDefaultUpdatedAt = baseDescUpdatedAt.UpdateDefault.(func() time.Time)
	// baseDescID is the schema descriptor for id field.
	baseDescID := baseFields[0].Descriptor()
	// base.IDValidator is a validator for the "id" field. It is called by the builders before save.
	base.IDValidator = baseDescID.Validators[0].(func(int) error)
	nodeMixin := schema.Node{}.Mixin()
	nodeMixinFields0 := nodeMixin[0].Fields()
	_ = nodeMixinFields0
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescCreatedAt is the schema descriptor for created_at field.
	nodeDescCreatedAt := nodeMixinFields0[1].Descriptor()
	// node.DefaultCreatedAt holds the default value on creation for the created_at field.
	node.DefaultCreatedAt = nodeDescCreatedAt.Default.(func() time.Time)
	// nodeDescUpdatedAt is the schema descriptor for updated_at field.
	nodeDescUpdatedAt := nodeMixinFields0[2].Descriptor()
	// node.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	node.UpdateDefaultUpdatedAt = nodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// nodeDescID is the schema descriptor for id field.
	nodeDescID := nodeFields[0].Descriptor()
	// node.IDValidator is a validator for the "id" field. It is called by the builders before save.
	node.IDValidator = nodeDescID.Validators[0].(func(int) error)
}
