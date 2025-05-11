package cfgprov

import (
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	// Prepare test case values
	configured1 := config.Default()
	configured1.Metadata.Number = "1.5"

	configured2 := config.Default()
	configured2.Metadata.Number = "2"

	info1 := comicinfo.New()
	info1.Number = "1.5"

	info2 := comicinfo.New()
	info2.Number = "2"

	// ----------------------------------------

	// Prepare Test cases, which error must not allowed
	type testCase struct {
		cfg     *config.ProgramConfig // configuration
		want    comicinfo.ComicInfo   // expected comicinfo
		wantErr bool                  // If error will occur
	}

	tests := []testCase{
		{config.Default(), comicinfo.New(), false},
		{configured1, info1, false},
		{configured2, info2, false},
		{nil, comicinfo.New(), true},
	}

	// Start test
	for idx, tt := range tests {
		// Prepare new comicinfo
		tc := comicinfo.New()
		c := &tc

		// Run provider
		prov := New(tt.cfg)
		c, err := prov.Fill(c)

		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error occur but nil.", idx)
		} else {
			assert.EqualValuesf(t, &tt.want, c, "Case %d: Unexpected comicinfo result", idx)
		}
	}
}
