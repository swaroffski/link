package user

// Генерация логина по порядку от @a, @b, ..., до @z, @aa, @ab...
var currentLogin int = 0

// Генерация уникального логина
func GenerateLogin() string {
	// Логика генерации логинов
	login := toBase26(currentLogin)
	currentLogin++
	return "@" + login
}

// Функция преобразования числа в строку с буквами (например 0 -> a, 1 -> b, 26 -> aa)
func toBase26(num int) string {
	var result string
	for num >= 0 {
		result = string(rune('a'+num%26)) + result
		num = num/26 - 1
		if num < 0 {
			break
		}
	}
	return result
}

// Генерация ключа из словаря
func GenerateKey(words []string) string {
	key := words[0] + words[1] + words[2] + words[3]
	return key
}
