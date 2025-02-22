package converter

import (
	"fmt"
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/converter/cbz"
	"github.com/metafates/mangal/converter/pdf"
	"github.com/metafates/mangal/converter/plain"
	"github.com/metafates/mangal/converter/zip"
	"github.com/metafates/mangal/source"
)

// Converter is the interface that all converters must implement.
type Converter interface {
	Save(chapter *source.Chapter) (string, error)
	SaveTemp(chapter *source.Chapter) (string, error)
}

var converters = map[string]Converter{
	constant.Plain: plain.New(),
	constant.CBZ:   cbz.New(),
	constant.PDF:   pdf.New(),
	constant.ZIP:   zip.New(),
}

// Available returns a list of available converters.
func Available() []string {
	return []string{
		constant.Plain,
		constant.CBZ,
		constant.PDF,
		constant.ZIP,
	}
}

// Get returns a converter by name.
// If the converter is not available, an error is returned.
func Get(name string) (Converter, error) {
	if converter, ok := converters[name]; ok {
		return converter, nil
	}

	return nil, fmt.Errorf("unkown format \"%s\"", name)
}
