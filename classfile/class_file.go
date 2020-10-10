package classfile

import (
	"fmt"
)

/*
ClassFile {
	u4				magic;
	u2 				minor_version;
	u2 				major_version;
	u2 				constant_pool_count;
	cp_info 		constant_pool[constant_pool_count-1];
	u2 				access_flags;
	u2 				this_class;
	u2 				super_class;
	u2 				interfaces_count;
	u2 				interfaces[interfaces_count];
	u2 				fields_count;
	field_info 		fields[fields_count];
	u2 				methods_count;
	method_info 	methods[methods_count];
	u2 				attributes_count;
	attribute_info 	attributes[attributes_count];
}
*/
type ClassFile struct {
	size         int
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  ClassAccessFlag
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []MemberInfo
	methods      []MemberInfo
	attributes   []AttributeInfo
}

// Parse a class file from raw bytes
func Parse(bytes []byte) *ClassFile {
	reader := NewClassReader(bytes)
	return ParseFromClassReader(reader)
}

// ParseFromClassReader parses a class from a class reader
func ParseFromClassReader(reader IClassReader) *ClassFile {
	cf := &ClassFile{}
	cf.size = reader.Length()
	cf.magic = reader.ReadUint32()
	cf.minorVersion = reader.ReadUint16()
	cf.majorVersion = reader.ReadUint16()
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = ClassAccessFlag(reader.ReadUint16())
	cf.thisClass = reader.ReadUint16()
	cf.superClass = reader.ReadUint16()
	cf.readInterfaces(reader)
	//TODO from now on, we can speed up by run concurrently
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
	return cf
}

func (cf *ClassFile) readInterfaces(reader IClassReader) {
	cf.interfaces = make([]uint16, reader.ReadUint16())
	for i := 0; i < len(cf.interfaces); i++ {
		cf.interfaces[i] = reader.ReadUint16()
	}
}

// Size of the class in bytes
func (cf *ClassFile) Size() int {
	return cf.size
}

// Magic number of the class file
func (cf *ClassFile) Magic() uint32 {
	return cf.magic
}

// MajorVersion of the class file
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// MinorVersion of the class file
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// ConstantPool holds all constant values in the class file
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// AccessFlags for the class
func (cf *ClassFile) AccessFlags() ClassAccessFlag {
	return cf.accessFlags
}

// ThisClass name
func (cf *ClassFile) ThisClass() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

// SuperClass name
func (cf *ClassFile) SuperClass() string {
	return cf.constantPool.getClassName(cf.superClass)
}

// Interfaces names as strings
func (cf *ClassFile) Interfaces() []string {
	r := []string{}
	for _, i := range cf.interfaces {
		r = append(r, cf.constantPool.getClassName(i))
	}
	return r
}

// Fields info
func (cf *ClassFile) Fields() []MemberInfo {
	return cf.fields
}

// Methods info
func (cf *ClassFile) Methods() []MemberInfo {
	return cf.methods
}

// Attributes info
func (cf *ClassFile) Attributes() []AttributeInfo {
	return cf.attributes
}

// Print the class info
func (cf *ClassFile) Print() {
	fmt.Printf("Size: %d bytes\n", cf.size)
	fmt.Printf("magic: %x\n", cf.magic)
	fmt.Printf("minor version: %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)

	fmt.Printf("accessFlags: %d\n", cf.accessFlags)
	fmt.Printf("thisClass: #%d\n", cf.thisClass)
	fmt.Printf("superClass: #%d\n", cf.superClass)

	fmt.Println("**********************************************************")
	for i, length := 1, len(cf.constantPool); i < length; i++ {
		fmt.Printf(" #%2d = ", i)
		if cp, ok := cf.constantPool[i].(*ConstantClassInfo); ok {
			fmt.Printf("Class\t\t#%d\t\t\t// %s", cp.nameIndex, cp.String(cf.constantPool))
		} else if cp, ok := cf.constantPool[i].(*ConstantFieldrefInfo); ok {
			fmt.Printf("Fieldref\t\t#%d.#%d\t\t\t// %s", cp.classIndex, cp.nameAndTypeIndex, cp.String(cf.constantPool))
		} else if cp, ok := cf.constantPool[i].(*ConstantMethodrefInfo); ok {
			fmt.Printf("Methodref\t#%d.#%d\t\t\t// %s", cp.classIndex, cp.nameAndTypeIndex, cp)
		} else if cp, ok := cf.constantPool[i].(*ConstantUtf8Info); ok {
			fmt.Printf("Utf8\t\t%s", cp.String())
		} else if cp, ok := cf.constantPool[i].(*ConstantStringInfo); ok {
			fmt.Printf("String\t\t#%d\t\t\t// %s", cp.stringIndex, cp.String(cf.constantPool))
		} else if cp, ok := cf.constantPool[i].(*ConstantNameAndTypeInfo); ok {
			fmt.Printf("NameAndType\t#%d:#%d\t\t\t// %s", cp.nameIndex, cp.descriptorIndex, cp.String(cf.constantPool))
		}
		fmt.Println()
	}
}
