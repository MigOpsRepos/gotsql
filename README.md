# GoTSQL - A Better Way to Organize SQL codebase using Templates in Golang
​


1. Installation through Go Get command
    
    ```bash
    $ go get github.com/migopsrepos/gotsql
    ```
    
2. Initialize the GoTSQL object
    
    ```go
    import "github.com/migopsrepos/gotsql"
    ...
    g := gotsql.GoTSQL{}
    ```
    
3. Load the template file or directory with template files. (Templates are discussed in detail in next section)
    
    ```go
    g.Load("library/")
    ...
    g.Load("library_dev/")
    ```
    
4. Get the query using namespace and use it in Golang
    
    ```go
    query, err := g.Get("library/books/crud/getAllBooks", nil)
    ...
    rows, err := db.Query(query, ...)


> You can find the usage and more examples in this [article](https://www.migops.com/blog/2021/10/22/organizing-postgresql-codebase-using-templates-in-golang/) 
