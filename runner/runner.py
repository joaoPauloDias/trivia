import requests
import os
import html
import random
import sys

api_port = os.getenv('API_PORT', '3000')
base_url = f"http://172.17.0.1:{api_port}"

def deconstruct_json(data):
    """
    Deconstructs the JSON data received from the API into a simplified format.

    Args:
        data (list): The JSON data to be deconstructed.

    Returns:
        list: The deconstructed data in a simplified format.
    """
    simplified_data = [
        {
            "Category": html.unescape(item["category"]),
            "Difficulty": html.unescape(item["difficulty"]),
            "Question": html.unescape(item["question"]),
            "Correct Answer": html.unescape(item["correct_answer"]),
            "Incorrect Answers": [html.unescape(answer) for answer in item["incorrect_answers"]],
            "Points": item["points"]
        }
        for item in data
    ]
    return simplified_data

def get_trivia_page(amount, category):
    """
    Retrieves a trivia page from the API.

    Args:
        amount (int): The number of trivia questions to retrieve.
        category (str): The category of the trivia questions.

    Returns:
        dict: The JSON response containing the trivia questions.
    """
    url = f"{base_url}/trivia"
    if str(amount).isdigit() and category:
        url += f"?amount={amount}&category={category}"
    try:
        response = requests.get(url)
        response.raise_for_status()
        return response.json()
    except ValueError as e:
        print("Error:", e)

if __name__ == "__main__":
    
    if len(sys.argv) == 3:
        data = deconstruct_json(get_trivia_page(sys.argv[1] , sys.argv[2]))
    else:
        data = deconstruct_json(get_trivia_page(None, None))
    
    points = 0
    for item in data:
        print(f'\nCategory: {item["Category"]} - Points: {item["Points"]}')
        print(f'Question: {item["Question"]}\n')
        questions = item["Incorrect Answers"] + [item["Correct Answer"]]
        random.shuffle(questions)
        for i, answer in enumerate(questions):
            print(f'{i+1}. {answer}')
        answer = input("\nEnter the number of the correct answer: ")
        if questions[int(answer)-1] == item["Correct Answer"]:
            points += item["Points"]
            print("Correct!\n")
        else:
            print(f'Incorrect! - The correct answer is: {item["Correct Answer"]}\n')
    
    print(f'\nYou scored {points} points!')
