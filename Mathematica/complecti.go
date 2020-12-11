package Mathematica

import(
	"crypto/sha256"
	"encoding/hex"
	"github.com/WisdomDwarfs/Triplicata/Model"
	"reflect"
	"strings"
	//"fmt"
)


var (
	size int = 0
	
)

type MathematicaAbstract interface {
	
	CreateTriplicata(s string) *Model.Abstract
	HashTriplicata(blockA  , blockB string) *Model.Abstract
	// AddBlockinList(blockA Abstract) []Abstract
	EmptyList(len int) bool
	BlocksinLevel(block *Model.Abstract) int
	ReduceSizeOf(m *Model.Abstract)*Model.Abstract
	FieldCompare(s, st1 string)int
	Value(i interface{})string
	Nil(i interface{}) bool
}

type AbstractCipher struct{}

func Init() MathematicaAbstract{
	return &AbstractCipher{}
}


func SizeUpdate() int {
	size +=1 
	return size
}

func (abstract *AbstractCipher) CreateTriplicata(s string) *Model.Abstract {
	inc := SizeUpdate()
	art := &Model.Abstract{
		Complecti : s, 
		Length : inc, 
		Status : false,
		Level : 0,
	}
	return art
}

func (abstract * AbstractCipher) ReduceSizeOf(m *Model.Abstract)*Model.Abstract{

	value := reflect.ValueOf(m).Elem()

	valueX := abstract.FieldCompare(value.Field(0).String(), " ")
	if valueX == 1{
		code := abstract.CreateTriplicata(value.Field(0).String())
		return code
	}
	return nil
		
}

func (abstract *AbstractCipher) FieldCompare(s, st1 string)int {
	return strings.Compare(s, st1)
}

func (abstract *AbstractCipher) Nil(i interface{}) bool{
	inter := reflect.ValueOf(i).IsNil()
	return inter
}

func (abstract *AbstractCipher) Value(i interface{})string{
	value := reflect.ValueOf(i).Elem()
	return value.Field(0).String()
}

func (abstract *AbstractCipher) HashTriplicata(blockA , blockB string) *Model.Abstract {
	
	
	a := ([]byte(blockA))
	b := ([]byte(blockB))

	 xorblock := make([]byte, len(a))
	 for i, _ := range xorblock {
	 	xorblock[i] = a[i] ^ b[i]
	 }
	 operation := BlockHash(xorblock)
	 block := abstract.CreateTriplicata(operation)
	 block.Level+=1
	 block.Status=true
	 return block
	
}

// func (abstract *AbstractCipher) AddBlockinList(blockA Abstract) []Abstract {
	
// 	state := abstract.EmptyList(size)
// 	blockList := make([]Abstract, blockA.Length)

// 	if !state {
		
// 		blockList = append(blockList, blockA)
// 	}
// 	return blockList
// }


func (abstract *AbstractCipher) EmptyList(len int) bool  {
	if len == 0{
		return true
	}
	return false

}

func BlockHash(block []byte) string{
	h := sha256.New()
	h1 := h.Sum(block)
	encode := hex.EncodeToString(h1)
	return encode
}

func (abstract *AbstractCipher) BlocksinLevel(block *Model.Abstract) int  {
	count := 0;
	return count
}




