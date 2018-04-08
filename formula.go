package brewer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
)

const (
	tagRegexp = `(?i):tag => "(.*)"`
	revRegexp = `(?i):revision => "(.*)"`
	urlRegexp = `(?i)url "(.*)"`
	shaRegexp = `(?i)sha256 "(.*)"`
)

// Formula represents a Homebrew formula that can be updated.
type Formula struct {
	Path     string
	Contents string
}

// ContentsSHA returns the SHA256 hash of the formula contents.
func (f *Formula) ContentsSHA() string {
	hash := sha256.New()
	hash.Write([]byte(f.Contents))

	return hex.EncodeToString(hash.Sum(nil))
}

// Tag returns the value for `:tag => "{{tag}}"` if the formula has one.
func (f *Formula) Tag() string {
	r, _ := regexp.Compile(tagRegexp)
	matches := r.FindStringSubmatch(f.Contents)

	if matches == nil {
		return ""
	}

	return matches[len(matches)-1]
}

// UpdateTag updates the `:tag => "{{tag}}"` part of a formula.
func (f *Formula) UpdateTag(tag string) error {
	r, err := regexp.Compile(tagRegexp)
	if err != nil {
		return err
	}

	f.Contents = r.ReplaceAllString(f.Contents, fmt.Sprintf(`:tag => "%s"`, tag))

	return nil
}

// Revision returns the value for `:revision => "{{revision}}"` if the formula
// has one.
func (f *Formula) Revision() string {
	// TODO
	return ""
}

// UpdateRevision updates the `:revision => "{{revision}}"` part of a formula.
func (f *Formula) UpdateRevision(rev string) error {
	// TODO
	return nil
}

// URL returns the value for `url "{{url}}"` if the formula has one.
func (f *Formula) URL() string {
	// TODO
	return ""
}

// UpdateURL updates the `url "{{url}}"` part of a formula.
func (f *Formula) UpdateURL(url string) error {
	// TODO
	return nil
}

// SHA256 returns the value for the `sha256 "{{sha}}"` if the formula has one.
func (f *Formula) SHA256() string {
	// TODO
	return ""
}

// UpdateSHA256 updates the `sha256 "{{sha}}"` part of a formula.
func (f *Formula) UpdateSHA256(sha string) error {
	// TODO
	return nil
}
