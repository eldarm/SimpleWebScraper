echo "Compiling main for Mac..."
GOARCH="arm64" GOOS="darwin" go build main/scraper.go
echo "Compiling main for linux..."
GOARCH="amd64" GOOS="linux" go build -o scraperlx main/scraper.go
echo "Compiling main for windows..."
GOARCH="amd64" GOOS="windows" go build main/scraper.go
echo "Publishing..."
cp scraper bin/scraper.app
cp scraperlx bin/
cp scraper.exe bin/
