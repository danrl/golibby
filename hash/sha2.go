package hash

var sha256K = []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
	0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01,
	0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa,
	0x5cb0a9dc, 0x76f988da, 0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
	0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967, 0x27b70a85, 0x2e1b2138,
	0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624,
	0xf40e3585, 0x106aa070, 0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
	0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3, 0x748f82ee, 0x78a5636f,
	0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

// SHA256 implements the NIST SHA2-256 cryptographic hashsum
func SHA256(message []byte) [32]byte {
	h := []uint32{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f,
		0x9b05688c, 0x1f83d9ab, 0x5be0cd19}

	// original length in bits
	l := uint64(len(message)) * 8
	// pad message
	message = append(message, 0x80)
	for (len(message) % 64) < 56 {
		message = append(message, 0x00)
	}
	// append length information
	for j := uint(0); j < 64; j += 8 {
		message = append(message, byte(l>>(56-j)))
	}

	// chunk message
	for i := 0; i < len(message); i += 64 {
		// initialize words (first 16 words copied over from current chunk)
		w := [64]uint32{}
		for j := 0; j < 16; j++ {
			for n := 0; n < 4; n++ {
				w[j] |= uint32(message[i+(j*4)+n]) << uint(24-n*8)
			}
		}

		// extend words (word 16 to 63 computed based on first 16 words)
		for j := 16; j < 64; j++ {
			s0 := (w[j-15]>>7 | w[j-15]<<25) ^ (w[j-15]>>18 | w[j-15]<<14)
			s0 ^= w[j-15] >> 3
			s1 := (w[j-2]>>17 | w[j-2]<<15) ^ (w[j-2]>>19 | w[j-2]<<13)
			s1 ^= w[j-2] >> 10
			w[j] = w[j-16] + s0 + w[j-7] + s1
		}

		// init working variables
		v := [8]uint32{}
		for j := 0; j < 8; j++ {
			v[j] = h[j]
		}

		// main compression loop (variable names loosely based on NIST document)
		for j := 0; j < 64; j++ {
			s1 := (v[4]>>6 | v[4]<<26) ^ (v[4]>>11 | v[4]<<21) ^ (v[4]>>25 | v[4]<<7)
			ch := (v[4] & v[5]) ^ ((^v[4]) & v[6])
			t1 := v[7] + s1 + ch + sha256K[j] + w[j]
			s0 := (v[0]>>2 | v[0]<<30) ^ (v[0]>>13 | v[0]<<19) ^ (v[0]>>22 | v[0]<<10)
			maj := (v[0] & v[1]) ^ (v[0] & v[2]) ^ (v[1] & v[2])
			t2 := s0 + maj

			// rotate working variables around
			for j := 7; j > 0; j-- {
				v[j] = v[j-1]
			}
			v[4] += t1
			v[0] = t1 + t2
		}

		// add compression result
		for j := 0; j < 8; j++ {
			h[j] += v[j]
		}
	}

	// move sums bytewise to digest slice
	var digest [32]byte
	for i, s := range h {
		digest[i*4] = byte(s >> 24)
		digest[i*4+1] = byte(s >> 16)
		digest[i*4+2] = byte(s >> 8)
		digest[i*4+3] = byte(s)
	}
	return digest
}
