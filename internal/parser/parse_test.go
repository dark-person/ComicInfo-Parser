package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFolder(t *testing.T) {
	// Prepare Struct
	type testCase struct {
		foldername   string // Folder name, as Input
		wantBookName string // Output Book name
		wantAuthor   string // Output Author
		wantMarket   string // Output Market
	}

	// Prepare List
	//cSpell:disable
	tests := []testCase{
		{`(C97) [老眼郷 (老眼)] ネロねろ！2 finale (Fate／Grand Order)[ExtraAi个人汉化]`,
			"ネロねろ！2 finale (Fate／Grand Order)[ExtraAi个人汉化]", "老眼郷 (老眼)", "C97"},
		{`[Sweet Avenue (カヅチ)] 田舎の黒ギャルJKと結婚しました 1-3`,
			`田舎の黒ギャルJKと結婚しました 1-3`, "Sweet Avenue (カヅチ)", ""},
		{`[白桃亭 (rikazu)] 子作り実施科目。絶倫の僕を優しく筆おろししてくれるクラスの人気ギャル [中国翻訳]`,
			"子作り実施科目。絶倫の僕を優しく筆おろししてくれるクラスの人気ギャル [中国翻訳]", "白桃亭 (rikazu)", ""},
		{`[生き恥ハミングバード (天野どん)] Gals Showdown (勝利の女神：NIKKE) [中国翻訳] [DL版]`,
			"Gals Showdown (勝利の女神：NIKKE) [中国翻訳] [DL版]", "生き恥ハミングバード (天野どん)", ""},
		{`[しまぱん (立花オミナ)] 異世界ハーレム物語 4.5 (Side Story)`,
			"異世界ハーレム物語 4.5 (Side Story)", "しまぱん (立花オミナ)", ""},
	}
	//cSpell:enable

	// Run Test Case
	for idx, tt := range tests {
		gotMarket, gotAuthor, gotBookName := ParseFolder(tt.foldername)

		// Compare value
		assert.EqualValuesf(t, tt.wantBookName, gotBookName, "Case %d : Book name is not equal.", idx)
		assert.EqualValuesf(t, tt.wantAuthor, gotAuthor, "Case %d : Author is not equal.", idx)
		assert.EqualValuesf(t, tt.wantMarket, gotMarket, "Case %d : Market is not equal.", idx)
	}
}
