// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/ratelimit/v4alpha/rls.proto

package envoy_config_ratelimit_v4alpha

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

	v4alpha "github.com/datawire/ambassador/v2/pkg/api/envoy/config/core/v4alpha"
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

	_ = v4alpha.ApiVersion(0)
)

// Validate checks the field values on RateLimitServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return RateLimitServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitServiceConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := v4alpha.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		return RateLimitServiceConfigValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for UseAlpha

	return nil
}

// RateLimitServiceConfigValidationError is the validation error returned by
// RateLimitServiceConfig.Validate if the designated constraints aren't met.
type RateLimitServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitServiceConfigValidationError) ErrorName() string {
	return "RateLimitServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitServiceConfigValidationError{}
