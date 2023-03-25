// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nsq.proto

package nsq

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
)

// Validate checks the field values on ComplexMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ComplexMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Cmd

	// no validation rules for Time

	return nil
}

// ComplexMessageValidationError is the validation error returned by
// ComplexMessage.Validate if the designated constraints aren't met.
type ComplexMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComplexMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComplexMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComplexMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComplexMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComplexMessageValidationError) ErrorName() string { return "ComplexMessageValidationError" }

// Error satisfies the builtin error interface
func (e ComplexMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComplexMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComplexMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComplexMessageValidationError{}

// Validate checks the field values on Mail with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Mail) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MailValidationError is the validation error returned by Mail.Validate if the
// designated constraints aren't met.
type MailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MailValidationError) ErrorName() string { return "MailValidationError" }

// Error satisfies the builtin error interface
func (e MailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MailValidationError{}

// Validate checks the field values on FriendMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FriendMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Sender

	// no validation rules for Receiver

	// no validation rules for Cmd

	return nil
}

// FriendMessageValidationError is the validation error returned by
// FriendMessage.Validate if the designated constraints aren't met.
type FriendMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FriendMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FriendMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FriendMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FriendMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FriendMessageValidationError) ErrorName() string { return "FriendMessageValidationError" }

// Error satisfies the builtin error interface
func (e FriendMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFriendMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FriendMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FriendMessageValidationError{}