package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

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

func TestGetQuestionsStruct(t *testing.T) {
	// Define test data
	successfulResponse := []byte(`[{"id":1,"question":"How to delete a directory in Linux?","description":"delete folder","answers":{"answer_a":"ls","answer_b":"delete","answer_c":"remove","answer_d":"rmdir","answer_e":null,"answer_f":null},"multiple_correct_answers":"false","correct_answers":{"answer_a_correct":"false","answer_b_correct":"false","answer_c_correct":"false","answer_d_correct":"true","answer_e_correct":"false","answer_f_correct":"false"},"explanation":"rmdir deletes an empty directory","tip":null,"tags":[],"category":"linux","difficulty":"Easy"}]`)

	unsuccessfulResponse := `{"error":"Unauthenticated"}`
	_ = unsuccessfulResponse

	// Call the function being tested
	result := get_questions_struct([]byte(successfulResponse))

	// Define the expected result
	expected := []QuizResponse{
		{
			ID:          1,
			Question:    "How to delete a directory in Linux?",
			Description: "delete folder",
			Answers: map[string]string{
				"answer_a": "ls",
				"answer_b": "delete",
				"answer_c": "remove",
				"answer_d": "rmdir",
				"answer_e": "",
				"answer_f": "",
			},
			MultipleCorrectAnswers: "false",
			CorrectAnswers: map[string]string{
				"answer_a_correct": "false",
				"answer_b_correct": "false",
				"answer_c_correct": "false",
				"answer_d_correct": "true",
				"answer_e_correct": "false",
				"answer_f_correct": "false",
			},
			Explanation: "rmdir deletes an empty directory",
			Tip:         "",
			Tags:        []any{},
			Category:    "linux",
			Difficulty:  "Easy",
		},
	}

	// expected := ErrorResponse{
	// 		Error: "Unauthenticated",
	// }

	// Check if the actual result matches the expected result
	if len(result) != len(expected) {
		t.Errorf("Length mismatch. Expected: %d, Got: %d", len(expected), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i].Question != expected[i].Question {
			t.Errorf("Mismatch at index %d. Expected: %+v, Got: %+v", i, expected[i], result[i])
		}
	}
}

// Mock input and simulate the timer
func TestCheckAnswersValue(t *testing.T) {
	// Create a correct answers map
	correctAnswers := map[string]string{
		"answer_a_correct": "true",
		"answer_b_correct": "false",
	}

	// Create a timer
	timer := time.NewTimer(5 * time.Second)

	// Redirect input
	input := "a\n"
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	// Call the function
	score, err := check_answers_value(correctAnswers, timer)
	// Check for errors
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check the score
	expectedScore := 1
	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}
