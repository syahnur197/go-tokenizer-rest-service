# Installation Instruction
1. Run `go mod tidy` in both api and client directories
2. Run `go run .` in both api and client directories
3. `http://localhost:8000` is the API web server to tokenise and sort text based on the number of occurence
4. `http://localhost:8001` is the client which consumes the tokeniser API

# API Routes

**Route:** `GET http://localhost:8000/top-ten-words/` 

will return the top ten occuring words in text

### Params
text:`string` - The text that you want to tokenise and sort


# Client Routes

1. `http://localhost:8001/consumer-one/`
2. `http://localhost:8001/consumer-two/`
3. `http://localhost:8001/consumer-three/`

# TODO

## API Script

- [X] Tokenize text
- [X] Count tokens
- [x] Sort tokens
- [x] Return top ten tokens as JSON

## Client Script

- [x] Consume first script API using http client
- [x] Display the top ten as response in JSON
