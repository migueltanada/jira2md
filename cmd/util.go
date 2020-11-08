package jira2md

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// ConvertAll function
func ConvertAll() {

	if len(os.Args) != 2 {
		log.Fatal(fmt.Errorf("This util can only accept exactly 1 parameters. \n \t1. File name/path to be converted to MD format"))
	}

	inputFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file \"%s\". %s", inputFile, err))
	}

	result := string(inputFile)

	ConvertHeadings(&result)
	ConvertCodeBlocks(&result)
	ConvertTextEffects(&result)

	fmt.Print(result)

}

// ConvertHeadings function to convert headings https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=headings
func ConvertHeadings(result *string) {

	header1Jira := regexp.MustCompile(`(?m)^h1\. `)
	header2Jira := regexp.MustCompile(`(?m)^h2\. `)
	header3Jira := regexp.MustCompile(`(?m)^h3\. `)
	header4Jira := regexp.MustCompile(`(?m)^h4\. `)
	header5Jira := regexp.MustCompile(`(?m)^h5\. `)
	header6Jira := regexp.MustCompile(`(?m)^h6\. `)

	*result = header1Jira.ReplaceAllString(*result, "# ")
	*result = header2Jira.ReplaceAllString(*result, "## ")
	*result = header3Jira.ReplaceAllString(*result, "### ")
	*result = header4Jira.ReplaceAllString(*result, "#### ")
	*result = header5Jira.ReplaceAllString(*result, "##### ")
	*result = header6Jira.ReplaceAllString(*result, "###### ")

}

// ConvertCodeBlocks function to convert code blocks
func ConvertCodeBlocks(result *string) {

	/*
		 From:

		 {code:json}{
			"attribute1": "value1"
		 }{code}

		 To:

		 ```json
		 {
			"attribute1": "value1"
		 }
		 ```
	*/
	codeBlockJira := regexp.MustCompile(`(?mU)\{code(:([a-z]+))?\}((?s).*)\{code\}`)

	*result = codeBlockJira.ReplaceAllString(*result, "```"+`$2`+"\n"+`$3`+"\n"+"```")

}

// ConvertTextEffects function to conver text effects https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=texteffects
func ConvertTextEffects(result *string) {

	// *strong in jira* -> `strong in jira`
	strongJira := regexp.MustCompile(`(?U)\*([^}]+)\*`)

	// _emphasis in jira_ -> *emphasis in jira*
	emphasisJira := regexp.MustCompile(`(?U)_([^}]+)_`)

	// ??citation in jira?? -> <cite>citation in jira</cite>
	citationJira := regexp.MustCompile(`(?U)\?\?([^}]+)\?\?`)

	// -deleted in jira- -> ~~deleted in jira~~
	deletedJira := regexp.MustCompile(`(?U)\-([^}]+)\-`)

	// +inserted in jira+ -> <ins>inserted in jira</ins>
	insertedJira := regexp.MustCompile(`(?U)\+([^}]+)\+`)

	// +superscript in jira+ -> <sup>superscript in jira</sup>
	superscriptJira := regexp.MustCompile(`(?U)\^([^}]+)\^`)

	// ~subscript in jira~ -> <sub>subscript in jira</sub>
	subscriptJira := regexp.MustCompile(`(?U)\~([^}]+)\~`)

	// {{monospaced in jira}} -> `codes in jira`
	monospacedJira := regexp.MustCompile(`(?U)\{\{([^}]+)\}\}`)

	*result = strongJira.ReplaceAllString(*result, "**"+`$1`+"**")
	*result = emphasisJira.ReplaceAllString(*result, "*"+`$1`+"*")
	*result = citationJira.ReplaceAllString(*result, "<cite>"+`$1`+"</cite>")
	*result = subscriptJira.ReplaceAllString(*result, "<sub>"+`$1`+"</sub>")
	*result = deletedJira.ReplaceAllString(*result, "~~"+`$1`+"~~")
	*result = insertedJira.ReplaceAllString(*result, "<ins>"+`$1`+"</ins>")
	*result = superscriptJira.ReplaceAllString(*result, "<sup>"+`$1`+"</sup>")
	*result = monospacedJira.ReplaceAllString(*result, "`"+`$1`+"`")

}
