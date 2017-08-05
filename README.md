# RLWE : Ring Learning With Errors

This go package implements the RLWE ( Ring learning with errors) a Lattice based Key Exchange.

The algorithm is safe from quantum computer attacks. ( Since it is not based on factoring problem)
This implementation is  a non-constant time Gaussian based RLWE with params **n = 1024 q=40961**
The above parameters have been proved to deliver **minimal 2^256 bit security, upper limit not yet found.**





You can read more  at https://en.wikipedia.org/wiki/Ring_learning_with_errors

Why n= 1024 & q=40961 can be found at http://www.ringlwe.info/parameters-for-rlwe.html