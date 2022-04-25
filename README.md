<h1 align="center">Welcome to hi 👋</h1>

> 一个以太坊类型公链的命令行工具，可以进行快速address创建，原生币的一对一、一对多、多对一、多对多转账。

# Roadmap
## account
- [x] 批量创建

## transfer
### native
- [x] 一对一
- [x] 一对多
- [x] 多对多
- [x] 多对一

### token
- [ ] erc20 
- [ ] erc721
- [ ] erc1155

**备注：标准的ERC**
## contract
- [ ] compile - 编译成API 或者 go 源代码文件，自动下载 solc & abigen
- [ ] depoly - 部署, 如果未编译则自动编译并部署到指定节点
- [ ] call - 指定ABI文件调用指定函数


# Prerequisites

## Install
```shell
go mod download
```

## Build

```shell
make build
```

## Author

👤 **lambda-yu**


## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
