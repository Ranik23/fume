// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fume/internal/storage/ent/request"
	"fume/internal/storage/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescEndpoint is the schema descriptor for endpoint field.
	requestDescEndpoint := requestFields[1].Descriptor()
	// request.EndpointValidator is a validator for the "endpoint" field. It is called by the builders before save.
	request.EndpointValidator = func() func(string) error {
		validators := requestDescEndpoint.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(endpoint string) error {
			for _, fn := range fns {
				if err := fn(endpoint); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// requestDescMethod is the schema descriptor for method field.
	requestDescMethod := requestFields[2].Descriptor()
	// request.MethodValidator is a validator for the "method" field. It is called by the builders before save.
	request.MethodValidator = func() func(string) error {
		validators := requestDescMethod.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(method string) error {
			for _, fn := range fns {
				if err := fn(method); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// requestDescCreatedAt is the schema descriptor for created_at field.
	requestDescCreatedAt := requestFields[3].Descriptor()
	// request.DefaultCreatedAt holds the default value on creation for the created_at field.
	request.DefaultCreatedAt = requestDescCreatedAt.Default.(func() time.Time)
	// requestDescUpdatedAt is the schema descriptor for updated_at field.
	requestDescUpdatedAt := requestFields[4].Descriptor()
	// request.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	request.DefaultUpdatedAt = requestDescUpdatedAt.Default.(func() time.Time)
	// request.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	request.UpdateDefaultUpdatedAt = requestDescUpdatedAt.UpdateDefault.(func() time.Time)
	// requestDescID is the schema descriptor for id field.
	requestDescID := requestFields[0].Descriptor()
	// request.IDValidator is a validator for the "id" field. It is called by the builders before save.
	request.IDValidator = requestDescID.Validators[0].(func(int) error)
}
