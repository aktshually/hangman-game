package utils

type Chances int

func (chances Chances) DecreaseByOne() {
	chances -= 1
}
