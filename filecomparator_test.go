package filecomparator_test

import (
	"os/user"
	"testing"

	"github.com/chixm/filecomparator"
)

func TestCompare(t *testing.T) {
	// create file path string
	user, _ := user.Current()
	file1 := user.HomeDir + `/512.zip`
	file2 := user.HomeDir + `/512.zip`

	comp := filecomparator.New()
	result, _ := comp.CompareMD5(file1, file2)

	// result is true since same file are compared.
	t.Log(result)
}

func TestCompareDiff(t *testing.T) {
	// create file path string
	user, _ := user.Current()
	file1 := user.HomeDir + `/512.zip`
	file2 := user.HomeDir + `/NTUSER.DAT`

	comp := filecomparator.New()
	result, _ := comp.CompareMD5(file1, file2)

	// result is false since different files are chosen.
	t.Log(result)
}

func TestCompareError(t *testing.T) {
	// create file path string
	user, _ := user.Current()
	file1 := user.HomeDir + `/512.zip`
	file2 := user.HomeDir + `/notExistFile.txt`

	comp := filecomparator.New()
	result, err := comp.CompareMD5(file1, file2)
	// result is false when error occurs. for example, FILE NOT FOUND
	t.Log(result)
	t.Log(err)
}

func TestCompareSha1(t *testing.T) {
	user, _ := user.Current()
	file1 := user.HomeDir + `/512.zip`
	file2 := user.HomeDir + `/512.zip`

	comp := filecomparator.New()
	result, _ := comp.CompareSHA1(file1, file2)

	// result is true since same file are compared.
	t.Log(result)
}
