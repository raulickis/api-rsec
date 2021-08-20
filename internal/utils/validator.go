package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsZipCode(doc string) bool {
	var zipCodeRgx = regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !zipCodeRgx.MatchString(doc) {
		return false
	}
	return true
}

func IsFederativeUnit(state string) bool {
	var federativeUnits = []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	return Contains(federativeUnits, state)
}

func FederativeUnit(stateFullName string) string {
	var states = map[string]string{
		"acre"               :"AC",  // "Acre"
		"alagoas"            :"AL",  // "Alagoas"
		"amapa"              :"AP",  // "Amapá"
		"amazonas"           :"AM",  // "Amazonas"
		"bahia"              :"BA",  // "Bahia"
		"ceara"              :"CE",  // "Ceará"
		"distrito federal"   :"DF",  // "Distrito Federal"
		"espirito santo"     :"ES",  // "Espírito Santo"
		"goias"              :"GO",  // "Goiás"
		"maranhao"           :"MA",  // "Maranhão"
		"mato grosso"        :"MT",  // "Mato Grosso"
		"mato grosso do sul" :"MS",  // "Mato Grosso do Sul"
		"minas gerais"       :"MG",  // "Minas Gerais"
		"para"               :"PA",  // "Pará"
		"paraiba"            :"PB",  // "Paraíba"
		"parana"             :"PR",  // "Paraná"
		"pernambuco"         :"PE",  // "Pernambuco"
		"piaui"              :"PI",  // "Piauí"
		"rio de janeiro"     :"RJ",  // "Rio de Janeiro"
		"rio grande do norte":"RN",  // "Rio Grande do Norte"
		"rio grande do sul"  :"RS",  // "Rio Grande do Sul"
		"rondonia"           :"RO",  // "Rondônia"
		"roraima"            :"RR",  // "Roraima"
		"santa catarina"     :"SC",  // "Santa Catarina"
		"sao paulo"          :"SP",  // "São Paulo"
		"sergipe"            :"SE",  // "Sergipe"
		"tocantins"          :"TO"}  // "Tocantins"
	state, _ := states[stateFullName]
	return state
}

func IsIP4(ipAddress string) bool {
	var ipRgx = regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	ipAddress = strings.Trim(ipAddress, " ")
	if ipRgx.MatchString(ipAddress) {
		return true
	}
	return false
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	var emailRgx = regexp.MustCompile(".+@.+\\..+")
	if !emailRgx.MatchString(email) {
		return false
	}
	return true
}

func IsValidCPF(cpfFull string) bool {
	cpf := RemoveAllNonNumericCharacteres(cpfFull)
	//cpf = strings.Replace(cpf, ".", "", -1)
	//cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return false
	}

	return true
}

func IsValidCNPJ(cnpjFull string) bool {
	cnpj := RemoveAllNonNumericCharacteres(cnpjFull)
	//cnpj = strings.Replace(cnpj, ".", "", -1)
	//cnpj = strings.Replace(cnpj, "-", "", -1)
	//cnpj = strings.Replace(cnpj, "/", "", -1)
	if len(cnpj) != 14 {
		return false
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	char12, _ := strconv.Atoi(string(cnpj[12]))
	if char12 != digit1 {
		return false
	}
	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))

		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	char13, _ := strconv.Atoi(string(cnpj[13]))
	if char13 != digit2 {
		return false
	}

	return true
}

func IsValidDate(dateRaw string) bool {
	//var dateRgx = regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)") // dd/mm/yyyy
	_, err := time.Parse("02/01/2006", dateRaw)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
Examples for this validator in:
   https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835 (tutorial)
   https://godoc.org/gopkg.in/go-playground/validator.v9
   https://github.com/go-playground/universal-translator
*/
func ValidatorForm() (*validator.Validate, ut.Translator) {
	v := validator.New()

	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} é uma informação obrigatória", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("required_without", trans, func(ut ut.Translator) error {
		return ut.Add("required_without", "{0} é uma informação obrigatória", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("len", trans, func(ut ut.Translator) error {
		return ut.Add("len", "{0} ultrapassou o tamanho máximo permitido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("len", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} ultrapassou o tamanho máximo permitido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} possui menos dígitos do que o mínimo permitido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} precisa ser um email válido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("isvalidcnpj", trans, func(ut ut.Translator) error {
		return ut.Add("isvalidcnpj", "{0} inválido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvalidcnpj", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("isvalidcpf", trans, func(ut ut.Translator) error {
		return ut.Add("isvalidcpf", "{0} inválido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvalidcpf", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("isvalidcreditcard", trans, func(ut ut.Translator) error {
		return ut.Add("isvalidcreditcard", "{0} inválido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvalidcreditcard", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("isvaliddate", trans, func(ut ut.Translator) error {
		return ut.Add("isvaliddate", "{0} inválida", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvaliddate", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("isvalidzipcode", trans, func(ut ut.Translator) error {
		return ut.Add("isvalidzipcode", "{0} informado é inválido", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isvalidzipcode", fe.Field())
		return t
	})

	return v, trans
}

func IsValidCreditCardNumber(cardNumber string) bool {
	// TODO: Traduzir de Java para GOLANG, de maneira a ter uma validacao mais precisa: https://github.com/wirecardBrasil/credit-card-validator
	//var creditCardRgx = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	var creditCardRgx = regexp.MustCompile(`^(\d{13,19})$`)
	cardNumberRaw := RemoveAllNonNumericCharacteres(cardNumber)
	if !creditCardRgx.MatchString(cardNumberRaw) {
		return false
	}
	return true
}

