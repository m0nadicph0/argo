package parser

import (
	"bufio"
	"os"
)

type Parser interface {
	GetTokens() ([]string, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p parser) GetTokens() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	tokens := make([]string, 0)
	for scanner.Scan() {
		token := scanner.Text()
		tokens = append(tokens, token)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return tokens, nil
}
