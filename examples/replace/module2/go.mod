module github.com/example/module2

go 1.21

replace github.com/example/module1 => ../module1

require github.com/example/module1 v0.0.0-00010101000000-000000000000
