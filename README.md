# Syntax Helper CLI

Syntax Helper is a CLI that helps developers quickly get access to built-in code syntax. Syntax Helper supports a variety of code packages for both Golang and Python. Data is scrapped directly from https://golang.org/pkg/ & https://docs.python.org/3/library/.

## Installation

Install locally using:
```bash
git clone git@github.com:alishalabi/syntax-helper.git
```

## Usage
- Note: All functionality requires broadband connection

- Get all Golang functions for given package name:
```bash
./syntax-cli golang <pkg-name>
```

- Get syntax for specific public function from package (Note: function names all begin with a capital letter)
```bash
./syntax-cli golang <pkg-name> <function-name>
```

- Get example of specific public function from package (Note: function names all begin with a capital letter)
```bash
./syntax-cli golang <pkg-name> <function-name> example
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
- [] Finalize ReadMe.md
- [] Implement badges
- [] Deploy as Homebrew package


## License
Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements. See the NOTICE file distributed with this work for additional information regarding copyright ownership. The ASF licenses this file to you under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
