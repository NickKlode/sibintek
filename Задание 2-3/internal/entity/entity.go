package entity

type number struct {
	Num int `json:"random_number"`
}

type NumberBlocks struct {
	Blocks []number `json:"blocks"`
}
