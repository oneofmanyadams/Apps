package main

type CharacterPool struct {
	Pool string
	PoolSize int
}

func NewPool(characters string) (cp CharacterPool) {
	cp.Pool = characters
	cp.PoolSize = len(cp.Pool)
	return
}

func (cp CharacterPool) CharacterFromNumber(num int) string {
	if num < 0 {
		return cp.CharacterFromNumber(0)
	}

	if num >= cp.PoolSize {
		quotent := (int(num) / int(26)) - 1
		remainder := num % cp.PoolSize
		
		return cp.CharacterFromNumber(quotent) + cp.CharacterFromNumber(remainder)
	}

	return string(cp.Pool[num])
}