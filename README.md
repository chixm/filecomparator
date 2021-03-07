# filecomparator
Compares two files are equal or not.

# Project FUSO
File Utils for Simple Objectives

### Compare Files in multiple ways
filecomparator can tell two files are same or not.
it compares files with md5 hash values, sha1 values and other original compare methods.

file compare methods
- MD5
- sha1
- other

## Examples
create a instance with New method and
call a method you want to compare files with.
```
	// create file path string
	user, _ := user.Current()
	file1 := user.HomeDir + `/512.zip`
	file2 := user.HomeDir + `/512.zip`

	comp := filecomparator.New()
	result, _ := comp.CompareMD5(file1, file2)

	// result is true since same file are compared.
	t.Log(result)

```
