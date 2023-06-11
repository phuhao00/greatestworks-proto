// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat.proto

package chat

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

// Validate checks the field values on PrivateChatMsg with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PrivateChatMsg) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	// no validation rules for SendTime

	if v, ok := interface{}(m.GetSender()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PrivateChatMsgValidationError{
				field:  "Sender",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReceiver()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PrivateChatMsgValidationError{
				field:  "Receiver",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsVoice

	// no validation rules for Degree

	return nil
}

// PrivateChatMsgValidationError is the validation error returned by
// PrivateChatMsg.Validate if the designated constraints aren't met.
type PrivateChatMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrivateChatMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrivateChatMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrivateChatMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrivateChatMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrivateChatMsgValidationError) ErrorName() string { return "PrivateChatMsgValidationError" }

// Error satisfies the builtin error interface
func (e PrivateChatMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrivateChatMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrivateChatMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrivateChatMsgValidationError{}

// Validate checks the field values on ChatBaseInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ChatBaseInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BubbleFrame

	// no validation rules for Head

	// no validation rules for HeadFrame

	// no validation rules for NickName

	// no validation rules for Id

	return nil
}

// ChatBaseInfoValidationError is the validation error returned by
// ChatBaseInfo.Validate if the designated constraints aren't met.
type ChatBaseInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatBaseInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatBaseInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatBaseInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatBaseInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatBaseInfoValidationError) ErrorName() string { return "ChatBaseInfoValidationError" }

// Error satisfies the builtin error interface
func (e ChatBaseInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatBaseInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatBaseInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatBaseInfoValidationError{}

// Validate checks the field values on ChatStatistics with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ChatStatistics) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UseNum

	// no validation rules for TimeNum

	// no validation rules for Speaker

	// no validation rules for LastTime

	// no validation rules for PriNum

	return nil
}

// ChatStatisticsValidationError is the validation error returned by
// ChatStatistics.Validate if the designated constraints aren't met.
type ChatStatisticsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatStatisticsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatStatisticsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatStatisticsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatStatisticsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatStatisticsValidationError) ErrorName() string { return "ChatStatisticsValidationError" }

// Error satisfies the builtin error interface
func (e ChatStatisticsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatStatistics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatStatisticsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatStatisticsValidationError{}