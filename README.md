# Go with Ginkgo
### Aims
- To use BDD with Go and Ginkgo to test drive basic applications.
- To familiarise myself with Go syntax
- To understand common practices and quirks of Go

## Personal notes

### Ginkgo set up

- First create a go.mod file and specify:
    module example.com(If version controlled link to repo)/the name of the package (directory) that I am working in/testing

- Create your go file in that directory and ensure that you state the *import name of your package/directory*

- Next run:

`` go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo;
go get github.com/onsi/gomega/... ``

- Then cd into that directory and run:

`` ginkgo bootstrap ``

- Observe the file generates and ensure that you click through your go mod file so that you have your go.sum checklist.

### Sytnax
#### Public vs private
- Setting the first letter of a variable to uppercase makes it public, meaning it can be exported.
- var Age == public
- var age == private

### Commands

`` ginkgo -r `` 
Will run tests recursively across all test suites


