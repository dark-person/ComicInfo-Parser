package main

// import (
// 	"gui-comicinfo/internal/archive"
// 	"gui-comicinfo/internal/comicinfo"

// 	"encoding/xml"
// 	"fmt"
// 	"gui-comicinfo/internal/scanner"
// 	"os"
// 	"path/filepath"
// 	"testing"
// )

// // Test Scanning a Complete Directory.
// func TestScan(t *testing.T) {
// 	pages := scanner.GetPageInfo(`temp\Give My Regards to Black Jack\Give My Regards to Black Jack_v01`)

// 	v := comicinfo.New()
// 	v.Pages = pages
// 	v.PageCount = len(pages)

// 	output, err := xml.MarshalIndent(v, "", "  ")
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}

// 	result :=
// 		`<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
//   <Title></Title>
//   <Series></Series>
//   <Number></Number>
//   <Volume>0</Volume>
//   <AlternateSeries></AlternateSeries>
//   <AlternateNumber></AlternateNumber>
//   <StoryArc></StoryArc>
//   <StoryArcNumber></StoryArcNumber>
//   <SeriesGroup></SeriesGroup>
//   <Summary></Summary>
//   <Notes></Notes>
//   <Writer></Writer>
//   <Publisher></Publisher>
//   <Imprint></Imprint>
//   <Genre></Genre>
//   <Tags></Tags>
//   <PageCount>0</PageCount>
//   <LanguageISO></LanguageISO>
//   <Format></Format>
//   <AgeRating></AgeRating>
//   <Manga></Manga>
//   <Characters></Characters>
//   <Teams></Teams>
//   <Locations></Locations>
//   <ScanInformation></ScanInformation>
//   <Pages></Pages>
// </ComicInfo>`

// 	if string(output) != result {
// 		t.Errorf("Result not matched")
// 		os.Stdout.Write(output)
// 	}

// 	os.Stdout.Write(output)
// }

// // Test below process:
// //  1. Create ComicInfo.xml for testFolder
// //  2. Zip the testFolder into .zip
// //  3. Rename .zip to .cbz, and wrap it with folder
// //
// // This test is complicated and require manual compare to determine test pass or not.
// func TestZip(t *testing.T) {
// 	testFolder := `temp\(C102) [そちゃ屋 (にこびぃ)] 高雄さんの性事情 (艦隊これくしょん -艦これ-) [空気系☆漢化]`

// 	// Create ComicInfo.xml first
// 	// Load Abs Path
// 	c := scanner.ScanBooks(testFolder)

// 	output, err := xml.MarshalIndent(c, "  ", "    ")
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}

// 	// Open File for reading
// 	f, err := os.Create(filepath.Join(testFolder, "ComicInfo.xml"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	// Write XML Content to file
// 	f.Write([]byte("<?xml version=\"1.0\"?>\n"))
// 	f.Write(output)
// 	f.Sync()

// 	// Start Archive
// 	filename, _ := archive.CreateZip(testFolder)
// 	err = archive.RenameZip(filename)
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}
// }

// // This Test is for checking internal whole process only. This test is passed with any situation
// func TestProduce(t *testing.T) {
// 	absPath := `D:\Desktop\(C102) [丸杏亭 (マルコ)] ヒミツの××開発 (原神)｜秘密的××开发 [黎欧出资汉化]`

// 	if absPath == "" {
// 		panic("Missing absolute path")
// 	}

// 	// Load Abs Path
// 	c := scanner.ScanBooks(absPath)

// 	output, err := xml.MarshalIndent(c, "  ", "    ")
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}

// 	// Open File for reading
// 	f, err := os.Create(filepath.Join(absPath, "ComicInfo.xml"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	// Write XML Content to file
// 	f.Write([]byte("<?xml version=\"1.0\"?>\n"))
// 	f.Write(output)
// 	f.Sync()

// 	// Start Archive
// 	filename, _ := archive.CreateZip(absPath)
// 	err = archive.RenameZip(filename)
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}
// }
