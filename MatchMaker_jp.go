package main

import "fmt" // fmt.Println --> allows to print text
import "bufio" // provides input/output operations
import "strings" // allows to manipulate strings
import "os" // allows for interacting with the operating system
import "strconv" // converts strings to whatever else is needed

func validate() (bool, int) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter a number between 1 and 5.")
    userResponse, _ := reader.ReadString('\n')
    userResponse = strings.TrimSpace(userResponse)
    fmt.Println("You entered:", userResponse)
        
    userResponse2, err := strconv.Atoi(userResponse)
    if err != nil {
        fmt.Println("This is not a number!")
        return false, 0
    }
    if userResponse2 < 1 || userResponse2 > 5 {
        fmt.Println("\nThis is not a number between 1 and 5.")
        return false, 0
    }
    fmt.Println("\nThis is a number between 1 and 5. Great job!")
    return true, userResponse2
}
    
func determineRangeCompatibility(userTotalCompatibility int) {
	if userTotalCompatibility >= 95 {
		fmt.Println("Okay... I guess you can call me.")
	} else if userTotalCompatibility >= 75 {
		fmt.Println("Uh... friends?")
	} else {
		fmt.Println("Ugh, as if!")
	}
}

func main() {
    introduction := `
*****************************************************
                Justina's Matchmaker
        Are you worthy of being called mine?
                WinningMyLoveSoft, Inc.
*****************************************************       

This program figures out if you and I are meant to be.
You will answer 5 questions. To each question, answer 5
if you strongly agree, 4 is you agree, 3 if you neither
agree nor disagree, 2 if you disagree, and 1 if you 
strongly disagree.

Our happiness depends on you. Don't let us down...

If you do let us down, do not come running back...
my love will not be reciprocated.

*****************************************************
`
    questions := []string {
        "People should strive to eat MORE meat.", //1
        "Traveling is a MUST.", //5
        "EDM music is the best music genre.", //3
        "You should buy Justina a Birkin Bag.", //5
        "Social media is the best thing invented.", //3
    }

    desired_responses := []int {
        1, // strongly disagree
        5, // strongly agree
        3, // neither agree or disagree
        5, // strongly agree
        3, // neither agree or disagree
    }

    var max_score = 5 * len(questions)

    fmt.Println(introduction)

    var compatibility []int
	totalCompatibility := 0

	for i, question := range questions {
		var userResponse int
		var userValid bool
		fmt.Println(question)
		for !userValid {
			userValid, userResponse = validate()
		}

		questionCompatibility := 5 - absoluteValue(userResponse-desired_responses[i])
		compatibility = append(compatibility, questionCompatibility)

		fmt.Printf("Question %d compatibility: %d\n\n", i+1, questionCompatibility)
	}

	for _, c := range compatibility {
		totalCompatibility += c
	}

	totalCompatibility = totalCompatibility * 100 / max_score
	fmt.Printf("Total Compatibility: %d\n\n", totalCompatibility)

	determineRangeCompatibility(totalCompatibility)
}

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
