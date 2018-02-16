# Simple github api front end
    The project is implemented in Go programming language, we use the opensource
    "github.com/spf13/cobra" as the cli/console application framework 

# How to build
    install Go, https://golang.org/doc/install
    make sure $GOPATH has been configured

```
    cd $GOPATH
    go get github.com/wy3148/github_cli
    cd $GOPATH/src/github.com/wy3148/github_cli
    go build
```

# How to use
    the default github repo is "https://api.github.com/repos/gorilla/mux", if you want to
    change a different repo, set the system environment variable 'github_repo'.
    Regarding the configuration of system environment in windows, refer the 
    "http://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-windows-command-line-and-registry/"

    in mac os or linux, simply 'export' this variable before run the application

    ./github_cli help
    ./github_cli pulls --state open
    ./github_cli pulls --state closed