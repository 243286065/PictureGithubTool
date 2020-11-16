# PictureGithubTool
使用github作为图床的工具，快速上传图片到github，并获取链接地址。

# 运行方法
Linux下可以直接执行：
```
./start-all.sh
```

Windows下还没写启动后脚本，可以直接使用go运行：
```
go run service/apigw/main.go
```

然后就可以通过浏览器访问`http://ip:8081/`, ip为启动服务的主机ip，端口号可以通过修改`config/web_server.go`文件中的`WebServerHost`字段来修改，默认是8081。

# 使用的工具及插件
## bootstrap
有名的web ui前端开发框架。 可访问https://www.bootcss.com/ 。

## bootstrap-fileinput
基于bootstrap实现的文件上传插件，示例：https://plugins.krajee.com/file-basic-usage-demo 。

## bootstrap-select
基于bootstrap实现的下拉搜索插件，官网：https://developer.snapappointments.com/bootstrap-select/ 。


# 运行效果
![image](https://raw.githubusercontent.com/243286065/pictures_markdown/master/test/cbdb0a20d3b1ecda8469068d31cd3f79.png)

![image](https://raw.githubusercontent.com/243286065/pictures_markdown/master/test/66041c900017ef153dc7a2c0ca9c42fb.png)

# 实现逻辑
* 主要是使用[github v3 api](https://developer.github.com/v3/)。这里使用golang开发主要是为了锻炼，完全可以只用js实现。

* 本身工具是完全的客户端程序，但是因为对bootstrap-fileinput插件，如何直接上传到github没弄清楚，所以暂时还是有服务端，客户端先通过bootstrap-fileinput上传图片到服务端，服务端再上传到github.


# 注意事项
* 敬告：github已禁用用户名和密码的认证方式，改用access token的方式，详见[官方文档](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)。


* 如果上传失败可能的原因：
    * github暂时无法连接，你可以重新登录试试；
    * 没有选择仓库；
    * 路径填写不正确，可以填写例如"a/b"，这样会将文件存放在'a/b'目录。

* 请注意仓库不要使用私有仓库，这会导致未认证时无法访问；也请注意不要轻易删除存放图片的仓库里的文件。

* 实现比较仓促，可能会有很多bug或者错误逻辑。

# 版本说明
* v1.0: 支持基础的图片上传功能;
* v1.1: 支持粘贴板粘贴文件.(每次仍只允许最多上传1个文件,因此粘贴板粘贴的文件可能第一次上传失败,可以刷新重新上传).
* v1.2:改用access token认证。