package convetor

import (
	"fmt"
	"strconv"
	"encoding/binary"
)

func Float64ToString(f float64) string {
	return fmt.Sprintf("%g", f)
}

func StringToFloat64(val string) float64 {
	currVal, _ := strconv.ParseFloat(val, 64)
	return currVal
}

// 四舍五入
func Float64Rand(v float64, dig int) float64 {
	cDig := strconv.Itoa(dig)
	val := fmt.Sprintf("%0."+cDig+"f", v)
	return StringToFloat64(val)
}

// 浮点数串化(左边是整数位置,右边是小数位,dig参数控制)
func FloatToFDig(floVal float64, dig int) string {
	return fmt.Sprintf("%10."+strconv.Itoa(dig)+"f", floVal) //十位整数，8位小数
}

// byte转浮点数数组
func ByteToFloat32Array(bf []byte, featureSize int) []float32 {
	feature := make([]float32, featureSize)
	for i := 0; i < featureSize; i++ {
		off := i * 4
		feature[i] = byteToFloat32(bf[off : off+4])
	}
	return feature
}

func byteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}


func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

// Float32ArrayToByte
func Float32ArrayToByte(fs []float32) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(fs); i++ {
		tmp := Float32ToByte(fs[i])
		buf.Write(tmp)
	}
	return buf.Bytes()
}
