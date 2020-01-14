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

	var scy = scytale(toIntArr("EIEPAADSNERDUATREVCNIIFEAONTURTRPYGSINRAEIOAONITNMSDNY"), 6)
	fmt.Println(toString(scy))
}

func atbash(text []int) []int {
	var res []int
	for _, char := range text {
		res = append(res, 25-char)
	}
	return res
}

func scytale_dec(text []int, lines int) []int {
	var res [6][]int
	for pos, char := range text {
		res[pos%lines] = append(res[pos%lines], char)
	}
	var flat []int
	for _, content := range res {
		flat = append(flat, content...)
	}
	return flat
}
