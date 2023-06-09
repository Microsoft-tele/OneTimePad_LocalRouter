# 项目名称

OneTimePadLocalRouter
# 项目描述

该项目是一个密码管理工具，旨在为用户提供安全、高效、易用的密码管理服务。用户可以使用该工具生成、存储和管理他们的密码，从而避免因使用相同密码或简单密码而导致的安全问题。

该项目使用了OTP（One-Time Pad）和RSA加密算法，以及Go语言编写的服务器程序，为用户提供强大的安全保障。用户可以通过一个网页界面方便地管理他们的密码，无需下载任何客户端软件。
# 安装指南

该项目的服务器程序运行在Go语言环境中，您需要在安装之前安装Go语言。
首先，您需要从GitHub上获取源代码。您可以使用Git将源代码克隆到本地：

```
git clone https://github.com/Microsoft-tele/OneTimePad_LocalRouter
```

进入项目目录，使用以下命令编译并运行服务器程序：
```
go build
./OneTimePadLocalRouter
```

该程序会监听本地的8080端口，并开始运行。现在，您可以在浏览器中访问http://localhost:8080来使用该程序。

# 使用指南

    注册账户：在使用该程序之前，您需要注册一个账户。在登录界面中，点击“注册”按钮，填写所需信息即可完成注册。

    登录：在注册完成后，您可以使用您的邮箱和密码登录该程序。

    生成OTP密码：登录成功后，您可以点击“生成OTP”按钮，填写相关信息，生成一个新的OTP密码。

    查看OTP密码：在OTP列表中，您可以查看您生成的所有OTP密码。

    加密密码：在“加密”页面中，您可以使用OTP算法加密您的密码。

    查看个人信息：在登录成功后，您可以在主页中查看您的个人信息，包括已存储的OTP密码。

# 联系方式

如果您有任何问题或建议，请随时联系我们。您可以通过发送电子邮件至liweijun0302@gmail.com来联系我们。
