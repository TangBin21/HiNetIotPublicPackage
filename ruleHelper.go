// ruleHelper
package HiNetEdgeDAS

const (
	CellFormatBit   = "bit"   //1个位
	CellFormatByte  = "byte"  //1个字节
	CellFormatWord  = "word"  //2个字节
	CellFormatTbyte = "tbyte" //3个字节
	CellFormatDword = "dword" //4个字节
	CellFormatFbyte = "fbyte" //5个字节
	CellFormatTword = "tword" //6个字节
)

/*
func GainVarAddressLen(cellType string, dType string, strLen uint32) int32 {
	var ret int32 = -1
	switch dType {
	case BOOL:
		ret = 1
	case INT8, UINT8:
		if cellType == CellFormatByte {
			ret = 1
		} else {
			ret = -1
		}
	case INT16, UINT16:
		if cellType == CellFormatByte {
			ret = 2
		} else if cellType == CellFormatWord {
			ret = 1
		} else {
			ret = -1
		}
	case INT32, UINT32, FLOAT:
		if cellType == CellFormatByte {
			ret = 4
		} else if cellType == CellFormatWord {
			ret = 2
		} else if cellType == CellFormatDword {
			ret = 1
		} else {
			ret = -1
		}
	case DOUBLE:
		if cellType == CellFormatByte {
			ret = 8
		} else if cellType == CellFormatWord {
			ret = 4
		} else if cellType == CellFormatDword {
			ret = 2
		} else {
			ret = -1
		}
	case STRING:
		if strLen == 0 {
			return -1
		}
		ret = int32(strLen)
	case FP_STRING16:
		if strLen == 0 {
			return -1
		}
		ret = int32(strLen)
		// if strLen == 0 {
		// 	return -1
		// }
		// if cellType == CellFormatWord {
		// 	ret = int32((strLen + 1) / 2)
		// } else {
		// 	ret = -1
		// }
	case FP_STRING32:
		if strLen == 0 {
			return -1
		}
		ret = int32(strLen)
		// if strLen == 0 {
		// 	return -1
		// }
		// if cellType == CellFormatWord {
		// 	ret = int32((strLen + 1) / 2)
		// } else {
		// 	ret = -1
		// }
	case HEX:
		if strLen == 0 {
			return -1
		}
		ret = int32(strLen)
		// if cellType == CellFormatByte {
		// 	ret = int32(strLen)
		// } else if cellType == CellFormatWord {
		// 	ret = int32((strLen + 1) / 2)
		// } else {
		// 	ret = -1
		// }
	case OBJECT:
		ret = 1
	case S7TCP_Counter:
		if cellType == CellFormatTbyte {
			ret = 1
		} else {
			ret = -1
		}
	case S7TCP_Timers:
		if cellType == CellFormatFbyte {
			ret = 1
		} else {
			ret = -1
		}
	case S7TCP_Date:
		if cellType == CellFormatByte {
			ret = 2
		} else {
			ret = -1
		}
	case S7TCP_Time:
		if cellType == CellFormatByte {
			ret = 4
		} else {
			ret = -1
		}
	default:
		return -1
	}
	return ret
}
*/
