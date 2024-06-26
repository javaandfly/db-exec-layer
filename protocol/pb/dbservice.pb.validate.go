// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protocol/pb/dbservice_protofile/dbservice.proto

package pb

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

// Validate checks the field values on DB_PING_REQ with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DB_PING_REQ) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DB_PING_REQ with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DB_PING_REQMultiError, or
// nil if none found.
func (m *DB_PING_REQ) ValidateAll() error {
	return m.validate(true)
}

func (m *DB_PING_REQ) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Timestamp

	if len(errors) > 0 {
		return DB_PING_REQMultiError(errors)
	}

	return nil
}

// DB_PING_REQMultiError is an error wrapping multiple validation errors
// returned by DB_PING_REQ.ValidateAll() if the designated constraints aren't met.
type DB_PING_REQMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DB_PING_REQMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DB_PING_REQMultiError) AllErrors() []error { return m }

// DB_PING_REQValidationError is the validation error returned by
// DB_PING_REQ.Validate if the designated constraints aren't met.
type DB_PING_REQValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DB_PING_REQValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DB_PING_REQValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DB_PING_REQValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DB_PING_REQValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DB_PING_REQValidationError) ErrorName() string { return "DB_PING_REQValidationError" }

// Error satisfies the builtin error interface
func (e DB_PING_REQValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDB_PING_REQ.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DB_PING_REQValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DB_PING_REQValidationError{}

// Validate checks the field values on DB_PING_RESP with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DB_PING_RESP) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DB_PING_RESP with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DB_PING_RESPMultiError, or
// nil if none found.
func (m *DB_PING_RESP) ValidateAll() error {
	return m.validate(true)
}

func (m *DB_PING_RESP) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TimeStamp

	if len(errors) > 0 {
		return DB_PING_RESPMultiError(errors)
	}

	return nil
}

// DB_PING_RESPMultiError is an error wrapping multiple validation errors
// returned by DB_PING_RESP.ValidateAll() if the designated constraints aren't met.
type DB_PING_RESPMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DB_PING_RESPMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DB_PING_RESPMultiError) AllErrors() []error { return m }

// DB_PING_RESPValidationError is the validation error returned by
// DB_PING_RESP.Validate if the designated constraints aren't met.
type DB_PING_RESPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DB_PING_RESPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DB_PING_RESPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DB_PING_RESPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DB_PING_RESPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DB_PING_RESPValidationError) ErrorName() string { return "DB_PING_RESPValidationError" }

// Error satisfies the builtin error interface
func (e DB_PING_RESPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDB_PING_RESP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DB_PING_RESPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DB_PING_RESPValidationError{}

// Validate checks the field values on DB_HEART_BEAT_REQ with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DB_HEART_BEAT_REQ) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DB_HEART_BEAT_REQ with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DB_HEART_BEAT_REQMultiError, or nil if none found.
func (m *DB_HEART_BEAT_REQ) ValidateAll() error {
	return m.validate(true)
}

func (m *DB_HEART_BEAT_REQ) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Timestamp

	// no validation rules for Tick

	if len(errors) > 0 {
		return DB_HEART_BEAT_REQMultiError(errors)
	}

	return nil
}

// DB_HEART_BEAT_REQMultiError is an error wrapping multiple validation errors
// returned by DB_HEART_BEAT_REQ.ValidateAll() if the designated constraints
// aren't met.
type DB_HEART_BEAT_REQMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DB_HEART_BEAT_REQMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DB_HEART_BEAT_REQMultiError) AllErrors() []error { return m }

// DB_HEART_BEAT_REQValidationError is the validation error returned by
// DB_HEART_BEAT_REQ.Validate if the designated constraints aren't met.
type DB_HEART_BEAT_REQValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DB_HEART_BEAT_REQValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DB_HEART_BEAT_REQValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DB_HEART_BEAT_REQValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DB_HEART_BEAT_REQValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DB_HEART_BEAT_REQValidationError) ErrorName() string {
	return "DB_HEART_BEAT_REQValidationError"
}

// Error satisfies the builtin error interface
func (e DB_HEART_BEAT_REQValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDB_HEART_BEAT_REQ.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DB_HEART_BEAT_REQValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DB_HEART_BEAT_REQValidationError{}

// Validate checks the field values on DB_HEART_BEAT_RESP with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DB_HEART_BEAT_RESP) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DB_HEART_BEAT_RESP with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DB_HEART_BEAT_RESPMultiError, or nil if none found.
func (m *DB_HEART_BEAT_RESP) ValidateAll() error {
	return m.validate(true)
}

func (m *DB_HEART_BEAT_RESP) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Timestamp

	// no validation rules for Ticl

	if len(errors) > 0 {
		return DB_HEART_BEAT_RESPMultiError(errors)
	}

	return nil
}

// DB_HEART_BEAT_RESPMultiError is an error wrapping multiple validation errors
// returned by DB_HEART_BEAT_RESP.ValidateAll() if the designated constraints
// aren't met.
type DB_HEART_BEAT_RESPMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DB_HEART_BEAT_RESPMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DB_HEART_BEAT_RESPMultiError) AllErrors() []error { return m }

// DB_HEART_BEAT_RESPValidationError is the validation error returned by
// DB_HEART_BEAT_RESP.Validate if the designated constraints aren't met.
type DB_HEART_BEAT_RESPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DB_HEART_BEAT_RESPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DB_HEART_BEAT_RESPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DB_HEART_BEAT_RESPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DB_HEART_BEAT_RESPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DB_HEART_BEAT_RESPValidationError) ErrorName() string {
	return "DB_HEART_BEAT_RESPValidationError"
}

// Error satisfies the builtin error interface
func (e DB_HEART_BEAT_RESPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDB_HEART_BEAT_RESP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DB_HEART_BEAT_RESPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DB_HEART_BEAT_RESPValidationError{}
