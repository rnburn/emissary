// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/gzip/v3/gzip.proto

package envoy_extensions_filters_http_gzip_v3

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
var _gzip_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Gzip with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Gzip) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMemoryLevel(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 9 {
			return GzipValidationError{
				field:  "MemoryLevel",
				reason: "value must be inside range [1, 9]",
			}
		}

	}

	if _, ok := Gzip_CompressionLevel_Enum_name[int32(m.GetCompressionLevel())]; !ok {
		return GzipValidationError{
			field:  "CompressionLevel",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := Gzip_CompressionStrategy_name[int32(m.GetCompressionStrategy())]; !ok {
		return GzipValidationError{
			field:  "CompressionStrategy",
			reason: "value must be one of the defined enum values",
		}
	}

	if wrapper := m.GetWindowBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 9 || val > 15 {
			return GzipValidationError{
				field:  "WindowBits",
				reason: "value must be inside range [9, 15]",
			}
		}

	}

	if v, ok := interface{}(m.GetCompressor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GzipValidationError{
				field:  "Compressor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedContentLength()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GzipValidationError{
				field:  "HiddenEnvoyDeprecatedContentLength",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HiddenEnvoyDeprecatedDisableOnEtagHeader

	// no validation rules for HiddenEnvoyDeprecatedRemoveAcceptEncodingHeader

	return nil
}

// GzipValidationError is the validation error returned by Gzip.Validate if the
// designated constraints aren't met.
type GzipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GzipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GzipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GzipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GzipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GzipValidationError) ErrorName() string { return "GzipValidationError" }

// Error satisfies the builtin error interface
func (e GzipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGzip.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GzipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GzipValidationError{}

// Validate checks the field values on Gzip_CompressionLevel with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Gzip_CompressionLevel) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// Gzip_CompressionLevelValidationError is the validation error returned by
// Gzip_CompressionLevel.Validate if the designated constraints aren't met.
type Gzip_CompressionLevelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Gzip_CompressionLevelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Gzip_CompressionLevelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Gzip_CompressionLevelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Gzip_CompressionLevelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Gzip_CompressionLevelValidationError) ErrorName() string {
	return "Gzip_CompressionLevelValidationError"
}

// Error satisfies the builtin error interface
func (e Gzip_CompressionLevelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGzip_CompressionLevel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Gzip_CompressionLevelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Gzip_CompressionLevelValidationError{}