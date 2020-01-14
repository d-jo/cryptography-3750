package main

import "fmt"

func toIntArr(text string) []int {
	var res []int
	for _, char := range text {
		res = append(res, int(char)-65)
	}
	return res
}

func toString(arr []int) string {
	var res string = ""
	for _, char := range arr {
		res += string(char + 65)
	}
	return res
}

func main() {
	var ct = toIntArr("KINGOFSHESHACHSHALLDRINK")
	ct = atbash(ct)
	fmt.Println(toString(ct))

	ct = toIntArr("GSRHDSLOVOZMWHSZOOYVXLNVZIFRMZMWZDZHGV")
	ct = atbash(ct)
	fmt.Println(ct)
	fmt.Println(toString(ct))

	ct = toIntArr("QFHGZHDZGVIIVUOVXGHGSVUZXVHLLMVSFNZMSVZIGIVUOVXGHZMLGSVI")
	ct = atbash(ct)
	fmt.Println(toString(ct))

	var scy = scytale_dec(toIntArr("EIEPAADSNERDUATREVCNIIFEAONTURTRPYGSINRAEIOAONITNMSDNY"), 6)
	fmt.Println(toString(scy))

	scy = scytale_dec(toIntArr("TWETIHAXROESAANSAMNCCNPSIYELPPTAEOHAROSELLFIREYAT"), 5)
	fmt.Println(toString(scy))

	var vg = vigenere_enc_autokey(toIntArr("WALTERRALEIGHBRINGSTOBACCOTOENGLANDFROMAMERICA"), toIntArr("B")[0])
	fmt.Println(toString(vg))
	fmt.Println(toString(vigenere_dec_autokey(vg, toIntArr("B")[0])))
}

func atbash(text []int) []int {
	var res []int
	for _, char := range text {
		res = append(res, 25-char)
	}
	return res
}

func scytale_dec(text []int, lines int) []int {
	var res [][]int
	for i := 0; i < lines; i++ {
		res = append(res, []int{})
	}
	for pos, char := range text {
		res[pos%lines] = append(res[pos%lines], char)
	}
	var flat []int
	for _, content := range res {
		flat = append(flat, content...)
	}
	return flat
}

func vigenere_enc(text []int, key []int) []int {
	var res []int
	for pos, char := range text {
		res = append(res, (char+key[pos])%26)
	}
	return res
}

func vigenere_enc_autokey(text []int, singleKey int) []int {
	var res []int
	var key []int
	key = append(key, singleKey)
	key = append(key, text[:len(text)-1]...)
	for pos, char := range text {
		res = append(res, (char+key[pos])%26)
	}
	return res
}

func vigenere_dec(text []int, key []int) []int {
	var res []int
	for pos, char := range text {
		res = append(res, (char-key[pos]+26)%26)
	}
	return res
}

func vigenere_dec_autokey(text []int, key int) []int {
	var res []int
	res = append(res, key)
	for pos, char := range text {
		res = append(res, (char-res[pos]+26)%26)
	}
	return res
}
