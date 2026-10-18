package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1
	golang := "GoLang"
	golang_low := strings.ToLower(golang)
	fmt.Println(golang_low)
	fmt.Println(strings.ToUpper(golang))
	// 2
	back := " backend developer "
	fmt.Println(strings.TrimSpace(back))
	// 3
	java := "I like Java"
	java_cont := strings.Contains(java, "Java")
	if java_cont {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}
	// 4
	text := "one,two,three"
	text_replace := strings.ReplaceAll(text, ",", ";")
	fmt.Println(text_replace)
	// 5
	var login string
	fmt.Println("Введите логин:")
	fmt.Scan(&login)
	login_adm := strings.HasPrefix(login, "admin")
	if login_adm {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
	// 6
	gogo := "Go is fast and Go is simple"
	gogo_count := strings.Count(gogo, "Go")
	fmt.Println(gogo_count)
	// 7
	var word string
	fmt.Println("Введите слово:")
	fmt.Scan(&word)
	word = strings.ToLower(word)
	letters := []rune(word)
	isPalindrome := true

	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}
	//8
	var word2 string
	fmt.Println("Введите строку:")
	fmt.Scanln(&word2)
	found := false

	for i := 0; i < len(word2)-2; i++ {
		if word2[i] == word2[i+1] && word2[i+1] == word2[i+2] {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Три одинаковых символа подряд найдены")
	} else {
		fmt.Println("Три одинаковых символа подряд не найдены")
	}
	//9
	var password string
	fmt.Println("Введите пароль:")
	fmt.Scan(&password)

	hasDigit := false
	hasUpper := false

	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
		}

		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
	}
	if len([]rune(password)) < 8 {
		fmt.Println("Слишком короткий пароль")
	} else if !hasDigit {
		fmt.Println("Пароль должен содержать цифру")
	} else if !hasUpper {
		fmt.Println("Пароль должен содержать заглавную букву")
	} else {
		fmt.Println("Пароль корректный")
	}
	// 10
	var email string
	fmt.Println("Введите email:")
	fmt.Scan(&email)

	atCount := strings.Count(email, "@")
	atIndex := strings.Index(email, "@")

	if atCount == 0 {
		fmt.Println("Email должен содержать @")
	} else if atCount > 1 {
		fmt.Println("Email должен содержать ровно один @")
	} else if !strings.Contains(email[atIndex+1:], ".") {
		fmt.Println("Email должен содержать точку после @")
	} else {
		fmt.Println("Email корректен")
	}
}
