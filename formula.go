package brewer

import (
	"crypto/sha256"
	"encoding/hex"
)

// Formula represents a Homebrew formula that can be updated.
type Formula struct {
	Path     string
	Contents string
}

// SHA returns the SHA256 hash of the formula contents.
func (f *Formula) SHA() string {
	hash := sha256.New()
	hash.Write([]byte(f.Contents))

	return hex.EncodeToString(hash.Sum(nil))
}

// UpdateTag updates the `:tag => "{{tag}}"` part of a formula.
func (f *Formula) UpdateTag(tag string) error {
	return nil
}

// UpdateRevision updates the `:revision => "{{revision}}"` part of a formula.
func (f *Formula) UpdateRevision(rev string) error {
	return nil
}

// UpdateURL updates the `url "{{url}}"` part of a formula.
func (f *Formula) UpdateURL(url string) error {
	return nil
}

// UpdateSHA256 updates the `sha256 "{{sha}}"` part of a formula.
func (f *Formula) UpdateSHA256(sha string) error {
	return nil
}
