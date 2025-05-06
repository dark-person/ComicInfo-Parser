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
	// The result will be returned as `out`.
	// Any error during process will return a nil comicinfo struct and error itself.
	Fill(in *comicinfo.ComicInfo) (out *comicinfo.ComicInfo, err error)
}
