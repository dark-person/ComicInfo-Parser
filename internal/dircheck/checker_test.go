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
		opt      DirectoryOpt
		want     bool
		hasError bool
	}{
		// Image Opt Test (1~12)
		{path1, DirectoryOpt{Image: Unspecific}, true, false},
		{path1, DirectoryOpt{Image: Allow}, true, false},
		{path1, DirectoryOpt{Image: AllowOnly}, true, false},
		{path1, DirectoryOpt{Image: Reject}, false, false},

		{path2, DirectoryOpt{Image: Unspecific}, true, false},
		{path2, DirectoryOpt{Image: Allow}, false, false},
		{path2, DirectoryOpt{Image: AllowOnly}, false, false},
		{path2, DirectoryOpt{Image: Reject}, true, false},

		{path3, DirectoryOpt{Image: Unspecific}, true, false},
		{path3, DirectoryOpt{Image: Allow}, true, false},
		{path3, DirectoryOpt{Image: AllowOnly}, false, false},
		{path3, DirectoryOpt{Image: Reject}, false, false},

		{path4, DirectoryOpt{Image: Unspecific}, true, false},
		{path4, DirectoryOpt{Image: Allow}, false, false},
		{path4, DirectoryOpt{Image: AllowOnly}, false, false},
		{path4, DirectoryOpt{Image: Reject}, true, false},

		// Subfolder Test (13~24)
		{path1, DirectoryOpt{SubFolder: Unspecific}, true, false},
		{path1, DirectoryOpt{SubFolder: Allow}, false, false},
		{path1, DirectoryOpt{SubFolder: AllowOnly}, false, false},
		{path1, DirectoryOpt{SubFolder: Reject}, true, false},

		{path2, DirectoryOpt{SubFolder: Unspecific}, true, false},
		{path2, DirectoryOpt{SubFolder: Allow}, true, false},
		{path2, DirectoryOpt{SubFolder: AllowOnly}, true, false},
		{path2, DirectoryOpt{SubFolder: Reject}, false, false},

		{path3, DirectoryOpt{SubFolder: Unspecific}, true, false},
		{path3, DirectoryOpt{SubFolder: Allow}, true, false},
		{path3, DirectoryOpt{SubFolder: AllowOnly}, false, false},
		{path3, DirectoryOpt{SubFolder: Reject}, false, false},

		{path4, DirectoryOpt{SubFolder: Unspecific}, true, false},
		{path4, DirectoryOpt{SubFolder: Allow}, false, false},
		{path4, DirectoryOpt{SubFolder: AllowOnly}, false, false},
		{path4, DirectoryOpt{SubFolder: Reject}, true, false},

		// Contradiction Test
		{path3, DirectoryOpt{SubFolder: AllowOnly, Image: Allow}, false, true},
		{path3, DirectoryOpt{Image: AllowOnly, SubFolder: Allow}, false, true},
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
