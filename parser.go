package gext

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

var (
	urlpath  string
	pageVars map[string]interface{}
)

// parser represents a parser.
type parser struct {
	s   *scanner
	buf struct {
		tok token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// newParser returns a new instance of parser.
func newParser(r io.Reader) *parser {
	return &parser{s: newScanner(r)}
}

// Parse parses a template statement.
func (p *parser) Parse(path string) (*page, error) {
	page := &page{}
	page.id = path
	var parent component
	parent = page

	pageVars = make(map[string]interface{})
	var bufRaw bytes.Buffer
	isNew := false
LoopBase:
	for {
		tok, lit := p.scan()
		//log.Println("tok=", tok, " lit=", lit)
		switch tok {
		case bRACEOPEN:
			tok2, lit2 := p.scan()
			if tok2 == pERCENT {
				isNew = true
				if bufRaw.Len() > 0 {
					cp := newRaw(bufRaw.String())
					appendChildren(parent, cp)
					bufRaw.Reset()
				}
			} else {
				bufRaw.WriteString(lit)
				bufRaw.WriteString(lit2)
			}
		case eXTENDS:
			if isNew {
				cp := newMaster(p.parseParameters())
				page.MasterPage = cp.GetID()
				//log.Println("master=", len(cp.GetChildren()))
				isNew = false
			} else {
				bufRaw.WriteString(lit)
			}
		case iNCLUDE:
			if isNew {
				cp := p.parseInclude()
				cp.Base().parent = page
				appendChildren(parent, cp)
				//log.Println("master=", len(cp.GetChildren()))
				isNew = false
			} else {
				bufRaw.WriteString(lit)
			}
		case bLOCK:
			if isNew {
				cp := newBlock(p.parseParameters())
				cp.Base().parent = page
				appendChildren(parent, cp)
				isNew = false
				parent = cp

			} else {
				bufRaw.WriteString(lit)
			}
		case eNDBLOCK:
			if isNew {
				parent = page
				p.scanToEndCode()
				isNew = false
			} else {
				bufRaw.WriteString(lit)
			}

		case eof:
			break LoopBase
		default:
			if !isNew {
				bufRaw.WriteString(lit)
			}

		}
	}
	if bufRaw.Len() > 0 {
		cp := newRaw(bufRaw.String())
		appendChildren(page, cp)
		bufRaw.Reset()
	}
	page.pageVarsInit = pageVars
	return page, nil
}

func appendChildren(parent component, child component) {
	parent.AppendChild(child)
}

func (p *parser) parseInclude() component {
	filename := p.parseQuotedStringWS()
	cp := newInclude(filename)
	p.scanToEndCode()
	return cp
}

func (p *parser) parseController() string {
	filename := p.parseQuotedStringWS()
	p.scanToEndCode()
	return filename
}

func (p *parser) parseParameters() map[string]interface{} {
	ret := make(map[string]interface{})
	varname := ""
Loop:
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case eQUAL:
			varvalue := p.parseQuotedString()
			ret[varname] = varvalue
		case pERCENT:
			tok2, _ := p.scanIgnoreWhitespace()
			if tok2 == bRACECLOSE {
				break Loop
			}
		case bRACECLOSE:
			break Loop
		default:
			varname = lit
		}
	}
	return ret
}

func (p *parser) parseParam() string {

	tok, _ := p.scanIgnoreWhitespace()
	if tok == eQUAL {
		_, l := p.scanIgnoreWhitespace()
		return l
	}
	return ""
}

func (p *parser) parseParamString() string {

	tok, _ := p.scanIgnoreWhitespace()
	if tok == eQUAL {
		return p.parseQuotedString()
	}
	return ""
}

func (p *parser) parseQuotedString() string {
	var buffer bytes.Buffer
	init := false
	for {
		tok, lit := p.scan()
		if tok != qUOTE {
			buffer.WriteString(lit)
			init = true
		} else {
			if init {
				break
			}
		}
	}
	return buffer.String()
}

func (p *parser) parseQuotedStringWS() string {
	var buffer bytes.Buffer
	init := false
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if tok != qUOTE {
			buffer.WriteString(lit)
			init = true
		} else {
			if init {
				break
			}
		}
	}
	return buffer.String()
}

func (p *parser) parseParamInt() (int, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != eQUAL {
		return 0, fmt.Errorf("found %q, expected = ", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	i, err := strconv.Atoi(lit)
	if err != nil {
		return 0, fmt.Errorf("found %q, expected integer number ", lit)
	}
	return i, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *parser) scan() (tok token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *parser) scanIgnoreWhitespace() (tok token, lit string) {
	tok, lit = p.scan()
	if tok == ws {
		tok, lit = p.scan()
	}
	return
}

func (p *parser) scanEndLine() (lit string) {
	return p.s.scanLine()
}

func (p *parser) scanToEndCode() {
	for {
		tok, _ := p.scan()
		if tok == pERCENT {
			tok2, _ := p.scan()
			if tok2 == bRACECLOSE {
				break
			}
		}
		if tok == eof {
			break
		}
	}

}

// unscan pushes the previously read token back onto the buffer.
func (p *parser) unscan() { p.buf.n = 1 }

func joinMap(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
