package utils

type AlreadyTried []string

func (alreadyTried AlreadyTried) ContainsLetter(letter string) bool {
	for _, alreadyTriedLetter := range alreadyTried {
		if letter == alreadyTriedLetter {
			return true
		}
	}
	return false
}

func (alreadyTried AlreadyTried) AddLetter(letter string) {
	alreadyTried = append(alreadyTried, letter)
}
