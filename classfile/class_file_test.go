package classfile

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parse(t *testing.T, path string) *ClassFile {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return Parse(d)
}

func TestClassFile(t *testing.T) {
	c := parse(t, "../testdata/org.springframework.boot.loader.jar.JarFileEntries.class")

	assert.Equal(t, "org/springframework/boot/loader/jar/JarFileEntries", c.ThisClass())
	assert.Equal(t, "java/lang/Object", c.SuperClass())
	assert.Equal(t, []string{
		"org/springframework/boot/loader/jar/CentralDirectoryVisitor",
		"java/lang/Iterable",
	}, c.Interfaces())
	assert.Equal(t, ClassAccessFlagSuper, c.AccessFlags())
}
