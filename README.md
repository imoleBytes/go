# Learnig Go

## Install Go
1.   download the version you want from go.dev/learn
2.   make sure you remove the one you have before by deleting the go directory in your /usr/local
`sudo rm -rf /usr/local/go`
3.   extract the archive you just downloaded into /usr/local, this will create a fresh go directory
`sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz`
4.   Add the /bin path to the the PATH variablesoyou can access go command from everywhere
`export PATH=$PATH:/usr/local/go/bin`
5. finally verify that you have istalled go by runnig `go version`. You should get the version installed.


