# Go with Ginkgo
### Aims
- To use BDD with Go and Ginkgo to test drive basic applications.
- To familiarise myself with Go syntax
- To understand common practices and quirks of Go

## Personal notes

### Ginkgo set up

- First create a go.mod file using 

    `` go mod init example.com/greetings 
    
    <br>

    (Use my github)/the name of the module (root directory) that I am working in. ``

- Create your go file in that directory and ensure that you state the *import name of your package/sub-directory*

- Next run:

`` go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo

    <br>

go get github.com/onsi/gomega/... ``

- This will install ginkgo and gomega into your go.mod file and generate a sum file (Checklist).

- cd into your sub-directory and run:

`` ginkgo bootstrap ``

- This will generate a suite test which links ginkgo and gomega to go test.

- Use ginkgo generate to create a seperate test file where you will write your tests for the package.

### Sytnax
#### Public vs private
- Setting the first letter of a variable to uppercase makes it public, meaning it can be exported.
- var Age == public
- var age == private

### Commands

`` ginkgo -r `` 
Will run tests recursively across all test suites


