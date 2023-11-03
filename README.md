## Nepali utilities 🇳🇵

It's a package that encompasses a wide range of common Nepali utilities in Golang, designed to assist fellow Nepali developers in their work.

#### Features:
- [✅] English to Nepali Converter
- [✅] Number to Nepali Word Converter
- [❎] Date Conversion
- Many yet to come

#### Examples:
```
package main

import nepali_utils "github.com/PrameshKarki/nepali_utils/lib"

func main() {
	nepali,_:=nepali_utils.EnglishToNepali("Mero Nepal") // Result: मेरो नेपाल
	currencyInWord:=nepali_utils.NumberToNepaliWord(12345) // Result: बाह्र हजार तीन सय पैंतालिस
}
```