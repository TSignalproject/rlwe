
package main

import "fmt"
import "rlwe"

func main() {
	fmt.Printf("Init RWLE\n")

	/*Exclusively For Alice*/
	var s_alice [2 * rlwe.M]rlwe.RINGELT   /* Alice's Private Key */
	var mu_alice [rlwe.Muwords]uint64 /* Alice's recovered mu */

	/*Exclusively For Bob*/
	var mu_bob [rlwe.Muwords]uint64 /* Bob's version of mu */

	/*Information that gets shared by Alice and Bob*/
	var b_alice [rlwe.M]rlwe.RINGELT   /* Alice's Public Key */
	var u [rlwe.M]rlwe.RINGELT         /* Bob's Ring Element from Encapsulation */
	var cr_v [rlwe.Muwords]uint64 /* Cross Rounding of v */

	/*for i:=0 ; i < 100;i++{
		fmt.Printf("%4d: %d\n",i,RANDOM8())
	}*/

	rlwe.KEM1_Generate(&s_alice, &b_alice)
	// KEM1_Generate(s_alice,b_alice)

	for i := 1000; i < 1024; i++ {
		fmt.Printf("%4d: %8d  %8d  %d  %d\n", i, s_alice[i], s_alice[i+1024], b_alice[i], len(s_alice))
	}

	fmt.Printf("Keys initialised\n")

	public_alice := b_alice[:1024]

	rlwe.KEM1_Encapsulate(&u, &cr_v, &mu_bob, public_alice)

	/* 	 for i:=0 ; i < 16;i++{
	   	 	fmt.Printf("%4d: %d %d %dx\n",i,u[i],cr_v[i],mu_bob[i])
	   	 }
	*/

	private_alice := s_alice[1024:]
	u_copy := u[:]
	rlwe.KEM1_Decapsulate(&mu_alice, u_copy, private_alice, cr_v)

	for i := 0; i < 16; i++ {
		fmt.Printf("%4d: %d %d \n", i, mu_bob[i], mu_alice[i])
	}

}
 
