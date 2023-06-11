// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: friend.proto

package friend

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

// Validate checks the field values on FriendBaseInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FriendBaseInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlayerId

	// no validation rules for Name

	// no validation rules for IsOnline

	// no validation rules for Frame

	// no validation rules for Head

	// no validation rules for Model

	// no validation rules for Tag

	// no validation rules for Offline

	// no validation rules for FriendDegree

	// no validation rules for AddType

	// no validation rules for BaseLevel

	return nil
}

// FriendBaseInfoValidationError is the validation error returned by
// FriendBaseInfo.Validate if the designated constraints aren't met.
type FriendBaseInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FriendBaseInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FriendBaseInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FriendBaseInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FriendBaseInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FriendBaseInfoValidationError) ErrorName() string { return "FriendBaseInfoValidationError" }

// Error satisfies the builtin error interface
func (e FriendBaseInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFriendBaseInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FriendBaseInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FriendBaseInfoValidationError{}

// Validate checks the field values on RadarSearchPlayerInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RadarSearchPlayerInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Distance

	// no validation rules for PlayerId

	// no validation rules for BubbleFrame

	// no validation rules for Head

	// no validation rules for HeadFrame

	// no validation rules for NickName

	return nil
}

// RadarSearchPlayerInfoValidationError is the validation error returned by
// RadarSearchPlayerInfo.Validate if the designated constraints aren't met.
type RadarSearchPlayerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RadarSearchPlayerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RadarSearchPlayerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RadarSearchPlayerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RadarSearchPlayerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RadarSearchPlayerInfoValidationError) ErrorName() string {
	return "RadarSearchPlayerInfoValidationError"
}

// Error satisfies the builtin error interface
func (e RadarSearchPlayerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRadarSearchPlayerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RadarSearchPlayerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RadarSearchPlayerInfoValidationError{}
