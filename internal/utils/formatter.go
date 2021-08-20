package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/runes"
)

// Format a valid zip code, eg: "12345678" -> "12345-678"
func ZipCodeFormatter(zipCode string) string {
	zipCodeSearch := zipCode
	if len(zipCode) == 8 {
		zipCodeSearch = zipCode[0:5] + "-" + zipCode[5:8]
	}
	return zipCodeSearch
}

func Normalize(s string) string  {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return strings.ToLower(output)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RemoveAllNonNumericCharacteres(text string) string{
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(text, "")
	return processedString
}

func RemoveAllNonAlphanumericCharacteres(text string) string{
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(text, "")
	return processedString
}

func DateParsed(dateRaw string) (string, string, string) {
	t, err := time.Parse("02/01/2006", dateRaw)
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	year, month, day := t.Date()
	fmt.Printf("|%06d|%6d|\n", 12, 345)
	return PadZeros(year,4), PadZeros(int(month),2), PadZeros(day,2)
}

func PadZeros(number int, pos int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(pos)+"d", number)
}

func PadSpaces(number int, pos int) string {
	return fmt.Sprintf("%"+strconv.Itoa(pos)+"d", number)
}
















