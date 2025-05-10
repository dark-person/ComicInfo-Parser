// Package for providers, which allow filling comicinfo struct by different source.
package dataprovider

import "github.com/dark-person/comicinfo-parser/internal/comicinfo"

// Provider interface for struct that fill comicinfo data.
//
// All provider should implement this interface for consistant structure.
// To ensure interface is valid, developer MUST add below code:
//  var _ dataprovider.DataProvider = (*MyDataProvider)(nil)
type DataProvider interface {

	// Fill input comicinfo by provider's method.
	// There has two possiblity of fill method:
	//
	//  1. Fill success, return changed comicinfo and nil error
	//  2. Fill failed, return unchanged comicinfo (same as input) and error it self
	Fill(input *comicinfo.ComicInfo) (result *comicinfo.ComicInfo, err error)
}
