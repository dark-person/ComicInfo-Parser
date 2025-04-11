package comicinfo

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
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
