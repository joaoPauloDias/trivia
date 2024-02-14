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

Também pode não receber nenhum argumento como parâmetro, nesse caso retorna um número arbitrário de questões de categorias aleatórias

```docker run -ti --rm -e API_PORT=3000  mconf/runner:candidato```

<details>

<summary>Categorias Disponíveis</summary>

- General Knowledge
- Books
- Film
- Music
- Theatre
- Television
- Video Games
- Board Games
- Nature
- Computers
- Mathematics
- Mythology
- Sports
- Geography
- History
- Politics
- Art
- Celebrities
- Animals
- Vehicles
- Comics
- Gadgets
- Anime
- Cartoons
  
</details>



