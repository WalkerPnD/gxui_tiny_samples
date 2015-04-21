# 1. はじめの一歩

環境:  
OSX yosemite 10.10.3  
Golang 1.4.2

**"~/.bash_profile"**に下の２行を追加.  
理由はわかりませんが**".bashrc"に追加しても動きませんでした**
```
export GOPATH=$HOME/Documents/go
export PATH=$PATH:$GOPATH/bin
```

ターミナルを再起動するか、下のコードを打つこ度でもリロードできます。
```
> source ~/.bash_profile
```

# 2. gxuiのインストール

この順番でインストールすると良いようです。	
```terminal
brew install glew
go get http://code.google.com/p/freetype-go/freetype/raster
go get code.google.com/p/freetype-go/freetype/truetype
go get github.com/go-gl/gl/v3.2-core/gl
go get github.com/go-gl/glfw/v3.1/glfw
go get github.com/google/gxui
```


# 3. フォントのダウンロード
[SourceCodePro](https://www.google.com/fonts#UsePlace:use/Collection:Source+Code+Pro)

# 4. 付属サンプルのダウンロード
```
go install github.com/google/gxui/samples/...
```

# # 公式サイトなど
[gxui on GitHub](https://github.com/google/gxui)  
[gxui/sample  on GitHub](https://github.com/google/gxui/tree/master/samples)  
[go-gl/gl on GitHub](https://github.com/go-gl/gl)  
[go-gl/glfw on GitHub](https://github.com/go-gl/glfw)  
[go-gl/examples on GitHub](https://github.com/go-gl/examples)  
