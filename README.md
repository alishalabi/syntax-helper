[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity) [![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://GitHub.com/Naereen/ama) [![Documentation Status](https://readthedocs.org/projects/ansicolortags/badge/?version=latest)](http://ansicolortags.readthedocs.io/?badge=latest) [![GitHub issues](https://img.shields.io/github/issues/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/issues/) [![Awesome Badges](https://img.shields.io/badge/badges-awesome-green.svg)](https://github.com/Naereen/badges) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)



# Syntax Helper CLI

Syntax Helper is a CLI that helps developers quickly get access to built-in code syntax. Syntax Helper supports a variety of code packages for both Golang and Python. Data is scrapped directly from https://golang.org/pkg/ & https://docs.python.org/3/library/.

## Installation

Install using `go get` command:
```bash
go get github.com/alishalabi/syntax-helper
```

Install locally using:
```bash
git clone git@github.com:alishalabi/syntax-helper.git
```

## Usage
- Note: All functionality requires broadband connection

Get all Golang functions for given package name:
```bash
./syntax-cli golang <pkg-name>
```
Example:
```
➜  syntax-cli git:(master) ./syntax-cli golang unsafe                
func Alignof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Sizeof(x ArbitraryType) uintptr
type ArbitraryType
type Pointer
```


Get syntax for specific public function from package (Note: function names all begin with a capital letter)
```bash
./syntax-cli golang <pkg-name> <function-name>
```
Example:
```
➜  syntax-cli git:(master) ./syntax-cli golang unsafe Offsetof
Function Syntax:
func Offsetof(x ArbitraryType) uintptr
For Synatx example:
$ ./syntax-cli golang unsafe Offsetof example
```


Get example of specific public function from package (Note: function names all begin with a capital letter)
```bash
./syntax-cli golang <pkg-name> <function-name> example
```
Example:
```
➜  syntax-cli git:(master) ./syntax-cli golang fmt Fprintf example    
Function Example:
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Milestones
- [x] Generate skeleton structure
- [x] Implement `Golang` command
- [x] Implement `Golang package` parameter
- [x] Implement `Golang function` parameter
- [x] Implement `Golang function example` parameter
- [x] Implement Go-Colly to populate messages
- [x] Implement `Python` command
- [x] Implement `Python package` parameter
- [] Implement `Python function` parameter
- [] Implement `Python function example` parameter
- [x] Finalize ReadMe.md
- [x] Implement badges
- [] Deploy as Homebrew package


## License
Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements. See the NOTICE file distributed with this work for additional information regarding copyright ownership. The ASF licenses this file to you under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[![ForTheBadge built-with-love](http://ForTheBadge.com/images/badges/built-with-love.svg)](https://GitHub.com/Naereen/)
