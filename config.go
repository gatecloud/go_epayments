package epayments

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	validator "gopkg.in/go-playground/validator.v8"
)

type Config struct {
	SignKey  string `validate:"required"`
	Endpoint string `validate:"required"`
}

func (e *Config) Sign(sig Signaturer) (int, error) {
	var (
		keys       []string
		paymentMap = make(map[string]string)
		err        error
	)

	// Validate if the fields are null
	config := validator.Config{
		TagName: "validate",
	}

	validate := validator.New(&config)
	if err := validate.Struct(*e); err != nil {
		return http.StatusBadRequest, err
	}

	// Sort the fields
	if err := scan(sig, paymentMap); err != nil {
		return http.StatusInternalServerError, err
	}

	for k := range paymentMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	sortedQueryString := ""
	for _, key := range keys {
		sortedQueryString += "&"
		sortedQueryString += key
		sortedQueryString += "="
		sortedQueryString += paymentMap[key]
	}
	sortedQueryString += e.SignKey

	// Generate the signature
	fmt.Println("===", sortedQueryString[1:len(sortedQueryString)])
	h := md5.New()
	n, err := io.WriteString(h, sortedQueryString[1:len(sortedQueryString)])
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if n < 0 {
		return http.StatusInternalServerError, errors.New("Write <= bytes")
	}

	sig.SetSignature(fmt.Sprintf("%x", h.Sum(nil)))
	sig.SetSignType("MD5")

	if err := validate.Struct(sig); err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (e *Config) Verify(ver Verifier) (int, error) {
	var (
		keys       []string
		paymentMap = make(map[string]string)
		err        error
	)

	// Sort the fields
	if err := scan(ver, paymentMap); err != nil {
		return http.StatusInternalServerError, err
	}

	for k := range paymentMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	sortedQueryString := ""
	for _, key := range keys {
		sortedQueryString += "&"
		sortedQueryString += key
		sortedQueryString += "="
		sortedQueryString += paymentMap[key]
	}
	sortedQueryString += e.SignKey

	// Generate the signature
	fmt.Println("===", sortedQueryString[1:len(sortedQueryString)])
	h := md5.New()
	n, err := io.WriteString(h, sortedQueryString[1:len(sortedQueryString)])
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if n < 0 {
		return http.StatusInternalServerError, errors.New("Write <= bytes")
	}

	signature := fmt.Sprintf("%x", h.Sum(nil))
	if ver.GetSignature() != signature {
		return http.StatusBadRequest, errors.New("Signature not match")
	}

	return http.StatusOK, nil
}

func ToURLParams(sig Signaturer) (string, error) {
	rfPayment := reflect.ValueOf(sig).Elem()
	if !rfPayment.IsValid() {
		return "", errors.New("reflect error")
	}

	v := url.Values{}
	for i := 0; i < rfPayment.NumField(); i++ {
		tag := rfPayment.Type().Field(i).Tag.Get("json")
		p := rfPayment.Field(i)
		if reflect.DeepEqual(p.Interface(), reflect.Zero(reflect.TypeOf(p.Interface())).Interface()) {
			continue
		}

		switch p.Kind() {
		case reflect.String:
			s := rfPayment.Field(i).String()
			v.Set(tag, s)
		case reflect.Float64:
			f := rfPayment.Field(i).Float()
			v.Set(tag, fmt.Sprintf("%.2f", f))
		case reflect.Int64:
			f := rfPayment.Field(i).Int()
			v.Set(tag, fmt.Sprintf("%d", f))
		}
	}
	return v.Encode(), nil
}

func EncodeSpecialChar(src string) string {
	dst := url.QueryEscape(src)
	dst = strings.Replace(dst, "+", "%20", -1)
	dst = strings.Replace(dst, "*", "%2A", -1)
	return strings.Replace(dst, "%7E", "~", -1)
}

func scan(i interface{}, m map[string]string) error {
	val := reflect.ValueOf(i)
	if !val.IsValid() {
		return errors.New("reflect error")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("json")
		if tag != "signature" && tag != "sign_type" {
			p := val.Field(i)
			// Check if the value is 0, null or ""
			if reflect.DeepEqual(p.Interface(), reflect.Zero(reflect.TypeOf(p.Interface())).Interface()) {
				continue
			}
			fmt.Println("=====", tag)

			switch p.Kind() {
			case reflect.String:
				s := val.Field(i).String()
				m[tag] = s
			case reflect.Float64:
				f := val.Field(i).Float()
				m[tag] = fmt.Sprintf("%.2f", f)
			case reflect.Int64:
				f := val.Field(i).Int()
				m[tag] = fmt.Sprintf("%d", f)
			case reflect.Struct:
				if err := scan(p, m); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
