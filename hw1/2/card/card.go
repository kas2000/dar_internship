package card

//Card is a structure......
type card struct {
	num string
	lear string
}

//NewCard is a constructor of a Card structure
func NewCard(num string, lear string) card {
	c := card {num, lear}
	return c
}

//Num returns card's number
func (c card) Num() string {
	return c.num
}

//Lear returns card's lear
func (c card) Lear() string {
	return c.lear
}

//SetNum is setter for number of a card
func (c *card) SetNum(num string) {
	c.num = num
}

//SetLear is setter for lear of a card
func (c *card) SetLear(lear string) {
	c.lear = lear
}