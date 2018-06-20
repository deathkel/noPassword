package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
	"strconv"
	"os"
	"fmt"
	"flag"
	"encoding/json"
)

var keyMap = map[int]string{
	0: "0123456789",
	1: "abcdefghijklmnopqrstuvwxyz",
	2: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

type Item struct {
	Uid          string `json:"uid"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
}

type Result struct {
	Items []*Item `json:"items"`
}

func main() {
	argsWithoutProg := os.Args[1:]
	text := ""
	for _, arg := range argsWithoutProg {
		text += string(arg) + " "
	}
	options := ParseOptions()
	ln := getPassword(text)

	Output(ln[0:options.Length])
}

func Output(output string) {
	item := new(Item)
	item.Uid = "np"
	item.Type = "type"
	item.Title = output
	item.Subtitle = "按回车复制密码"
	item.Arg = output
	item.Autocomplete = "1"
	result := new(Result)
	result.Items = make([]*Item, 1)
	result.Items[0] = item
	res, _ := json.Marshal(result)
	fmt.Println(string(res))
}

type Options struct {
	Length int
}

func ParseOptions() *Options {
	opts := new(Options)
	flag.IntVar(&opts.Length, "l", 32, "Password Length")
	flag.Parse()
	return opts
}

func getPassword(text string) string {
	nums := getMD5Hash(text)
	numList := hexToIntArray(nums)
	password := ""
	lastIdx2 := 0
	for _key, num := range numList {
		if (_key < 3) {
			//确保数字 小写 大写英文存在
			idx1 := _key
			idx2 := num % len(keyMap[idx1])
			password += keyMap[idx1][idx2:idx2+1]
		} else {
			idx1 := _key % len(keyMap)
			idx2 := (num + lastIdx2) % len(keyMap[idx1])
			lastIdx2 = idx2
			password += keyMap[idx1][idx2:idx2+1]
		}
	}
	return password
}

func hexToIntArray(hexstr string) (res [32]int) {
	bi := big.NewInt(0)
	for key, num := range hexstr {
		bi.SetString(string(num), 16)
		res[key], _ = strconv.Atoi(bi.String())
	}
	return
}

func getMD5Hash(text string) (string) {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
