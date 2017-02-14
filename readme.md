# Goweb - live chat demo with Gin and Melody

### Installation
Goweb depends on [Gin](https://github.com/gin-gonic/gin) and [Melody](https://github.com/olahol/melody), so those need to be installed before Goweb is built. Otherwise, it's the usual clone and build pipeline.

```
go get github.com/gin-gonic/gin github.com/olahol/melody
git clone https://github.com/inoriy/goweb
cd goweb
go build main.go -o goweb
```

### Usage
Run the executable and point a browser to `localhost:3000`

```
./goweb
```