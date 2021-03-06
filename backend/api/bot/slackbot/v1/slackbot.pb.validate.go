// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bot/slackbot/v1/slackbot.proto

package slackbotv1

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

// define the regex for a UUID once up-front
var _slackbot_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Bot with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Bot) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Deleted

	// no validation rules for Name

	// no validation rules for Updated

	// no validation rules for AppId

	// no validation rules for Icons

	// no validation rules for TeamId

	return nil
}

// BotValidationError is the validation error returned by Bot.Validate if the
// designated constraints aren't met.
type BotValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BotValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BotValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BotValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BotValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BotValidationError) ErrorName() string { return "BotValidationError" }

// Error satisfies the builtin error interface
func (e BotValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBot.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BotValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BotValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for User

	// no validation rules for BotId

	if v, ok := interface{}(m.GetBotProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "BotProfile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Text

	// no validation rules for Ts

	// no validation rules for Channel

	// no validation rules for ChannelType

	// no validation rules for EventTs

	// no validation rules for ClientMsgId

	// no validation rules for Team

	if v, ok := interface{}(m.GetBlocks()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "Blocks",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on EventRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for IsExtSharedChannel

	// no validation rules for TeamId

	// no validation rules for ApiAppId

	if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventRequestValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	// no validation rules for EventId

	// no validation rules for EventTime

	// no validation rules for EventContext

	if v, ok := interface{}(m.GetAuthorizations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventRequestValidationError{
				field:  "Authorizations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Challenge

	// no validation rules for MinuteRateLimited

	return nil
}

// EventRequestValidationError is the validation error returned by
// EventRequest.Validate if the designated constraints aren't met.
type EventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventRequestValidationError) ErrorName() string { return "EventRequestValidationError" }

// Error satisfies the builtin error interface
func (e EventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventRequestValidationError{}

// Validate checks the field values on EventResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Challenge

	return nil
}

// EventResponseValidationError is the validation error returned by
// EventResponse.Validate if the designated constraints aren't met.
type EventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventResponseValidationError) ErrorName() string { return "EventResponseValidationError" }

// Error satisfies the builtin error interface
func (e EventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventResponseValidationError{}
