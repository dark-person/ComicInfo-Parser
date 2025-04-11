package comicinfo

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Create new ComicInfo object, then unmarshal to XML.
func TestNew(t *testing.T) {
	v := New()

	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if string(output) != emptyMarshal {
		t.Errorf("Result not matched")
		os.Stdout.Write(output)
	}
}

func TestAddTags(t *testing.T) {
	// Test Case Structure
	type testCase struct {
		c          *ComicInfo
		tags       []string
		wantedTags string
	}

	// Prepare Test Case
	case1 := New()
	case2 := New()
	case3 := New()
	case3.Tags = "pre-tag1,pre-tag2"
	case4 := New()
	case4.Tags = "pre-tag1,pre-tag2"
	case5 := New()
	case6 := New()

	tests := []testCase{
		// 1. Add Single Tag with empty tags in comic info
		{&case1, []string{"tag1"}, "tag1"},
		// 2. Add Multiple Tag with empty tags in comic info
		{&case2, []string{"tag1", "tag2"}, "tag1,tag2"},
		// 3. Add Single Tag with tags exist in comic info
		{&case3, []string{"tag1"}, "pre-tag1,pre-tag2,tag1"},
		// 4. Add multiple Tag with tags exist in comic info
		{&case4, []string{"tag1", "tag2"}, "pre-tag1,pre-tag2,tag1,tag2"},
		// 5. Add tag with "," characters
		{&case5, []string{"tag1,tag2"}, "tag1,tag2"},
		// 6. Add tag with empty slice
		{&case6, []string{""}, ""},
	}

	for idx, tt := range tests {
		tt.c.AddTags(tt.tags...)

		// Validate valid
		assert.EqualValuesf(t, tt.wantedTags, tt.c.Tags, "Case %d: not expected value.", idx)
	}
}

// ===============
const emptyMarshal = `<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <Title></Title>
  <Series></Series>
  <Number></Number>
  <Volume>0</Volume>
  <AlternateSeries></AlternateSeries>
  <AlternateNumber></AlternateNumber>
  <StoryArc></StoryArc>
  <StoryArcNumber></StoryArcNumber>
  <SeriesGroup></SeriesGroup>
  <Summary></Summary>
  <Notes></Notes>
  <Writer></Writer>
  <Publisher></Publisher>
  <Imprint></Imprint>
  <Genre></Genre>
  <Tags></Tags>
  <PageCount>0</PageCount>
  <LanguageISO></LanguageISO>
  <Format></Format>
  <AgeRating></AgeRating>
  <Manga></Manga>
  <Characters></Characters>
  <Teams></Teams>
  <Locations></Locations>
  <ScanInformation></ScanInformation>
  <Pages></Pages>
</ComicInfo>`
