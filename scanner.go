package gext

import (
	"bufio"
	"bytes"
	"io"
)

// scanner represents a lexical scanner.
type scanner struct {
	r *bufio.Reader
}

// newScanner returns a new instance of scanner.
func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *scanner) Scan() (tok token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	// if isWhitespace(ch) {
	// 	s.unread()
	// 	return s.scanWhitespace()
	// } else
	if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isDigit(ch) {
		s.unread()
		return s.scanIdent()
	}
	// Otherwise read the individual character.
	switch ch {
	case endOfFile:
		return eof, ""
	case ' ':
		return ws, string(ch)
	case '*':
		return aSTERISK, string(ch)
	case ',':
		return cOMMA, string(ch)
	case '(':
		return pARENTHESISON, string(ch)
	case ')':
		return pARENTHESISOFF, string(ch)
	case '=':
		return eQUAL, string(ch)
	case '"':
		return qUOTE, string(ch)
	case '{':
		return bRACEOPEN, string(ch)
	case '}':
		return bRACECLOSE, string(ch)
	case '[':
		return bRAKETOPEN, string(ch)
	case ']':
		return bRAKETCLOSE, string(ch)
	case '\'':
		return sIMPLEQUOTE, string(ch)
	case '%':
		return pERCENT, string(ch)
	case '@':
		return aT, string(ch)
	}

	return illegal, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (tok token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != '@' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return ws, buf.String()
}

func (s *scanner) scanLine() (lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if isEndLine(ch) {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *scanner) scanIdent() (tok token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != '@' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch buf.String() {

	case "id":
		return iD, buf.String()
	case "text":
		return tEXT, buf.String()
	case "block":
		return bLOCK, buf.String()
	case "endblock":
		return eNDBLOCK, buf.String()
	case "extends":
		return eXTENDS, buf.String()
	case "include":
		return iNCLUDE, buf.String()
	case "cache":
		return cACHE, buf.String()
	case "field":
		return fIELD, buf.String()
	case "fieldvalue":
		return fIELDVALUE, buf.String()
	case "fieldtext":
		return fIELDTEXT, buf.String()
	case "selected":
		return sELECTED, buf.String()
	case "value":
		return vALUE, buf.String()
	case "class":
		return cLASS, buf.String()
	case "textbox":
		return tEXTBOX, buf.String()
	case "textmode":
		return tEXTMODE, buf.String()
	case "security":
		return sECURITY, buf.String()
	case "endsecurity":
		return eNDSECURITY, buf.String()
	case "methods":
		return mETHODS, buf.String()
	case "roles":
		return rOLES, buf.String()
	}

	// Otherwise return as a regular identifier.
	return ident, buf.String()
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return endOfFile
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

func isOpenBrace(ch rune) bool { return ch == '{' }

func isCloseBrace(ch rune) bool { return ch == '}' }

func isPercent(ch rune) bool { return ch == '%' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isEndLine(ch rune) bool { return ch == '\n' || ch == '\r' }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var endOfFile = rune(0)
