GOARCH="amd64"
GOOS="windows"
echo "Compiling scrape GOARCH=${GOARCH} GOOS=${GOOS}..."
go build scrape/*
echo "Compiling main..."
go build main/scraper.go
# go -o bin/scraper.exe build main/scraper.go
