## My Go Goes to 11
ubuntu@ip-172-31-40-93:~/goWorkSpace$ go install $GOFISH

ubuntu@ip-172-31-40-93:~/goWorkSpace$ $RUNFISH

export GOPATH=$HOME/ubuntu/goWorkspace/src/github.com/EricRiney/GoFish

export GOBIN=$GOPATH/bin

export PATH=$PATH:$GOBIN

go install server.go

go run server.go