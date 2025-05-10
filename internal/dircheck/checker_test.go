package dircheck

import (
	"os"
	"path/filepath"
	"testing"
)

// Test CheckFolder function works in different options.
func TestCheckFolder(t *testing.T) {
	tempDir := t.TempDir()

	// Prepare Test Set
	//  1. Folder contain Image only
	path1 := filepath.Join(tempDir, "folder1")
	os.MkdirAll(path1, 0755)
	file1, _ := os.Create(filepath.Join(path1, "image1.jpg"))
	defer file1.Close()

	//  2. Folder contain another folder only
	path2 := filepath.Join(tempDir, "folder2")
	os.MkdirAll(filepath.Join(path2, "subfolder2"), 0755)

	//  3. Folder contain both subfolder & Image
	path3 := filepath.Join(tempDir, "folder3")
	os.MkdirAll(filepath.Join(path3, "subfolder3"), 0755)
	file2, _ := os.Create(filepath.Join(path3, "image3.jpg"))
	defer file2.Close()

	//  3. Empty Folder
	path4 := filepath.Join(tempDir, "folder4")
	os.MkdirAll(path4, 0755)

	// Start Image Test
	var tests = []struct {
		path     string
		opt      ScanOpt
		want     bool
		hasError bool
	}{
		// Image Opt Test (1~12)
		{path1, ScanOpt{Image: Unspecific}, true, false},
		{path1, ScanOpt{Image: Allow}, true, false},
		{path1, ScanOpt{Image: AllowOnly}, true, false},
		{path1, ScanOpt{Image: Reject}, false, false},

		{path2, ScanOpt{Image: Unspecific}, true, false},
		{path2, ScanOpt{Image: Allow}, false, false},
		{path2, ScanOpt{Image: AllowOnly}, false, false},
		{path2, ScanOpt{Image: Reject}, true, false},

		{path3, ScanOpt{Image: Unspecific}, true, false},
		{path3, ScanOpt{Image: Allow}, true, false},
		{path3, ScanOpt{Image: AllowOnly}, false, false},
		{path3, ScanOpt{Image: Reject}, false, false},

		{path4, ScanOpt{Image: Unspecific}, true, false},
		{path4, ScanOpt{Image: Allow}, false, false},
		{path4, ScanOpt{Image: AllowOnly}, false, false},
		{path4, ScanOpt{Image: Reject}, true, false},

		// Subfolder Test (13~24)
		{path1, ScanOpt{SubFolder: Unspecific}, true, false},
		{path1, ScanOpt{SubFolder: Allow}, false, false},
		{path1, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path1, ScanOpt{SubFolder: Reject}, true, false},

		{path2, ScanOpt{SubFolder: Unspecific}, true, false},
		{path2, ScanOpt{SubFolder: Allow}, true, false},
		{path2, ScanOpt{SubFolder: AllowOnly}, true, false},
		{path2, ScanOpt{SubFolder: Reject}, false, false},

		{path3, ScanOpt{SubFolder: Unspecific}, true, false},
		{path3, ScanOpt{SubFolder: Allow}, true, false},
		{path3, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path3, ScanOpt{SubFolder: Reject}, false, false},

		{path4, ScanOpt{SubFolder: Unspecific}, true, false},
		{path4, ScanOpt{SubFolder: Allow}, false, false},
		{path4, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path4, ScanOpt{SubFolder: Reject}, true, false},

		// Contradiction Test
		{path3, ScanOpt{SubFolder: AllowOnly, Image: Allow}, false, true},
		{path3, ScanOpt{Image: AllowOnly, SubFolder: Allow}, false, true},
	}

	// Loop the test case and check the result
	for i, testCase := range tests {
		result, err := CheckFolder(testCase.path, testCase.opt)

		// Prevent expected has error, but result is nil
		if testCase.hasError && err == nil {
			t.Errorf("Failed Test Case %d. Expected Error occur", i+1)
			continue
		}

		// Prevent expected not error, but result has error
		if !testCase.hasError && err != nil {
			t.Errorf("Failed Test Case %d. Expected Error Free, got %v", i+1, err)
		}

		// Prevent expected value not matched
		if result != testCase.want {
			t.Errorf("Failed Test Case %d. Expected %v, got %v", i+1, testCase.want, result)
		}
	}

}
