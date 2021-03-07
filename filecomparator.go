package filecomparator

// comparator file comparator
type comparator struct {
}

// New returns new comparator instance
func New() *comparator {
	return &comparator{}
}

// CompareMD5 runs md5 comparation of two files. if two are equal, it returns true
func (m *comparator) CompareMD5(filePath, compareFilePath string) (bool, error) {
	mdx, err := getFileMD5(filePath)
	if err != nil {
		return false, err
	}
	mdy, err := getFileMD5(compareFilePath)
	if err != nil {
		return false, err
	}
	return (mdx == mdy), nil
}
