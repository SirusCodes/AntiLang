# install curl, git, ...
apt-get update
apt-get install -y curl git jq

VERSION="1.24.0"

curl -OL https://go.dev/dl/go$VERSION.linux-amd64.tar.gz
tar -C /usr/local -xzf go$VERSION.linux-amd64.tar.gz
rm go$VERSION.linux-amd64.tar.gz

GOPATH="/usr/local/go"
echo "PATH=\"$PATH:$GOPATH:$GOPATH/bin\"" >> ~/.profile

source ~/.profile

go get -u -v golang.org/x/tools/cmd/gopls