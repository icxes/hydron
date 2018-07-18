//go:generate easyjson --all $GOFILE

package common

import "bytes"

type TagSource uint8

const (
	User TagSource = iota
	Gelbooru
	Hydrus
)

type TagType uint8

const (
	Undefined TagType = iota
	Author
	Character
	Series
	Rating
	System
)

// Tag of an image
type Tag struct {
	Type   TagType   `json:"type"`
	Source TagSource `json:"source"`
	ID     uint64    `json:"-"`
	Tag    string    `json:"tag"` // Not always defined for performance reasons
}

// Convert tag to normalized string representation
func (t Tag) String() string {
	var typ string
	switch t.Type {
	case Author:
		typ = "author"
	case Character:
		typ = "character"
	case Series:
		typ = "series"
	case System:
		typ = "system"
	}
	var buf bytes.Buffer
	if typ != "" {
		buf.WriteString(typ)
		buf.WriteByte(':')
	}
	buf.WriteString(t.Tag)
	return buf.String()
}

// Record of an image stored in the database
type Image struct {
	CompactImage
	Dims
	ID         int64  `json:"-"`
	ImportTime int64  `json:"import_time"`
	Size       int    `json:"size"`
	Duration   uint64 `json:"duration,omitempty"`
	MD5        string `json:"md5"`
	// Not always defined for performance reasons
	Tags []Tag `json:"tags,omitempty"`
}

// Common fields
type Dims struct {
	Width  uint64 `json:"width"`
	Height uint64 `json:"height"`
}

// Thumbnail of image
type Thumbnail struct {
	IsPNG bool `json:"is_png"`
	Dims
}

// Only provides the most minimal of fields. Optimal for thumbnail views.
type CompactImage struct {
	Type  FileType  `json:"type"`
	SHA1  string    `json:"sha1"`
	Thumb Thumbnail `json:"thumb"`
}
