package gext

type token int

const (
	// Special tokens

	illegal token = iota
	eof
	ws

	// Literals

	ident // main
	rAW   // html code

	// Misc characters

	aSTERISK       // *
	cOMMA          // ,
	pARENTHESISON  //(
	pARENTHESISOFF //)
	eQUAL          //=
	qUOTE          // "
	bRACEOPEN      // {
	bRACECLOSE     //}
	pERCENT        //%
	bRAKETOPEN     //[
	bRAKETCLOSE    //]
	sIMPLEQUOTE    //'
	aT             //@

	// Keywords

	fIELDVALUE  //fieldvalue
	fIELDTEXT   //fieldtext
	sELECTED    //selected
	vALUE       //value
	fIELD       //field
	iD          //ID
	tEXT        // text
	bLOCK       //block
	eNDBLOCK    //endblock
	eXTENDS     //extends
	iNCLUDE     //include
	cACHE       //cache
	cLASS       //class
	tEXTBOX     //textbox
	tEXTMODE    //textmode
	sECURITY    //security
	eNDSECURITY //endsecurity
	mETHOD      // method



)
