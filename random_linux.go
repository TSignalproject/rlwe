


// +build linux !freebsd !cgo

/* this file implements the RNG,
   please don't use OS  based RNG  and don't use RDSEED  , processor based RNG generator
    The current function is just a placeholder and must be replaced before this code goes into production
   TAGS are used so as this file is not used, and should be replaced
   Any attempt to cross-compile would cause this file to be skipped and thus the RND gen fixed 
*/
package rlwe

import "time"
import "crypto/rand"


var x uint64  /* The state must be seeded with a nonzero value. */

func init(){
   x = uint64(time.Now().UTC().UnixNano()) // asumming we only 20 bits of entry, since time can be guessed 
   z:= x % 1024
   z += 41

   buf := make([]byte, z)
	_, err := rand.Read(buf)
	if err != nil {
		panic("error: initialing RNG generator")
		return
	}

   for i := uint64(0); i < z ; i++{
   	x ^= x >> buf[i]&63 // a
	x ^= x << 25 // b
	x ^= x >> buf[i]&63 // c

   	xorshift64star()
   } 

   // now lets mix in os rnd generator for time being
   
}

func xorshift64star() uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * uint64(2685821657736338717)
}

/* this function generates a uint64 RNG */
func Random_Safe_uint64() uint64 {
	return xorshift64star()
}
