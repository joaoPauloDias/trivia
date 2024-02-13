# Trivia


## API externa utilizada:
- https://opentdb.com/api_category.php


## API:

Em resumo, esta API serve como um intermediário personalizado entre o usuário e uma API externa de perguntas de trivia, permitindo buscar perguntas por quantidade e categoria, e retornando-as em um formato fácil de consumir.

**Build:**
  
```docker build -t mconf/api:candidato .```

**Run:**

```docker run -ti --rm -e API_PORT=3000 -p 3000:3000  mconf/api:candidato```


## RUNNER:

**Build:**

```docker build -t mconf/runner:candidato .```

**Run:**

Recebe como argumentos respectivamente número de questões (inteiro positivo) e categoria das questões (string)

```docker run -ti --rm -e API_PORT=3000  mconf/runner:candidato 1 "History"```

<details>

<summary>Categorias Disponíveis</summary>

- General Knowledge
- Entertainment: Books
- Entertainment: Film
- Entertainment: Music
- Entertainment: Musicals & Theatres
- Entertainment: Television
- Entertainment: Video Games
- Entertainment: Board Games
- Science & Nature
- Science: Computers
- Science: Mathematics
- Mythology
- Sports
- Geography
- History
- Politics
- Art
- Celebrities
- Animals
- Vehicles
- Entertainment: Comics
- Science: Gadgets
- Entertainment: Japanese Anime & Manga
- Entertainment: Cartoon & Animations
  
</details>



