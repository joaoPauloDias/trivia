package main

// mapDifficultyToPoints maps the difficulty level of each question to a corresponding point value.
func mapDifficultyToPoints(questions []Question) {
	for i, question := range questions {
		switch question.Difficulty {
		case "easy":
			questions[i].Points = 1
		case "medium":
			questions[i].Points = 2
		case "hard":
			questions[i].Points = 3
		default:
			questions[i].Points = 0
		}
	}
}
