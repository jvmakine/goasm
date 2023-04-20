package class

import (
	"github.com/jvmakine/goasm/classfile"
)

type Class struct {
	file *classfile.Classfile
}

func NewClass(from *classfile.Classfile) *Class {
	return &Class{from}
}

func (c *Class) AccessFlags() *ClassAccess {
	return &ClassAccess{&c.file.AccessFlags}
}

func (c *Class) ThisClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.ThisClass}
}

func (c *Class) SetThisClass(to *ClassInfo) {
	validateFilesEqual(c.file, to.file)
	c.file.ThisClass = to.index
}

func (c *Class) SuperClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.SuperClass}
}

func (c *Class) SetSuperClass(to *ClassInfo) {
	validateFilesEqual(c.file, to.file)
	c.file.ThisClass = to.index
}

func (c *Class) Interfaces() []*ClassInfo {
	result := make([]*ClassInfo, len(c.file.Interfaces))
	for i, ci := range c.file.Interfaces {
		result[i] = &ClassInfo{c.file, ci}
	}
	return result
}

func (c *Class) SetInterfaces(to []*ClassInfo) {
	indices := make([]uint16, len(to))
	for i, ci := range to {
		validateFilesEqual(c.file, ci.file)
		indices[i] = ci.index
	}
	c.file.Interfaces = indices
}

func (c *Class) Constants() *Constants {
	return &Constants{c.file}
}

func (c *Class) Fields() *Fields {
	return &Fields{c.file}
}

func (c *Class) Methods() *Methods {
	return &Methods{c.file}
}

func (c *Class) Attributes() *Attributes {
	return &Attributes{c.file, c.file.Attributes}
}

func validateFilesEqual(f1, f2 *classfile.Classfile) {
	if f1 != f2 {
		panic("can not combine values from different files")
	}
}
