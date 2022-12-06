package main

import (
	"fmt"
	"math/big"
)




// check prime
type RSA struct {
	p *big.Int
	q *big.Int
	n *big.Int
	e *big.Int
	d *big.Int
}

func (x *RSA) init(p *big.Int, q *big.Int) {
	x.p = p
	x.q = q

}
func (x *RSA) genKey() {
	n := big.NewInt(0)
	n.Mul(x.p, x.q)
	x.n = n

	fmt.Println(x.n)

	phi := big.NewInt(0)
	p1, q1 := big.NewInt(0), big.NewInt(0)
	p1 = p1.Sub(x.p, big.NewInt(1))
	q1 = q1.Sub(x.q, big.NewInt(1))
	phi = phi.Mul(p1, q1)

	fmt.Println("Phi", phi)
	// calulate e
	e := big.NewInt(2)
	for true {
		tmp := big.NewInt(0)
		if tmp.GCD(nil, nil, e, phi).Cmp(big.NewInt(1)) == 0 {
			break
		}
		fmt.Println("e", e)
		e.Add(e, big.NewInt(1))
	}
	x.e = e
	fmt.Println("e", x.e)
	// calculate d
	d := big.NewInt(2)
	for true {
		tmp := big.NewInt(0)
		tmp = tmp.Mul(e, d)
		if tmp.Mod(tmp, phi).Cmp(big.NewInt(1)) == 0 {
			break
		}
		// fmt.Printf("\r%d",d)
		d.Add(d, big.NewInt(1))
	}
	x.d = d
	fmt.Println(d)

}

func main() {

	rsa := RSA{p: big.NewInt(1223), q: big.NewInt(1217)}
	rsa.genKey()
	fmt.Println(rsa.n)

	// tmp := big.NewInt(0)
	// tmp.GCD(nil, nil, big.NewInt(16),big.NewInt(24))
	// fmt.Println(tmp)

	// c := big.NewInt(729000)
	// d := big.NewInt(990635)
	// n := big.NewInt(1488391)

	// m := c.Exp(c, d, n)
	// fmt.Printf("Message %d", m)

}
