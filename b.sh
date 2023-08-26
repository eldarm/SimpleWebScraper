echo "Compiling scrape GOARCH=${GOARCH} GOOS=${GOOS}..."
go build scrape/*
echo "Compiling main..."
go build main/scraper.go
