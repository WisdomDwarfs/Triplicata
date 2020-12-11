package Model


// Data type
type Abstract struct {
	
	Complecti string 
	Length int 
	Status bool 
	Level int 
	
}


type MerkelBlock struct{
	*Abstract
	*MerkelBlock 
}