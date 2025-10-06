package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// -Key Components
// Deterministic 相同的输入总是产生相同的哈希值。
// Fast Computation 对于给定的输入，可以快速计算出哈希值。
// Pre-image Resistance 单向性（Pre-image Resistance）：给定一个哈希值h，很难反向计算出原始输入x，使得H(x)=h。
// Collision Resistance 很难找到两个不同的输入x和x'，使得H(x)=H(x')。
// Avalanche Effect：输入中微小的变化（比如一位改变）会导致输出的哈希值产生巨大的变化（大约一半的位会改变）。
// SHA-256
// SHA-512
// Salting
// -Best Practices
// Use of Standard Libraries
// Algorithm Updates



func main() {
	password:="password123"
	// hash:=sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(hash) //[239 146 183 120 186 254 119 30 137 36 91 137 236 188 8 164 74 78 22 108 6 101 153 17 136 31 56 61 68 115 233 79]
	// fmt.Println(hash512) //[190 212 239 161 212 253 189 149 75 211 112 93 106 42 120 39 14 201 165 46 207 191 176 16 198 24 98 175 92 118 175 23 97 255 235 26 239 106 202 27 245 208 43 55 129 170 133 79 171 210 182 156 121 13 231 78 23 236 254 195 203 106 196 191]
	// fmt.Printf("SHA-256 Hash hex val %x\n",hash) //ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512) //bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffeb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf


	salt,err:=generateSalt()
	fmt.Println("Original Salt:", salt) //[79 1 22 182 29 130 29 228 172 77 234 36 8 8 31 235]
	fmt.Printf("Original Salt: %x\n", salt) //4f0116b61d821de4ac4dea2408081feb
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	//hash password with salt
	signUphash:=hashPassword(password,salt)

	//Store the salt and password in database,just printing as of now
	saltStr:=base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:",saltStr) //TwEWth2CHeSsTeokCAgf6w==  simulate as storing in database
	fmt.Println("hash::",signUphash) //r4LxJ26iSBbVfj7uNHbKPixCk62BSEe6kri3tiDkp6k=  simulate as storing in database
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify password
	// retrieve the saltStr and decode it
	decodedSalt,err:=base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash:=hashPassword(password,decodedSalt)
	//compare the stored signUphash with loginHash
	if signUphash == loginHash {
		fmt.Println("Password is correct. you are logged in.")
	} else{
		fmt.Println("Login failed. Please check user credentials.")
	}
	
}

func generateSalt() ([]byte,error) {
	salt:=make([]byte,16)
	_,err:=io.ReadFull(rand.Reader,salt)
	if err!=nil {
		return nil,err
	}
	return salt,nil
}

//function to hash password
func hashPassword (password string, salt []byte) string{
	saltedPassword:=append(salt,[]byte(password)...)
	hash:=sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
