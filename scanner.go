package main

import "fmt"

type scanner struct {
	src    string
	start  int
	end    int
	line   int
	tokens []*Token
}

func newScanner(src string) *scanner {
	return &scanner{src, 0, 0, 1, make([]*Token, 0)}
}

func (s *scanner) scan() error {
	for s.end < len(s.src) {
		s.start = s.end
		if err := s.scanToken(); err != nil {
			return err
		}
	}
	return nil
}

func (s *scanner) scanToken() error {
	switch c := s.advance(); c {
	case '(':
		s.addToken(TokLeftParen)
	case ')':
		s.addToken(TokRightParen)
	case '{':
		s.addToken(TokLeftBrace)
	case '}':
		s.addToken(TokRightBrace)
	case ',':
		s.addToken(TokComma)
	case '.':
		s.addToken(TokDot)
	case '-':
		s.addToken(TokMinus)
	case '+':
		s.addToken(TokPlus)
	case ';':
		s.addToken(TokSemicolon)
	case '/':
		s.addToken(TokSlash)
	case '*':
		s.addToken(TokStar)
	case ' ', '\r', '\t', '\n':
	default:
		return fmt.Errorf("invalid character: %s", string(c))
	}
	return nil
}

func (s *scanner) advance() byte {
	defer func() { s.end++ }()
	return s.currentByte()
}

func (s *scanner) addToken(tokType TokenType) {
	s.tokens = append(s.tokens, &Token{tokType, s.currentText()})
}

func (s *scanner) currentByte() byte {
	return s.src[s.end]
}

func (s *scanner) currentText() string {
	return s.src[s.start:s.end]
}
