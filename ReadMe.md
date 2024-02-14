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

## Exemplo de interação do usuário

Ao rodar o script do runner, tendo em vista que a api está rodando corretamente em seu container e o runner recebeu como parâmetro a mesma porta que a api está escutando, espera-se receber da API uma série de perguntas (a quantidade pode ser definida pelo usuário). As perguntas são exibidas de maneira sequencial pro usuário que deve dar como input a alternativa que julga correta. Ao final de todas perguntas o score do usuário é exibido, perguntas mais difíceis valem mais pontos.

```
sudo docker run -ti --rm -e API_PORT=4000  mconf/runner:candidato 1 "History"

Category: History - Points: 1
Question: Which one of these was not a beach landing site in the Invasion of Normandy?

1. Silver
2. Sword
3. Juno
4. Gold

Enter the number of the correct answer: 2
Incorrect! - The correct answer is: Silver


You scored 0 points!```
