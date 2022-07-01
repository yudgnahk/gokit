# gokit
gokit is a toolkit to init and generate some useful components for a Golang RESTful web service

it uses:
- Repository service design pattern
- [go-gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm) 
- mysql
- [swaggo](https://github.com/swaggo/swag) for docs
- [beego](https://github.com/beego/bee) for migrations

## Installation
```
go get -u github.com/yudgnahk/gokit
```

## Usage
1. Init project:
```
gokit init
```

2. Generate some stuffs:
```
gokit new [type] [name]
```
In which type:
- `controller` - it will generate a controller, a service and a repository
- `service`    - it will generate a service and a repository
- `repository` - it will generate only a repository

3. Version:
```
gokit version
```

## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
