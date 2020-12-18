// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v4alpha/http_tracer.proto

package envoy_config_trace_v4alpha

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
var _http_tracer_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tracing) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TracingValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TracingValidationError is the validation error returned by Tracing.Validate
// if the designated constraints aren't met.
type TracingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TracingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TracingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TracingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TracingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TracingValidationError) ErrorName() string { return "TracingValidationError" }

// Error satisfies the builtin error interface
func (e TracingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TracingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TracingValidationError{}

// Validate checks the field values on Tracing_Http with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Tracing_Http) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return Tracing_HttpValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	switch m.ConfigType.(type) {

	case *Tracing_Http_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Tracing_HttpValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Tracing_HttpValidationError is the validation error returned by
// Tracing_Http.Validate if the designated constraints aren't met.
type Tracing_HttpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Tracing_HttpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Tracing_HttpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Tracing_HttpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Tracing_HttpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Tracing_HttpValidationError) ErrorName() string { return "Tracing_HttpValidationError" }

// Error satisfies the builtin error interface
func (e Tracing_HttpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing_Http.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Tracing_HttpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Tracing_HttpValidationError{}
