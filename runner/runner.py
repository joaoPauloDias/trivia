import requests
import os

api_port = os.getenv('API_PORT', '8080')

base_url = "http://172.17.0.1:" + api_port

def get_home_page():
    url = base_url + "/home"
    response = requests.get(url)
    if response.status_code == 200:
        print("Success:")
        print(response.text)
    else:
        print("Error:", response.status_code)

def get_welcome_page():
    url = base_url + "/welcome"
    response = requests.get(url)
    if response.status_code == 200:
        print("Success:")
        print(response.text)
    else:
        print("Error:", response.status_code)

def get_about_page():
    url = base_url + "/about"
    response = requests.get(url)
    if response.status_code == 200:
        print("Success:")
        print(response.text)
    else:
        print("Error:", response.status_code)


def get_trivia_page():
    url = base_url + "/trivia"
    response = requests.get(url)
    if response.status_code == 200:
        print("Success:")
        print(response.text)
    else:
        print("Error:", response.status_code)

if __name__ == "__main__":
    while True:
        s = input("Enter 'home', 'welcome', 'about' to view pages or 'q' to quit: ")
        if s == 'q':
            break
        elif s == 'home':
            get_home_page()
        elif s == 'welcome':
            get_welcome_page()
        elif s == 'about':
            get_about_page()
        elif s == 'trivia': 
            get_trivia_page()
        else:
            print("Invalid input. Try again.")
