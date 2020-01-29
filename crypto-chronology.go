package main

import "fmt"

func main() {

	var str = "EGDHE TGXIN XHCDI LXIWD JIBPC NUTPG HPCSS XHIPG ITHPC SPSKT GHXIN XHCDI LXIWD JIRDB UDGIH PCSWD ETH"
	var freq, count = freq(str)
	fmt.Println(count)
	fmt.Println(freq)
	var text = toIntArrBasic(str)
	for i := 0; i < 26; i++ {
		fmt.Printf("%d\n - %s\n - %s\n", i, str, toStringBasic(shift(text, i)))
	}

	/*
		var name = toIntArrBasic("DECLANJOHNSON")
		var ten = shift(name, 10)
		var sixteen = shift(name, 16)
		var both = shift(ten, 16)
		fmt.Println(toStringBasic(name))
		fmt.Println(toStringBasic(ten))
		fmt.Println(toStringBasic(sixteen))
		fmt.Println(toStringBasic(both))

		/*
			var freqMap, _ = rel_freq("test")
			fmt.Println(freqMap)

			var ct = toIntArrBasic("KINGOFSHESHACHSHALLDRINK")
			ct = atbash(ct)
			fmt.Println(toStringBasic(ct))

			ct = toIntArrBasic("GSRHDSLOVOZMWHSZOOYVXLNVZIFRMZMWZDZHGV")
			ct = atbash(ct)
			fmt.Println(ct)
			fmt.Println(toStringBasic(ct))

			ct = toIntArrBasic("QFHGZHDZGVIIVUOVXGHGSVUZXVHLLMVSFNZMSVZIGIVUOVXGHZMLGSVI")
			ct = atbash(ct)
			fmt.Println(toStringBasic(ct))

			var scy = scytale_dec(toIntArrBasic("EIEPAADSNERDUATREVCNIIFEAONTURTRPYGSINRAEIOAONITNMSDNY"), 6)
			fmt.Println(toStringBasic(scy))

			scy = scytale_dec(toIntArrBasic("TWETIHAXROESAANSAMNCCNPSIYELPPTAEOHAROSELLFIREYAT"), 5)
			fmt.Println(toStringBasic(scy))

			var vg = vigenere_enc_autokey(toIntArrBasic("WALTERRALEIGHBRINGSTOBACCOTOENGLANDFROMAMERICA"), toIntArrBasic("B")[0])
			fmt.Println(toStringBasic(vg))
			fmt.Println(toStringBasic(vigenere_dec_autokey(vg, toIntArrBasic("B")[0])))

			var kk = kid_krypto_init(47, 22, 11, 5)
			fmt.Println(kk)
			var kkEnc = kid_krypto_enc(kk, 1958)
			fmt.Println(kkEnc)
			fmt.Println(kid_krypto_dec(kk, kkEnc))
	*/
}

func shift(text []int, shiftn int) []int {
	var res []int
	for _, char := range text {
		res = append(res, (char+shiftn)%26)
	}
	return res
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

func kid_krypto_init(a, b, A, B int) (res map[string]int) {
	res = make(map[string]int)
	res["M"] = (a * b) - 1
	res["e"] = (A * res["M"]) + a
	res["d"] = (B * res["M"]) + b
	res["n"] = ((res["e"] * res["d"]) - 1) / res["M"]
	return res
}

func kid_krypto_enc(kk map[string]int, x int) int {
	return (kk["e"] * x) % kk["n"]
}

func kid_krypto_dec(kk map[string]int, y int) int {
	return (kk["d"] * y) % kk["n"]
}
