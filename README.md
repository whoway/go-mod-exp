# go-mod-exp
Go mod的导入实验

背景介绍

1、如何使用go module导入gitlab/github之类的公有/私有仓库非常简单  
2、但是初学者在写自己的hello world的时候、或者项目开发的时候  
想要使用go的import导本地文件的时候，会遇到一些问题，很多人就容易卡在这里  

目标：
1、通过实验，梳理明白go module的import本地文件的方式  
2、注意GO111MODULE=on的设置  
3、仔细观察，其实你go mod init的包名，你也可以不那么规范  
- 规范的：go mod init github.com/whoway/go-mod-exp
- 不规范的：go mod init exp1


结论：

1、除去package main这个包，Go语言规定Go语言的一个目录下！无论你有多少个文件，都必须是同一个package【包】【见实验2】
- 特殊情况：package main这个包，在同一个目录下只能有1个文件里面写package main这个包，否则会报错！！【实验3】  
2、Go语言的目录下，你写的package【包】的名字，虽然不是必须和目录名一致！但是比如在百度、腾讯之类公司的编码规范中  
是非常建议你写成和目录名一致的【不然就会出现实验1那样的，不规范的事情，你在import的目是iamsoryy这样的目录，  
但是却用的他里面的package为two的包，就很令人费解】  
3、我总结的import规则是: go.mod中对应的module名字/相对那个go.mod的一系列路径【而不是到你的package的名字！】  


## 导入本地项目的包

### exp1

- 实验1是一个集各种不规范的一个合集【但是就是能跑】
 - 1、package main这个包，故意写在一个不叫main.go的文件中
 - 2、iamsorry目录下的go故意不写成iamsorry.go，并且这个go代码中他的package名字故意写为two而不是iamsorry

```bash
➜  exp1 git:(main) ✗ go run demo.go 
main function
我是One目录下的package one的Demo函数...
我是iamsorry目录下的package two的Demo函数...
```

### exp2

```bash
➜  exp2 git:(main) ✗ go run main.go 
Hello, 世界
我是在one.go中用packeage one的Demo函数
我是在oneplus.go中用packeage one的DemoPlus函数
```


### exp3

```bash
➜  exp3 git:(main) ✗ go run main.go 
# command-line-arguments
./main.go:9:2: undefined: Temp
```


## 跨项目的导入+开发时候临时使用require和replace

### exp4

使用场景：要导入的package不在同一个项目下+【关于require和replace】的使用
有一种开发场景：在本地开发的时候，你还没把自己的其他包上传到gitlab/github上，但是你写了一个demo
那你就要暂时使用require来引入你本地的包
- 见参考资料

```bash
module github.com/q1mi/p1

go 1.14


require "liwenzhou.com/q1mi/p2" v0.0.0
replace "liwenzhou.com/q1mi/p2" => "../p2"

```


## 参考资料

- https://www.liwenzhou.com/posts/Go/import-local-package/#c-0-2-1


## 附录

```bash
go env -w GO111MODULE=on
go mod init exp1
go mod init exp2
go mod init exp3
```