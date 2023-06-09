package class_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/jvmakine/goasm/class"
	"github.com/jvmakine/goasm/classfile"
	"github.com/stretchr/testify/require"
)

func TestAccessFlags(t *testing.T) {
	t.Run("returns access flags for a public concrete class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		access := clazz.AccessFlags()

		require.True(t, access.IsPublic())
		require.False(t, access.IsFinal())
		require.True(t, access.IsSuper())
		require.False(t, access.IsAbstract())
		require.False(t, access.IsAnnotation())
		require.False(t, access.IsEnum())
		require.False(t, access.IsInterface())
		require.False(t, access.IsSynthetic())
	})
	t.Run("changes access flags", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		clazz.AccessFlags().SetSynthetic(true)

		require.True(t, clazz.AccessFlags().IsSynthetic())
	})
}

func TestClassInfo(t *testing.T) {
	t.Run("returns name of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		require.Equal(t, "com/github/jvmakine/test/Hello", clazz.ThisClass().Name().Text())
	})
	t.Run("returns name of the super class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		require.Equal(t, "java/lang/Object", clazz.SuperClass().Name().Text())
	})
	t.Run("returns interfaces correctly", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, ci := range clazz.Interfaces().List() {
			names = append(names, ci.Name().Text())
		}
		require.Equal(t, []string{"java/io/Serializable"}, names)
	})
}

func TestFields(t *testing.T) {
	t.Run("returns field names of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, f := range clazz.Fields().List() {
			names = append(names, f.Name().Text())
		}
		require.Equal(t, []string{"foo", "empty"}, names)
	})
	t.Run("returns field descriptors of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, f := range clazz.Fields().List() {
			names = append(names, f.Descriptor().Text())
		}
		require.Equal(t, []string{"Ljava/lang/String;", "Lcom/github/jvmakine/test/Empty;"}, names)
	})
}

func TestMethods(t *testing.T) {
	t.Run("returns method names of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, f := range clazz.Methods().List() {
			names = append(names, f.Name().Text())
		}
		require.Equal(t, []string{"<init>", "hello"}, names)
	})
	t.Run("returns method descriptors of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, f := range clazz.Methods().List() {
			names = append(names, f.Descriptor().Text())
		}
		require.Equal(t, []string{
			"(Ljava/lang/String;Lcom/github/jvmakine/test/Empty;)V",
			"()V",
		}, names)
	})
}

func TestAttributes(t *testing.T) {
	t.Run("returns attribute names for a class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var names []string
		for _, f := range clazz.Attributes().List() {
			names = append(names, f.Name().Text())
		}
		require.Equal(t, []string{
			"SourceFile",
		}, names)
	})
}

func classFrom(t *testing.T, path string) *class.Class {
	t.Helper()

	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	classFile, err := classfile.Parse(bytes.NewReader(data))
	require.NoError(t, err)
	return class.NewClass(classFile)
}
