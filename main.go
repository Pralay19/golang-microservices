package main

import "coursemicro/app"

// now here whenever i import a package I'll have to use "coursemicro/..." that is because
// I have set the module name to coursemicro in go.mod file. And there will be only one module
// in a project, and its situated at the root of the project.

func main() {
	app.StartUp()
}
