package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"time"

	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"

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

// TODO: Check if answer can take multiple values
// TODO: Add timer with Huh
func main() {
	api_key, category, difficulty, que_limit, time_limit := read_arguments()
	questions := get_questions(api_key, category, difficulty, que_limit)
	begin_quiz(questions, que_limit, time_limit)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// -a API_KEY [-c Category] [-d Difficulty] [-limit Questions] [-time Time]
func read_arguments() (string, string, string, int, int) {
	api_key := flag.String("a", Config("API_KEY"), "Your QuizAPI key")
	category := flag.String("c", "", "Specify a category (Linux, DevOps, Networking, Programming, Cloud, Docker, Kubernetes)")
	difficulty := flag.String("d", "", "Specify a difficulty (Easy, Medium, Hard)")
	timer := flag.Int("time", 30, "Specify the time limit for the quiz in seconds")
	limit := flag.Int("limit", 1, "Specify the number of questions")

	flag.Parse()
	return *api_key, *category, *difficulty, *limit, *timer
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

func begin_quiz(questions []QuizResponse, limit int, time_limit int) {
	timer := time.NewTimer(time.Duration(time_limit) * time.Second)
	total_score := 0
	_ = timer
	var opt string

	formGroups := make([]*huh.Group, limit)
	for i, question := range questions {
		options_list := get_options(question.Answers)
		options := make([]huh.Option[string], 0, len(options_list))

		for key, option := range options_list {
			options = append(options, huh.NewOption(option, key))
		}

		formGroups[i] = huh.NewGroup(
			huh.NewSelect[string]().
				Title(question.Question).
				Options(options...).
				Value(&opt),
		)

		// Run the form for the current question
		form := huh.NewForm(formGroups[i])
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		key := fmt.Sprintf("answer_%s_correct", opt)
		if question.CorrectAnswers[key] == "true" {
			total_score += 1
		}
	}
	fmt.Printf("\nYou scored %d out of %d.", total_score, limit)

	// for _, question := range questions {
	// 	fmt.Println("Question: ", question.Question)
	// 	get_options(question.Answers)
	// 	score, err := check_answers_value(question.CorrectAnswers, timer)

	// 	if err != nil {
	// 		exit(fmt.Sprintf("You scored %d out of %d.", total_score, limit))
	// 	}
	// 	total_score += score
	// }
}

func get_options(answers map[string]string) map[string]string {
	options := make(map[string]string)

	for r := 'a'; r < 'g'; r++ {
		key := fmt.Sprintf("answer_%c", r)
		value := answers[key]
		if value != "" {
			options[fmt.Sprintf("%c", r)] = value
		}
	}
	return options
}

func check_answers_value(correct_answers map[string]string, timer *time.Timer) (int, error) {
	answerCh := make(chan string)
	fmt.Print("Pick an option: ")
	score := 0

	go func() {
		var input string
		fmt.Scanf("%s", &input)
		answerCh <- input
	}()

	select {
	case <-timer.C:
		return 0, fmt.Errorf("time's up")
	case input := <-answerCh:
		key := fmt.Sprintf("answer_%c_correct", input[0])
		if correct_answers[key] == "true" {
			score = 1
		}
	}

	return score, nil
}
