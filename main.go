package main

import(
	Britannica "github.com/WisdomDwarfs/Britannica/BritannicaMathematica"
	codex "github.com/WisdomDwarfs/Triplicata/Mathematica"
	"github.com/WisdomDwarfs/Triplicata/Model"
	"fmt"
	// "reflect"
	// "strings"
)
var ( 
	fileOp Britannica.FileOperations = Britannica.NewInstance()
	Alice Britannica.CodexBritannica = Britannica.NewVault()
	Codex codex.MathematicaAbstract = codex.Init() 
   )

func main() {

	file, err := fileOp.FileExist("iqbal"); if err != nil { 
		fmt.Println("Error", err) 
		return 
	} 
	
	fmt.Println(file) 
	f, err , hash := Alice.CodexMathematica(file); if err != nil { 
		fmt.Println("Error", err) 
		return 
	}
	fmt.Println("File:", f)
	
	dynamo := make([]*Model.Abstract, len(hash)-1)
	// easy way to solve problem
	 for i, _ := range hash {
		s1 := Codex.CreateTriplicata(hash[i])
		code := Codex.ReduceSizeOf(s1); if code != nil{
			dynamo = append(dynamo,code)
		}
	 }
	 MoveLikeTurning(dynamo)

	 
	 
  }


 func MoveLikeTurning(str []*Model.Abstract){
	 //s := make([]*Model.Abstract, len(str)-1)
 	if len(str)> 0 {
 		for i:=0; i < len(str)-1;i++ {
			state := Codex.Nil(str[i]); if !state{
				a := Codex.Value(str[i])
				b := Codex.Value(str[i+1])
				h := Codex.HashTriplicata(a,b)
				fmt.Println(h)
				
			}
			// block := Codex.BlocksinLevel(str[i])
			
    	}
	 }
 }