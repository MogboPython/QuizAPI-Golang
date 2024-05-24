package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"unicode"

	"github.com/joho/godotenv"
)

type QuizResponse struct {
	ID                     int               `json:"id"`
	Question               string            `json:"question"`
	Description            string            `json:"description"`
	Answers                map[string]string `json:"answers"`
	MultipleCorrectAnswers string            `json:"multiple_correct_answers"`
	CorrectAnswers         map[string]string `json:"correct_answers"`
	Explanation            string            `json:"explanation"`
	Tip                    any               `json:"tip"`
	Tags                   []any             `json:"tags"`
	Category               string            `json:"category"`
	Difficulty             string            `json:"difficulty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	api_key, category, difficulty, limit := read_arguments()
	questions := get_questions(api_key, category, difficulty, limit)
	score := begin_quiz(questions)
	fmt.Printf("You scored %d out of %d.", score, limit)
}

// -a API_KEY [-c Category] [-d Difficulty]
func read_arguments() (string, string, string, int) {
	// TODO: Remove default default API Key
	api_key := flag.String("a", Config("API_KEY"), "Your QuizAPI key")
	category := flag.String("c", "", "Specify a category (Linux, DevOps, Networking, Programming, Cloud, Docker, Kubernetes)")
	difficulty := flag.String("d", "", "Specify a difficulty (Easy, Medium, Hard)")
	limit := flag.Int("limit", 1, "Specify the number of questions")
	flag.Parse()
	return *api_key, *category, *difficulty, *limit
}

func begin_quiz(questions []QuizResponse) int {
	total_score := 0

	for _, question := range questions {
		fmt.Println("Question: ", question.Question)
		show_options(question.Answers)
		total_score += check_answers_value(question.CorrectAnswers)
	}
	return total_score
}

func get_questions(api_key string, category string, difficulty string, limit int) []QuizResponse {
	request, err := http.NewRequest("GET", "https://quizapi.io/api/v1/questions", nil)

	if err != nil {
		exit(fmt.Sprintf("Error in making request: %s", err))
	}

	request.Header.Set("X-Api-Key", api_key)

	q := request.URL.Query()
	q.Add("category", category)
	q.Add("difficulty", difficulty)
	q.Add("limit", strconv.Itoa(limit))
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		exit(fmt.Sprintf("Error in making request: %s", err.Error()))
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		exit(fmt.Sprintf("Error in reading request  body: %s", err))
	}

	var questions []QuizResponse
	err = json.Unmarshal(responseData, &questions)

	if err != nil {
		var errorResponse ErrorResponse
		err = json.Unmarshal([]byte(responseData), &errorResponse)
		if err != nil {
			fmt.Println("Error:", err)
		}
		exit(fmt.Sprintf("Error: %s", errorResponse.Error))
	}

	if len(questions) == 0 {
		exit("No questions found... Try with different category or tag.")
	}

	return questions

}

func get_questions_struct(responseData []byte) []QuizResponse {
	var questions []QuizResponse
	err := json.Unmarshal(responseData, &questions)

	if err != nil {
		var errorResponse ErrorResponse
		err = json.Unmarshal([]byte(responseData), &errorResponse)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Error:", errorResponse.Error)
	}
	return questions
}

func show_options(answers map[string]string) {
	println("Enter the correct option: ")

	for r := 'a'; r < 'g'; r++ {
		key := fmt.Sprintf("answer_%c", r)
		value := answers[key]
		if value != "" {
			fmt.Printf("%c) %s\n", r, value)
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func check_answers_value(correct_answers map[string]string) int {
	var input string
	fmt.Print("Pick an option: ")
	_, err := fmt.Scanf("%s", &input)

	if err != nil {
		exit(fmt.Sprintf("Error reading input: %v", err))
	}

	if len(input) != 1 || !unicode.IsLetter(rune(input[0])) {
		exit("Invalid input.")
	}

	// fmt.Printf("You entered: %c\n", input[0])
	score := 0
	key := fmt.Sprintf("answer_%c_correct", input[0])
	if correct_answers[key] == "true" {
		score = 1
	}

	return score
	// else if correct_answers[key] == "false" {
	// 	fmt.Println("better luck next time")
	// } else {
	// 	fmt.Println("option not available")
	// }
}

// func handleResponse(responseData string) []QuizResponse {
// 	var response []QuizResponse

// 	err := json.Unmarshal([]byte(responseData), &response)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	return response
// }
