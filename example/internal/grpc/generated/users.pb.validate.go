// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: users.proto

package generated

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetByIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetByIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetByIDRequestMultiError,
// or nil if none found.
func (m *GetByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetByIDRequestMultiError(errors)
	}

	return nil
}

// GetByIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetByIDRequest.ValidateAll() if the designated constraints
// aren't met.
type GetByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetByIDRequestMultiError) AllErrors() []error { return m }

// GetByIDRequestValidationError is the validation error returned by
// GetByIDRequest.Validate if the designated constraints aren't met.
type GetByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIDRequestValidationError) ErrorName() string { return "GetByIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIDRequestValidationError{}

// Validate checks the field values on UserList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserListMultiError, or nil
// if none found.
func (m *UserList) ValidateAll() error {
	return m.validate(true)
}

func (m *UserList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserListMultiError(errors)
	}

	return nil
}

// UserListMultiError is an error wrapping multiple validation errors returned
// by UserList.ValidateAll() if the designated constraints aren't met.
type UserListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListMultiError) AllErrors() []error { return m }

// UserListValidationError is the validation error returned by
// UserList.Validate if the designated constraints aren't met.
type UserListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListValidationError) ErrorName() string { return "UserListValidationError" }

// Error satisfies the builtin error interface
func (e UserListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetUsername()) < 11 {
		err := UserValidationError{
			field:  "Username",
			reason: "value length must be at least 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Email

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}
