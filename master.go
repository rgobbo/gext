package gext

import "log"

func newMaster(args map[string]interface{}) *page {

	// file, err := os.Open("./" + PathTemplates + "/" + filename)
	// if err != nil {
	// 	log.Println("Master error:=", err)
	// 	return nil
	// }
	// defer file.Close()

	// parser := newParser(file)
	// p, err := parser.Parse("")
	// if err != nil {
	// 	log.Println("Master error:=", err)
	// 	return nil
	// }
	// p.id = filename
	// p.tagName = "masterpage"
	filename := args["file"].(string)
	//	name := strings.Replace(filename, ".html", "", 1)
	p, err := getPage(filename)
	if err != nil {
		log.Println("Master error:=", err)
		return nil
	}

	p.tagName = "masterpage"
	return p
}
