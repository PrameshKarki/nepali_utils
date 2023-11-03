package main

import nepali_utils "github.com/PrameshKarki/nepali_utils/lib"

func main() {
	nepali,_:=nepali_utils.EnglishToNepali("Mero Nepal") // Result: मेरो नेपाल
	currencyInWord:=nepali_utils.NumberToNepaliWord(12345) // Result: बाह्र हजार तीन सय पैंतालिस
}