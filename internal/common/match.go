package common 

type Match struct {
	Season int 
	Mode int // foreign key reference to 
	TotalKills int 
	TotalDamage float64
}
