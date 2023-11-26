package main

import (
	"changeme/internal/archive"
	"changeme/internal/comicinfo"

	"changeme/internal/parser"
	"changeme/internal/scanner"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// Test Scanning a Complete Directory.
func TestScan(t *testing.T) {
	pages := scanner.GetPageInfo(`temp\Give My Regards to Black Jack\Give My Regards to Black Jack_v01`)

	v := comicinfo.New()
	v.Pages = pages
	v.PageCount = len(pages)

	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	result :=
		`<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
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

	if string(output) != result {
		t.Errorf("Result not matched")
		os.Stdout.Write(output)
	}

	os.Stdout.Write(output)

}

// Test Parse filenames to related comic info.
func TestParse(t *testing.T) {
	// Empty Now
	filenames := []string{
		`(C97) [老眼郷 (老眼)] ネロねろ！2 finale (Fate／Grand Order)[ExtraAi个人汉化]`,
		`[Sweet Avenue (カヅチ)] 田舎の黒ギャルJKと結婚しました 1-3`,
		`[白桃亭 (rikazu)] 子作り実施科目。絶倫の僕を優しく筆おろししてくれるクラスの人気ギャル [中国翻訳]`,
		`[生き恥ハミングバード (天野どん)] Gals Showdown (勝利の女神：NIKKE) [中国翻訳] [DL版]`,
	}

	// Result
	titleResult := []string{"ネロねろ！2 finale (Fate／Grand Order)[ExtraAi个人汉化]",
		"田舎の黒ギャルJKと結婚しました 1-3",
		"子作り実施科目。絶倫の僕を優しく筆おろししてくれるクラスの人気ギャル [中国翻訳]",
		"Gals Showdown (勝利の女神：NIKKE) [中国翻訳] [DL版]"}
	authorResult := []string{"老眼郷 (老眼)", "Sweet Avenue (カヅチ)", "白桃亭 (rikazu)", "生き恥ハミングバード (天野どん)"}
	marketResult := []string{"C97", "", "", ""}
	tagResult := []string{"C97", "", "", "DL版"}

	for i, filename := range filenames {

		// Test XML
		market, author, bookName := parser.Parse(filename)
		c := comicinfo.New()
		c.Title = bookName
		c.Writer = author
		c.Manga = comicinfo.Manga_Yes
		if market != "" {
			c.Imprint = market
			c.AddTags(market)
		}

		// Add Special Tags
		tags := scanner.GetSpecialTags(filename)
		c.AddTags(tags...)

		// Matches
		if c.Title != titleResult[i] {
			t.Errorf("Title not matched, %v, %v", c.Title, titleResult[i])
		}

		if c.Writer != authorResult[i] {
			t.Errorf("Writer not matched, %v, %v", c.Writer, authorResult[i])
		}

		if c.Imprint != marketResult[i] {
			t.Errorf("Imprint not matched, %v, %v", c.Imprint, marketResult[i])
		}

		if c.Tags != tagResult[i] {
			t.Errorf("Tags not matched, %v, %v", c.Tags, tagResult[i])

		}
	}
}

// Test below process:
//  1. Create ComicInfo.xml for testFolder
//  2. Zip the testFolder into .zip
//  3. Rename .zip to .cbz, and wrap it with folder
//
// This test is complicated and require manual compare to determine test pass or not.
func TestZip(t *testing.T) {
	testFolder := `temp\(C102) [そちゃ屋 (にこびぃ)] 高雄さんの性事情 (艦隊これくしょん -艦これ-) [空気系☆漢化]`

	// Create ComicInfo.xml first
	// Load Abs Path
	c := scanner.ScanBooks(testFolder)

	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Open File for reading
	f, err := os.Create(filepath.Join(testFolder, "ComicInfo.xml"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write XML Content to file
	f.Write([]byte("<?xml version=\"1.0\"?>\n"))
	f.Write(output)
	f.Sync()

	// Start Archive
	filename, _ := archive.CreateZip(testFolder)
	err = archive.RenameZip(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

// This Test is for checking internal whole process only. This test is passed with any situation
func TestProduce(t *testing.T) {
	absPath := `D:\Desktop\(C102) [丸杏亭 (マルコ)] ヒミツの××開発 (原神)｜秘密的××开发 [黎欧出资汉化]`

	if absPath == "" {
		panic("Missing absolute path")
	}

	// Load Abs Path
	c := scanner.ScanBooks(absPath)

	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Open File for reading
	f, err := os.Create(filepath.Join(absPath, "ComicInfo.xml"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write XML Content to file
	f.Write([]byte("<?xml version=\"1.0\"?>\n"))
	f.Write(output)
	f.Sync()

	// Start Archive
	filename, _ := archive.CreateZip(absPath)
	err = archive.RenameZip(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
