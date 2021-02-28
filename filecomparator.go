package filecomparator

var c *Comparator

func init() {
	c = &Comparator{}
}

// Comparator file comparator
type Comparator struct {
}

// CompareMD5 runs md5 comparation of two files. if two are equal, it returns true
func (m *Comparator) CompareMD5(filePath, compareFilePath string) (bool, error) {

	return false, nil
}
