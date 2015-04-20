# 1. Primeiro Passo

Adicione no **"~/.bash_profile"** abaixo.
```
export GOPATH=$HOME/Documents/go
export PATH=$PATH:$GOPATH/bin
```

Reinice ou digite o código para reiniciar
```
> . ~/.bash_profile
```

# 2. Intalar o gxui

```terminal
brew install glew // Dependência
go get http://code.google.com/p/freetype-go/freetype/raster // Opcional Font
go get code.google.com/p/freetype-go/freetype/truetype
go get github.com/go-gl/gl/v3.2-core/gl
go get github.com/go-gl/glfw/v3.1/glfw
go get github.com/google/gxui
```


# 3. Baixar as fonters
[SourceCodePro](https://www.google.com/fonts#UsePlace:use/Collection:Source+Code+Pro)

# 4. Os exemplos
```
go install github.com/google/gxui/samples/...
```

# Sites
[go-gl/gxui](https://github.com/google/gxui)
[go-gl/gl](https://github.com/go-gl/gl)
[go-gl/glfw](https://github.com/go-gl/glfw)
[go-gl/examples](https://github.com/go-gl/examples)
[https://github.com/google/gxui/tree/master/samples](https://github.com/google/gxui/tree/master/samples)
