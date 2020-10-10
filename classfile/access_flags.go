package classfile

// ClassAccessFlag decode permissions and properties of a class or interface
type ClassAccessFlag uint16

// Class Access Flags
var (
	ClassAccessFlagPublic     = ClassAccessFlag(0x0001)
	ClassAccessFlagFinal      = ClassAccessFlag(0x0010)
	ClassAccessFlagSuper      = ClassAccessFlag(0x0020)
	ClassAccessFlagInterface  = ClassAccessFlag(0x0200)
	ClassAccessFlagAbstract   = ClassAccessFlag(0x0400)
	ClassAccessFlagSynthetic  = ClassAccessFlag(0x1000)
	ClassAccessFlagAnnotation = ClassAccessFlag(0x2000)
	ClassAccessFlagEnum       = ClassAccessFlag(0x4000)
)

var classAccessFlags = []ClassAccessFlag{
	ClassAccessFlagPublic,
	ClassAccessFlagFinal,
	ClassAccessFlagSuper,
	ClassAccessFlagInterface,
	ClassAccessFlagAbstract,
	ClassAccessFlagSynthetic,
	ClassAccessFlagAnnotation,
	ClassAccessFlagEnum,
}

var classAccessFlagNames = []string{
	"public", "final", "super", "interface", "abstract", "synthetic", "annotation", "enum",
}

// ToString converts access flags to the corresponding string names
func (a ClassAccessFlag) ToString() []string {
	r := []string{}
	for i, f := range classAccessFlags {
		if f&a != 0 {
			r = append(r, classAccessFlagNames[i])
		}
	}
	return r
}

// MethodAccessFlag decode permissions and properties of a method
type MethodAccessFlag uint16

// MethodAccessFlag definitions
var (
	MethodAccessFlagPublic       = MethodAccessFlag(0x0001)
	MethodAccessFlagPrivate      = MethodAccessFlag(0x0002)
	MethodAccessFlagProtected    = MethodAccessFlag(0x0004)
	MethodAccessFlagStatic       = MethodAccessFlag(0x0008)
	MethodAccessFlagFinal        = MethodAccessFlag(0x0010)
	MethodAccessFlagSynchronized = MethodAccessFlag(0x0020)
	MethodAccessFlagBridge       = MethodAccessFlag(0x0040)
	MethodAccessFlagVarargs      = MethodAccessFlag(0x0080)
	MethodAccessFlagNative       = MethodAccessFlag(0x0100)
	MethodAccessFlagAbstract     = MethodAccessFlag(0x0400)
	MethodAccessFlagStrict       = MethodAccessFlag(0x0800)
	MethodAccessFlagSynthetic    = MethodAccessFlag(0x1000)
)

var methodAccessFlags = []MethodAccessFlag{
	MethodAccessFlagPublic,
	MethodAccessFlagPrivate,
	MethodAccessFlagProtected,
	MethodAccessFlagStatic,
	MethodAccessFlagFinal,
	MethodAccessFlagSynchronized,
	MethodAccessFlagBridge,
	MethodAccessFlagVarargs,
	MethodAccessFlagNative,
	MethodAccessFlagAbstract,
	MethodAccessFlagStrict,
	MethodAccessFlagSynthetic,
}

var methodAccessFlagNames = []string{
	"public", "private", "protected", "static", "final", "synchronized",
	"bridge", "varargs", "native", "abstract", "strict", "synthetic",
}

// ToString converts access flags to the corresponding string names
func (a MethodAccessFlag) ToString() []string {
	r := []string{}
	for i, f := range methodAccessFlags {
		if f&a != 0 {
			r = append(r, methodAccessFlagNames[i])
		}
	}
	return r
}
