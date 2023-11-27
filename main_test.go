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
