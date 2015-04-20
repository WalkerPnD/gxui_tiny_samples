# 1. Fist Step

My Environment:  
OSX yosemite 10.10.3  
Golang 1.4.2

Add the text below to **"~/.bash_profile"**.  
I don't know why but **Adding to "~/.bashrc" didn't work for me.**
```
export GOPATH=$HOME/Documents/go
export PATH=$PATH:$GOPATH/bin
```

Restart or type the code below to reload the terminal  

```
> source ~/.bash_profile
```

# 2. Install gxui

I sudgest to install in this order because of the dependencies.
```terminal
brew install glew
go get http://code.google.com/p/freetype-go/freetype/raster
go get code.google.com/p/freetype-go/freetype/truetype
go get github.com/go-gl/gl/v3.2-core/gl
go get github.com/go-gl/glfw/v3.1/glfw
go get github.com/google/gxui
```


# 3. Download fonts
[SourceCodePro](https://www.google.com/fonts#UsePlace:use/Collection:Source+Code+Pro)

# 4. Official examples
```
go install github.com/google/gxui/samples/...
```

# # Sites
[gxui on GitHub](https://github.com/google/gxui)  
[gxui/sample  on GitHub](https://github.com/google/gxui/tree/master/samples)  
[go-gl/gl on GitHub](https://github.com/go-gl/gl)  
[go-gl/glfw on GitHub](https://github.com/go-gl/glfw)  
[go-gl/examples on GitHub](https://github.com/go-gl/examples)  
