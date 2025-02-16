# install curl, git, ...
apt-get update
apt-get install -y curl git jq

curl -OL https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz
rm go1.23.5.linux-amd64.tar.gz

$GOPATH="/usr/local/go/bin"
echo "PATH=\"$PATH:$GOPATH\"" >> ~/.profile
echo "PATH=\"$PATH:$GOPATH/bin\"" >> ~/.profile

