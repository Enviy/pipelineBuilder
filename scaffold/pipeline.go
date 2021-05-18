package scaffold


// Pipeline is the pipeline.py scaffold
var Pipeline = `
name = "{{.name}}"
age = "{{.age}}"
print("Hello {0}, your age is {1}.".format(name, age))
`
