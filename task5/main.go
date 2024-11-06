package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var dictionary = []string{
	"автомобиль", "банан", "воздух", "город", "дерево", "еда", "футбол", "гора", "океан", "солнце",
	"луна", "звезда", "корабль", "самолет", "дом", "книга", "компьютер", "телефон", "часы", "метро",
	"парк", "река", "море", "озеро", "лес", "поле", "гора", "холм", "дорога", "мост", "здание",
	"окно", "дверь", "стена", "потолок", "пол", "стол", "стул", "кровать", "шкаф", "диван", "телевизор",
	"холодильник", "плита", "микроволновка", "чайник", "утюг", "пылесос", "фен", "бритва", "зубнаяпаста",
	"мыло", "шампунь", "гель", "крем", "маска", "помада", "ручка", "карандаш", "блокнот", "тетрадь",
	"книга", "журнал", "газета", "словарь", "энциклопедия", "карта", "глобус", "камера", "фотоаппарат",
	"фильм", "музыка", "песня", "барабан", "гитара", "пианино", "скрипка", "труба", "флейта", "саксофон",
	"балет", "опера", "театр", "кино", "музей", "библиотека", "университет", "школа", "детскийсад",
	"больница", "аптека", "магазин", "рынок", "банк", "аэропорт", "вокзал", "порт", "пристань", "мост",
	"туннель", "мост",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("Выберите вариант ответа:")
		fmt.Println("1. Новая игра")
		fmt.Println("2. Выйти из приложения")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" || input == "Новая игра" {
			playGame()
		} else if input == "2" || input == "Выйти из приложения" {
			break
		} else {
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1 или 2.")
		}
	}
}

func playGame() {
	word := dictionary[rand.Intn(len(dictionary))]
	guessed := make(map[rune]bool)
	attemptsLeft := 6
	currentWord := make([]rune, len(word))
	for i := range currentWord {
		currentWord[i] = '_'
	}

	reader := bufio.NewReader(os.Stdin)

	for attemptsLeft > 0 {
		displayHangman(attemptsLeft)
		displayWord(currentWord)

		fmt.Print("Введите букву: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) != 1 {
			fmt.Println("Пожалуйста, введите одну букву.")
			continue
		}

		letter := rune(input[0])
		if guessed[letter] {
			fmt.Println("Вы уже угадали эту букву.")
			continue
		}

		guessed[letter] = true
		if strings.ContainsRune(word, letter) {
			for i, l := range word {
				if l == letter {
					currentWord[i] = letter
				}
			}
		} else {
			attemptsLeft--
		}

		if isGameOver(currentWord) {
			fmt.Println("Поздравляем! Вы выиграли!")
			break
		}
	}

	if attemptsLeft == 0 {
		displayHangman(attemptsLeft)
		fmt.Println("Вы проиграли. Загаданное слово было:", word)
	}
}

func displayHangman(attemptsLeft int) {
	fmt.Printf("Осталось попыток: %d\n", attemptsLeft)
	
}

func displayWord(currentWord []rune) {
	for _, letter := range currentWord {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}

func isGameOver(currentWord []rune) bool {
	for _, letter := range currentWord {
		if letter == '_' {
			return false
		}
	}
	return true
}