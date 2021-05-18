package builder

import (
    "fmt"
    "os"
    "bytes"
    "text/template"
    "github.com/Enviy/pipelineBuilder/scaffold"
)

const outputDirPath string = "/Users/enviy/"

// Param is argument to populate query.yaml
type Param struct {
	RuleName string
	RuleDescription string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// BuildPipeline takes string input and writes to a python file
func BuildPipeline(name,age  string) {
	var doc bytes.Buffer

	f, err := os.Create(outputDirPath + name + ".py")
	check(err)
	defer f.Close()

	m := map[string]interface{}{"name": name, "age": age}
	t := template.Must(template.New("").Parse(scaffold.Pipeline))
	t.Execute(&doc, m)

	fmt.Println(doc.String())
	writtenValue, err := f.WriteString(doc.String())
	check(err)
	fmt.Printf("wrote %d bytes\n", writtenValue)
	f.Sync()
}

// BuildQuery populates a new query.yaml file
func BuildQuery(param Param) {
	var doc bytes.Buffer
	f, err := os.Create(outputDirPath + param.RuleName + ".yaml")
	check(err)
	defer f.Close()

	m := map[string]interface{}{"ruleName": param.RuleName, "ruleDescription": param.RuleDescription}
	t := template.Must(template.New("").Parse(scaffold.Query))
	t.Execute(&doc, m)
	_, err = f.WriteString(doc.String())
	check(err)
	f.Sync()
}

