// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/meta_protocol_proxy/filters/local_ratelimit/v1alpha/local_ratelimit.proto

package aeraki_meta_protocol_proxy_filters_local_ratelimit_v1alpha

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

// Validate checks the field values on LocalRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LocalRateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		return LocalRateLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "TokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDescriptors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRateLimitValidationError{
					field:  fmt.Sprintf("Descriptors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LocalRateLimitValidationError is the validation error returned by
// LocalRateLimit.Validate if the designated constraints aren't met.
type LocalRateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRateLimitValidationError) ErrorName() string { return "LocalRateLimitValidationError" }

// Error satisfies the builtin error interface
func (e LocalRateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRateLimitValidationError{}

// Validate checks the field values on LocalRatelimitMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LocalRatelimitMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMetadata() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRatelimitMatchValidationError{
					field:  fmt.Sprintf("Metadata[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LocalRatelimitMatchValidationError is the validation error returned by
// LocalRatelimitMatch.Validate if the designated constraints aren't met.
type LocalRatelimitMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRatelimitMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRatelimitMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRatelimitMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRatelimitMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRatelimitMatchValidationError) ErrorName() string {
	return "LocalRatelimitMatchValidationError"
}

// Error satisfies the builtin error interface
func (e LocalRatelimitMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRatelimitMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRatelimitMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRatelimitMatchValidationError{}
