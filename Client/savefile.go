package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
)

var book []string

type server struct {
	ip string
}

func (s *server) Send(text string) {
	conn, _ := net.Dial("tcp", s.ip)
	fmt.Fprint(conn, text)

}

func chunk() []string {
	//define ip and port
	s1 := server{ip: "192.168.242.2:8081"}
	s2 := server{ip: "192.168.242.3:8081"}
	s3 := server{ip: "192.168.242.4:8081"}
	s4 := server{ip: "192.168.242.5:8081"}
	var chunkServ []string

	fileToBeChunked := "send.enc"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 1 * (1 << 20) // Fichier de 1Mb

	//NB chunk
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	y := 0
	for i := uint64(0); i < totalPartsNum; i++ {
		//y permet d'enregistrer dans 4 endroit différent
		y++

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		// define name
		var Numchunk = strconv.FormatUint(i, 10)
		fileName := "chunk" + Numchunk
		// save file
		save := rand.Intn(5-1) + 1

		for y == save {
			save = rand.Intn(5-1) + 1
		}

		savename1 := "s" + strconv.Itoa(y)
		savename2 := "s" + strconv.Itoa(save)
		fmt.Println(fileName, savename1, savename2)

		switch savename1 {
		case "s1":
			s1.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l1_s1 "+fileName)
		case "s2":
			s2.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l1_s2 "+fileName)
		case "s3":
			s3.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l1_s3 "+fileName)
		case "s4":
			s4.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l1_s4 "+fileName)
		default:
			fmt.Println("dude wtf")
		}
		switch savename2 {
		case "s1":
			s1.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l2_s1 "+fileName)
		case "s2":
			s2.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l2_s2 "+fileName)
		case "s3":
			s3.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l2_s3 "+fileName)
		case "s4":
			s4.Send(fileName + " " + string(partBuffer))
			chunkServ = append(chunkServ, "l2_s4 "+fileName)
		default:
			fmt.Println("dude wtf")
		}

		if i%4 == 0 {
			y = 0
		}
	}
	fmt.Println("STOP")
	s1.Send("STOP")
	s2.Send("STOP")
	s3.Send("STOP")
	s4.Send("STOP")
	return chunkServ
}
func encryptFile(key []byte, filename string, outFilename string) (string, error) {
	if len(outFilename) == 0 {
		outFilename = filename + ".enc"
	}

	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	of, err := os.Create(outFilename)
	if err != nil {
		return "", err
	}
	defer of.Close()

	// Write the original plaintext size into the output file first, encoded in
	// a 8-byte integer.
	origSize := uint64(len(plaintext))
	if err = binary.Write(of, binary.LittleEndian, origSize); err != nil {
		return "", err
	}

	// Pad plaintext to a multiple of BlockSize with random padding.
	if len(plaintext)%aes.BlockSize != 0 {
		bytesToPad := aes.BlockSize - (len(plaintext) % aes.BlockSize)
		padding := make([]byte, bytesToPad)
		if _, err := rand.Read(padding); err != nil {
			return "", err
		}
		plaintext = append(plaintext, padding...)
	}

	// Generate random IV and write it to the output file.
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	if _, err = of.Write(iv); err != nil {
		return "", err
	}

	// Ciphertext has the same size as the padded plaintext.
	ciphertext := make([]byte, len(plaintext))

	// Use AES implementation of the cipher.Block interface to encrypt the whole
	// file in CBC mode.
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	if _, err = of.Write(ciphertext); err != nil {
		return "", err
	}
	return outFilename, nil
}

// decryptFile decrypts the file specified by filename with the given key. See
// doc for encryptFile for more details.
func decryptFile(key []byte, filename string, outFilename string) (string, error) {
	if len(outFilename) == 0 {
		outFilename = filename + ".dec"
	}

	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	of, err := os.Create(outFilename)
	if err != nil {
		return "", err
	}
	defer of.Close()

	// cipertext has the original plaintext size in the first 8 bytes, then IV
	// in the next 16 bytes, then the actual ciphertext in the rest of the buffer.
	// Read the original plaintext size, and the IV.
	var origSize uint64
	buf := bytes.NewReader(ciphertext)
	if err = binary.Read(buf, binary.LittleEndian, &origSize); err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err = buf.Read(iv); err != nil {
		return "", err
	}

	// The remaining ciphertext has size=paddedSize.
	paddedSize := len(ciphertext) - 8 - aes.BlockSize
	if paddedSize%aes.BlockSize != 0 {
		return "", fmt.Errorf("want padded plaintext size to be aligned to block size")
	}
	plaintext := make([]byte, paddedSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext[8+aes.BlockSize:])

	if _, err := of.Write(plaintext[:origSize]); err != nil {
		return "", err
	}
	return outFilename, nil
}

func main() {
	slist := server{ip: "192.168.242.6:8081"}
	// key := bytes.Repeat([]byte("1"), 32)

	// outFilename, err := encryptFile(key, "100Mio.dat", "send.enc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Encrypted output file:", outFilename)
	returnlink := chunk()
	fmt.Println(returnlink)
	for _, value := range returnlink {
		fmt.Println(value)
		slist.Send(value)
	}
	slist.Send("STOP")
}
