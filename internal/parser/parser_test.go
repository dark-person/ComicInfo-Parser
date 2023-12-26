package parser

import (
	"gui-comicinfo/internal/comicinfo"
	"testing"
)

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
		market, author, bookName := ParseFolder(filename)
		c := comicinfo.New()
		c.Title = bookName
		c.Writer = author
		c.Manga = comicinfo.Manga_Yes
		if market != "" {
			c.Imprint = market
			c.AddTags(market)
		}

		// Add Special Tags
		tags := GetSpecialTags(filename)
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
