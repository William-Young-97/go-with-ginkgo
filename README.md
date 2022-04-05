## Personal notes

### Ginkgo set up

- First create a go.mod file and specify:
    module example.com(If version controlled link to repo)/the name of the package (directory) that I am working in/testing

- Create your go file in that directory and ensure that you state the *import name of your package/directory*

- Next run:

`` go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/... ``

- Then cd into that directory and run:

`` ginkgo bootstrap ``

- Observe the file generates and ensure that you click through your go mod file so that you have your go.sum checklist.


