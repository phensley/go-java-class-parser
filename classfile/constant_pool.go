package classfile

import "fmt"

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPool []ConstantPoolInfo

type ConstantPoolInfo interface {
	ReadInfo(reader IClassReader)
}

func (p ConstantPool) GetConstantInfo(index uint16) ConstantPoolInfo {
	return p[index]
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.GetConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.GetConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.GetConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.String()
}

func readConstantPool(reader IClassReader) ConstantPool {
	constantPool := make([]ConstantPoolInfo, reader.ReadUint16())
	for i := 1; i < len(constantPool); i++ {
		constType := reader.ReadUint8()
		cpInfo := newConstantPoolInfo(constType, constantPool)
		if cpInfo != nil {
			cpInfo.ReadInfo(reader)
			constantPool[i] = cpInfo
		}
		switch constType {
		case CONSTANT_Double:
			i++
		case CONSTANT_Long:
			i++
		}
	}
	return constantPool
}

func newConstantPoolInfo(constType uint8, cp ConstantPool) ConstantPoolInfo {
	switch constType {
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{cp: cp}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic(fmt.Sprintf("Invalid const type: %d", constType))
	}
}
