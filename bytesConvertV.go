// bytesConvertV
package HiNetEdgeDAS

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

var (
	AssignmentError error = errors.New("query content exceeds the original data")
	DataTypeError   error = errors.New("illegal data type")
)

func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

func SaveDecimalPlace(val float64, dec int) float64 {
	n10 := math.Pow10(dec)
	if val > 0 {
		return math.Trunc((val*n10 + 0.5)) / n10
	} else {
		return math.Trunc((val*n10 - 0.5)) / n10
	}
}

func GainBool(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos  int
		tempBit  = pPoint.BitAddr
		tempByte = make([]byte, 0)
	)
	switch pRule.CellFormat {
	case CellFormatBit:
		tempPos = int(pPoint.Offset - pRule.StartAddr)
		if tempPos/8+1 > len(src) {
			return 0, AssignmentError
		}
		tempByte = append(tempByte, src[tempPos/8])
		if (tempByte[0]>>(tempPos%8))&0x01 == 0x01 {
			return 1, nil
		} else {
			return 0, nil
		}
	case CellFormatByte:
		tempPos = int(pPoint.Offset - pRule.StartAddr)
		if tempPos+1 > len(src) {
			return 0, AssignmentError
		}
		tempByte = append(tempByte, src[tempPos])
		if (tempByte[0]>>tempBit)&0x01 == 0x01 {
			return 1, nil
		} else {
			return 0, nil
		}
	case CellFormatWord:
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
		if tempPos+2 > len(src) {
			return 0, AssignmentError
		}
		if order == "X1X2" {
			tempByte = append(tempByte, src[tempPos], src[tempPos+1])
		} else if order == "X2X1" {
			tempByte = append(tempByte, src[tempPos+1], src[tempPos])
		}
		if tempBit >= 0 && tempBit <= 7 {
			if (tempByte[1]>>tempBit)&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		} else if tempBit >= 8 && tempBit <= 15 {
			if (tempByte[0]>>(tempBit-8))&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case CellFormatDword:
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 4
		if tempPos+4 > len(src) {
			return 0, AssignmentError
		}
		if order == "X1X2X3X4" {
			tempByte = append(tempByte, src[tempPos], src[tempPos+1], src[tempPos+2], src[tempPos+3])
		} else if order == "X4X3X2X1" {
			tempByte = append(tempByte, src[tempPos+3], src[tempPos+2], src[tempPos+1], src[tempPos])
		} else if order == "X2X1X4X3" {
			tempByte = append(tempByte, src[tempPos+1], src[tempPos], src[tempPos+3], src[tempPos+2])
		} else if order == "X3X4X1X2" {
			tempByte = append(tempByte, src[tempPos+2], src[tempPos+3], src[tempPos], src[tempPos+1])
		}
		if tempBit >= 0 && tempBit <= 7 {
			if (tempByte[3]>>tempBit)&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		} else if tempBit >= 8 && tempBit <= 15 {
			if (tempByte[2]>>(tempBit-8))&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		} else if tempBit >= 16 && tempBit <= 23 {
			if (tempByte[1]>>(tempBit-8))&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		} else if tempBit >= 24 && tempBit <= 31 {
			if (tempByte[0]>>(tempBit-8))&0x01 == 0x01 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	}
	return 0, DataTypeError
}

func GainInt8(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos  int
		tempByte = make([]byte, 1)
		tempInt8 int8
	)
	tempPos = int(pPoint.Offset - pRule.StartAddr)
	if tempPos+1 > len(src) {
		return 0, AssignmentError
	}
	tempByte[0] = src[tempPos]
	tempInt8 = int8(tempByte[0])
	return SaveDecimalPlace(float64(tempInt8)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainUint8(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos   int
		tempByte  = make([]byte, 1)
		tempUint8 uint8
	)
	tempPos = int(pPoint.Offset - pRule.StartAddr)
	if tempPos+1 > len(src) {
		return 0, AssignmentError
	}
	tempByte[0] = src[tempPos]
	tempUint8 = uint8(tempByte[0])
	return SaveDecimalPlace(float64(tempUint8)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainInt16(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos   int
		tempByte  = make([]byte, 2)
		tempInt16 int16
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	}
	if tempPos+2 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
	} else if order == "X2X1" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
	}
	tempInt16 = int16(binary.BigEndian.Uint16(tempByte))
	return SaveDecimalPlace(float64(tempInt16)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainUint16(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos    int
		tempByte   = make([]byte, 2)
		tempUint16 uint16
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	}
	if tempPos+2 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
	} else if order == "X2X1" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
	}
	tempUint16 = uint16(binary.BigEndian.Uint16(tempByte))
	return SaveDecimalPlace(float64(tempUint16)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainInt32(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos   int
		tempByte  = make([]byte, 4)
		tempInt32 int32
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	} else if pRule.CellFormat == CellFormatDword {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 4
	}
	if tempPos+4 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2X3X4" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
		tempByte[2] = src[tempPos+2]
		tempByte[3] = src[tempPos+3]
	} else if order == "X2X1X4X3" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
		tempByte[2] = src[tempPos+3]
		tempByte[3] = src[tempPos+2]
	} else if order == "X4X3X2X1" {
		tempByte[0] = src[tempPos+3]
		tempByte[1] = src[tempPos+2]
		tempByte[2] = src[tempPos+1]
		tempByte[3] = src[tempPos]
	} else if order == "X3X4X1X2" {
		tempByte[0] = src[tempPos+2]
		tempByte[1] = src[tempPos+3]
		tempByte[2] = src[tempPos]
		tempByte[3] = src[tempPos+1]
	}
	tempInt32 = int32(binary.BigEndian.Uint32(tempByte))
	return SaveDecimalPlace(float64(tempInt32)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainUint32(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos    int
		tempByte   = make([]byte, 4)
		tempUint32 uint32
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	} else if pRule.CellFormat == CellFormatDword {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 4
	}
	if tempPos+4 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2X3X4" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
		tempByte[2] = src[tempPos+2]
		tempByte[3] = src[tempPos+3]
	} else if order == "X2X1X4X3" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
		tempByte[2] = src[tempPos+3]
		tempByte[3] = src[tempPos+2]
	} else if order == "X4X3X2X1" {
		tempByte[0] = src[tempPos+3]
		tempByte[1] = src[tempPos+2]
		tempByte[2] = src[tempPos+1]
		tempByte[3] = src[tempPos]
	} else if order == "X3X4X1X2" {
		tempByte[0] = src[tempPos+2]
		tempByte[1] = src[tempPos+3]
		tempByte[2] = src[tempPos]
		tempByte[3] = src[tempPos+1]
	}
	tempUint32 = uint32(binary.BigEndian.Uint32(tempByte))
	return SaveDecimalPlace(float64(tempUint32)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainFloat(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos   int
		tempByte  = make([]byte, 4)
		tempFloat float32
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	}
	if tempPos+4 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2X3X4" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
		tempByte[2] = src[tempPos+2]
		tempByte[3] = src[tempPos+3]
	} else if order == "X2X1X4X3" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
		tempByte[2] = src[tempPos+3]
		tempByte[3] = src[tempPos+2]
	} else if order == "X4X3X2X1" {
		tempByte[0] = src[tempPos+3]
		tempByte[1] = src[tempPos+2]
		tempByte[2] = src[tempPos+1]
		tempByte[3] = src[tempPos]
	} else if order == "X3X4X1X2" {
		tempByte[0] = src[tempPos+2]
		tempByte[1] = src[tempPos+3]
		tempByte[2] = src[tempPos]
		tempByte[3] = src[tempPos+1]
	}
	bits := binary.BigEndian.Uint32(tempByte)
	tempFloat = math.Float32frombits(bits)
	return SaveDecimalPlace(float64(tempFloat)*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainDouble(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (float64, error) {
	var (
		tempPos   int
		tempByte  = make([]byte, 8)
		tempFloat float64
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
	}
	if tempPos+8 > len(src) {
		return 0, AssignmentError
	}
	if order == "X1X2X3X4X5X6X7X8" {
		tempByte[0] = src[tempPos]
		tempByte[1] = src[tempPos+1]
		tempByte[2] = src[tempPos+2]
		tempByte[3] = src[tempPos+3]
		tempByte[4] = src[tempPos+4]
		tempByte[5] = src[tempPos+5]
		tempByte[6] = src[tempPos+6]
		tempByte[7] = src[tempPos+7]
	} else if order == "X2X1X4X3X6X5X8X7" {
		tempByte[0] = src[tempPos+1]
		tempByte[1] = src[tempPos]
		tempByte[2] = src[tempPos+3]
		tempByte[3] = src[tempPos+2]
		tempByte[4] = src[tempPos+5]
		tempByte[5] = src[tempPos+4]
		tempByte[6] = src[tempPos+7]
		tempByte[7] = src[tempPos+6]
	} else if order == "X8X7X6X5X4X3X2X1" {
		tempByte[0] = src[tempPos+7]
		tempByte[1] = src[tempPos+6]
		tempByte[2] = src[tempPos+5]
		tempByte[3] = src[tempPos+4]
		tempByte[4] = src[tempPos+3]
		tempByte[5] = src[tempPos+2]
		tempByte[6] = src[tempPos+1]
		tempByte[7] = src[tempPos]
	} else if order == "X7X8X5X6X3X4X1X2" {
		tempByte[0] = src[tempPos+6]
		tempByte[1] = src[tempPos+7]
		tempByte[2] = src[tempPos+4]
		tempByte[3] = src[tempPos+5]
		tempByte[4] = src[tempPos+2]
		tempByte[5] = src[tempPos+3]
		tempByte[6] = src[tempPos+0]
		tempByte[7] = src[tempPos+1]
	}
	bits := binary.BigEndian.Uint64(tempByte)
	tempFloat = math.Float64frombits(bits)
	return SaveDecimalPlace(tempFloat*pPoint.Ratio+pPoint.Cardinality, int(pPoint.DecimalPlace)), nil
}

func GainString(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (string, error) {
	var (
		tempPos   int
		retStr    string = ""
		dst, data        = make([]byte, 0), make([]byte, 0)
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
		if tempPos+int(pPoint.StrLen) > len(src) {
			return "", AssignmentError
		}
		dst = src[tempPos : tempPos+int(pPoint.StrLen)]
		retStr = string(dst)
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
		if tempPos+int(pPoint.StrLen)*2 > len(src) {
			return "", AssignmentError
		}
		dst = src[tempPos : tempPos+int(pPoint.StrLen)*2]
		if order == "X1X2" {

		} else if order == "X2X1" {
			for i := 0; i < len(dst); i += 2 {
				temp := dst[i]
				dst[i] = dst[i+1]
				dst[i+1] = temp
			}
		}
		for i := 0; i < len(dst); i++ {
			if dst[i] == 0 {
				break
			}
			data = append(data, dst[i])
		}
		retStr = string(data)
	}
	return retStr, nil
}

func GainHex(order string, src []byte, pRule *RuleMsg, pPoint *PointMsg) (string, error) {
	var (
		tempPos int
		retStr  string = ""
		dst            = make([]byte, 0)
	)
	if pRule.CellFormat == CellFormatByte {
		tempPos = int(pPoint.Offset - pRule.StartAddr)
		if tempPos+int(pPoint.StrLen) > len(src) {
			return "", AssignmentError
		}
		dst = src[tempPos : tempPos+int(pPoint.StrLen)]
		for _, val := range dst {
			retStr += fmt.Sprintf("%02X", val)
		}
	} else if pRule.CellFormat == CellFormatWord {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 2
		if tempPos+int(pPoint.StrLen)*2 > len(src) {
			return "", AssignmentError
		}
		dst = src[tempPos : tempPos+int(pPoint.StrLen)*2]
		for _, val := range dst {
			retStr += fmt.Sprintf("%02X", val)
		}
	} else if pRule.CellFormat == CellFormatDword {
		tempPos = int(pPoint.Offset-pRule.StartAddr) * 4
		if tempPos+int(pPoint.StrLen)*4 > len(src) {
			return "", AssignmentError
		}
		dst = src[tempPos : tempPos+int(pPoint.StrLen)*4]
		for _, val := range dst {
			retStr += fmt.Sprintf("%02X", val)
		}
	}
	return retStr, nil
}
