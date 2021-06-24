// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-role-api/ocp-role-api.proto

package ocp_role_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ListRolesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRolesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListRolesV1RequestValidationError is the validation error returned by
// ListRolesV1Request.Validate if the designated constraints aren't met.
type ListRolesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesV1RequestValidationError) ErrorName() string {
	return "ListRolesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesV1RequestValidationError{}

// Validate checks the field values on ListRolesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRolesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRolesV1ResponseValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRolesV1ResponseValidationError is the validation error returned by
// ListRolesV1Response.Validate if the designated constraints aren't met.
type ListRolesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesV1ResponseValidationError) ErrorName() string {
	return "ListRolesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesV1ResponseValidationError{}

// Validate checks the field values on CreateRoleV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRoleV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	// no validation rules for Operation

	return nil
}

// CreateRoleV1RequestValidationError is the validation error returned by
// CreateRoleV1Request.Validate if the designated constraints aren't met.
type CreateRoleV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleV1RequestValidationError) ErrorName() string {
	return "CreateRoleV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleV1RequestValidationError{}

// Validate checks the field values on CreateRoleV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRoleV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoleId

	return nil
}

// CreateRoleV1ResponseValidationError is the validation error returned by
// CreateRoleV1Response.Validate if the designated constraints aren't met.
type CreateRoleV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleV1ResponseValidationError) ErrorName() string {
	return "CreateRoleV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleV1ResponseValidationError{}

// Validate checks the field values on MultiCreateRoleV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRoleV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateRoleV1RequestValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateRoleV1RequestValidationError is the validation error returned by
// MultiCreateRoleV1Request.Validate if the designated constraints aren't met.
type MultiCreateRoleV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRoleV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRoleV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRoleV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRoleV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRoleV1RequestValidationError) ErrorName() string {
	return "MultiCreateRoleV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRoleV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRoleV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRoleV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRoleV1RequestValidationError{}

// Validate checks the field values on MultiCreateRoleV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRoleV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiCreateRoleV1ResponseValidationError is the validation error returned by
// MultiCreateRoleV1Response.Validate if the designated constraints aren't met.
type MultiCreateRoleV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRoleV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRoleV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRoleV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRoleV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRoleV1ResponseValidationError) ErrorName() string {
	return "MultiCreateRoleV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRoleV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRoleV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRoleV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRoleV1ResponseValidationError{}

// Validate checks the field values on UpdateRoleV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRoleV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoleId

	// no validation rules for Service

	// no validation rules for Operation

	return nil
}

// UpdateRoleV1RequestValidationError is the validation error returned by
// UpdateRoleV1Request.Validate if the designated constraints aren't met.
type UpdateRoleV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleV1RequestValidationError) ErrorName() string {
	return "UpdateRoleV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleV1RequestValidationError{}

// Validate checks the field values on UpdateRoleV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRoleV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// UpdateRoleV1ResponseValidationError is the validation error returned by
// UpdateRoleV1Response.Validate if the designated constraints aren't met.
type UpdateRoleV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleV1ResponseValidationError) ErrorName() string {
	return "UpdateRoleV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleV1ResponseValidationError{}

// Validate checks the field values on RemoveRoleV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRoleV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoleId

	return nil
}

// RemoveRoleV1RequestValidationError is the validation error returned by
// RemoveRoleV1Request.Validate if the designated constraints aren't met.
type RemoveRoleV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRoleV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRoleV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRoleV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRoleV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRoleV1RequestValidationError) ErrorName() string {
	return "RemoveRoleV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRoleV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRoleV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRoleV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRoleV1RequestValidationError{}

// Validate checks the field values on RemoveRoleV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRoleV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveRoleV1ResponseValidationError is the validation error returned by
// RemoveRoleV1Response.Validate if the designated constraints aren't met.
type RemoveRoleV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRoleV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRoleV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRoleV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRoleV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRoleV1ResponseValidationError) ErrorName() string {
	return "RemoveRoleV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRoleV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRoleV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRoleV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRoleV1ResponseValidationError{}

// Validate checks the field values on DescribeRoleV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRoleV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoleId

	return nil
}

// DescribeRoleV1RequestValidationError is the validation error returned by
// DescribeRoleV1Request.Validate if the designated constraints aren't met.
type DescribeRoleV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRoleV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRoleV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRoleV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRoleV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRoleV1RequestValidationError) ErrorName() string {
	return "DescribeRoleV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRoleV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRoleV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRoleV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRoleV1RequestValidationError{}

// Validate checks the field values on DescribeRoleV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRoleV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeRoleV1ResponseValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeRoleV1ResponseValidationError is the validation error returned by
// DescribeRoleV1Response.Validate if the designated constraints aren't met.
type DescribeRoleV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRoleV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRoleV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRoleV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRoleV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRoleV1ResponseValidationError) ErrorName() string {
	return "DescribeRoleV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRoleV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRoleV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRoleV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRoleV1ResponseValidationError{}

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Service

	// no validation rules for Operation

	return nil
}

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on MultiCreateRoleV1Request_Role with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRoleV1Request_Role) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	// no validation rules for Operation

	return nil
}

// MultiCreateRoleV1Request_RoleValidationError is the validation error
// returned by MultiCreateRoleV1Request_Role.Validate if the designated
// constraints aren't met.
type MultiCreateRoleV1Request_RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRoleV1Request_RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRoleV1Request_RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRoleV1Request_RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRoleV1Request_RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRoleV1Request_RoleValidationError) ErrorName() string {
	return "MultiCreateRoleV1Request_RoleValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRoleV1Request_RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRoleV1Request_Role.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRoleV1Request_RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRoleV1Request_RoleValidationError{}
