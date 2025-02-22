package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	var j int64 = 2
	pSub := big.NewInt(0).Sub(p, big.NewInt(2))
	if pSub.Cmp(big.NewInt(0)) <= 0 {
		pSub = big.NewInt(1)
		j = 1
	}
	ran, _ := rand.Int(rand.Reader, pSub)
	return ran.Add(ran, big.NewInt(j))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	prv := PrivateKey(p)
	pub := PublicKey(prv, p, g)
	return prv, pub
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
