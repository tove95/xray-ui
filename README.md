xray-ui 管理脚本使用方法: 
----------------------------------------------
xray-ui              - 显示管理菜单
xray-ui start        - 启动 xray-ui 面板
xray-ui stop         - 停止 xray-ui 面板
xray-ui restart      - 重启 xray-ui 面板
xray-ui status       - 查看 xray-ui 状态
xray-ui enable       - 设置 xray-ui 开机自启
xray-ui disable      - 取消 xray-ui 开机自启
xray-ui log          - 查看 xray-ui 日志
xray-ui v2-ui        - 迁移本机器的 v2-ui 账号数据至 xray-ui
xray-ui update       - 更新 xray-ui 面板
xray-ui geoip        - 更新 geoip ip库
xray-ui update_shell - 更新 xray-ui 脚本
xray-ui install      - 安装 xray-ui 面板
xray-ui x25519       - REALITY  key 生成
xray-ui ssl_main     - SSL 证书管理
xray-ui ssl_CF       - Cloudflare SSL 证书
xray-ui crontab      - 添加geoip到任务计划每天凌晨1.30执行
xray-ui uninstall    - 卸载 xray-ui 面板

本项目基于上游X-UI项目进行略微的功能改动！后续将紧跟上游X-UI版本更新！在此感谢[vaxilu](https://github.com/vaxilu/x-ui)及各位为此项目做出贡献

----------------------------------------------------------------------------------------------------------------------------------------------
### 功能介绍

系统状态监控
支持多协议，网页可视化操作
支持的协议：vmess、vless、trojan、shadowsocks、dokodemo-door、socks、http
支持配置更多传输配置
流量统计，限制流量，限制到期时间
可自定义 xray 配置模板
支持 https 访问面板（自备域名 + ssl 证书）
更多高级配置项，详见面板

----------------------------------------------------------------------------------------------------------------------------------------------

本脚本显示功能更加人性化！已解决各种新老系统安装失败问题，并会长期更新，欢迎大家提建议！！
    
### VPS 安全加固 建议

<details>
  <summary>点击查看 VPS 安全加固 建议</summary>

1、保持内核版本更新修复内核级别漏洞 

2、开启防火墙千万别裸奔

3、安装 Fail2ban配置 ssh nginx 自动封禁可疑 IP 地址

4、不要使用简单密码，或者禁止使用密码登录，使用 RSA 私钥登录

5、配置ssh 密码重试次数

6、添加指定IP 登陆ssh 删除防火墙里面放行所以ssh 的服务 如果你自己没固定IP 使用clash 或者其它可以代理的ssh 让你vps ip 连接自己

clash 规则   `- DST-PORT,22,ACCESS-DENIED`

firewall 设置

Ubuntu/debain 安装firewall

```bash
# 1、关闭默认 ufw
# 停止ufw服务
sudo systemctl stop ufw
# 关闭开机启动
sudo systemctl disable ufw
# 删除ufw规则
sudo ufw --force reset
# 2 安装firewall
apt update 
apt install -y firewalld
#重载 （在增减规则后，需要重载）
firewall-cmd --reload
#启动
systemctl start firewalld
#重启
systemctl restart firewalld
#设置开机启动
systemctl enable firewalld
#关闭开机启动
systemctl disable firewalld
```

centos/Rocky/Redhat 安装firewall

```bash
yum install -y firewalld
#重载 （在增减规则后，需要重载）
firewall-cmd --reload
#启动
systemctl start firewalld
#重启
systemctl restart firewalld
#设置开机启动
systemctl enable firewalld
#关闭开机启动
systemctl disable firewalld
```

```bash
# 添加 指定ip 访问ssh 服务
firewall-cmd --permanent --add-rich-rule='rule family=ipv4 source address=10.0.0.1/32 service name=ssh accept'
firewall-cmd --permanent --add-rich-rule='rule family="ipv4" source address="103.119.132.41/32" service name="ssh" accept'
# 关闭 ssh 所有ip 访问 服务
firewall-cmd --remove-service=ssh --permanent
# 添加http https 服务
firewall-cmd --add-service=http --permanent
firewall-cmd --add-service=https --permanent
# 生效配置
firewall-cmd --reload
# 查看规则
firewall-cmd --list-all
```
</details>

----------------------------------------------------------------------------------------------------------------------------------------------
更新日志：

2023.8.8 更新到最新的依赖，添加fragment 用于控制发出的 TCP 分片，在某些情况下可以欺骗审查系统，比如绕过 SNI 黑名单。[客户端配置模块参考](./media/xray.json) [官方配置文档](https://xtls.github.io/config/outbounds/freedom.html#outboundconfigurationobject)
  
2023.5.29 添加xray-ui crontab 命令 添加geoip更新到计划任务默认凌晨1.30执行 你可以修改 /etc/crontab文件

2023.5.15 添加xray-ui geoip 更新IP库 添加数据库导入导出

2023.5.6 修复cipherSuites 配置多选 把分割符号从,改成: 由于没看文档就按照常规添加的。很抱歉啊，如果一开始配置了cipherSuites 请改成auto 然后在升级并添加网卡接口可见。

2023.5.4 sniffing 多选，tls cipherSuites 配置多选！

2023.4.28 添加REALITY分享随机选择sni方便使用any_SNI_No_SNI配置方案

2023.4.26 [添加 Nginx前置SNI分流](./Nginx前置SNI分流.md)

2023.4.24 添加一键更新geoip,geosite 添加geoip,geosite 更新版本号

2023.4.23 添加docker镜像

### SSL证书

<details>
  <summary>点击查看SSL证书详情</summary>

### ACME
使用ACME管理SSL证书：

1. 确保您的域名正确解析到服务器。
2. 在终端中运行 `xray-ui` 命令，然后选择 `SSL证书管理`。
3. 您将看到以下选项：

   - **获取 SSL:** 获取SSL证书。
   - **撤销证书:** 吊销现有的SSL证书。
   - **强制续期:** 强制更新SSL证书。

### Certbot

安装并使用Certbot：

```bash
apt-get install certbot -y
certbot certonly --standalone --agree-tos --register-unsafely-without-email -d yourdomain.com
certbot renew --dry-run
```

### Cloudflare

管理脚本内置了Cloudflare的SSL证书申请。要使用此脚本申请证书，您需要以下信息：

- Cloudflare注册的电子邮件
- Cloudflare全局API密钥
- 域名必须通过Cloudflare解析到当前服务器

**如何获取Cloudflare全局API密钥：**

1. 在终端中运行 `xray-ui` 命令，然后选择 `Cloudflare SSL证书`。
2. 访问链接：[Cloudflare API Tokens](https://dash.cloudflare.com/profile/api-tokens)。
3. 点击“查看全局API密钥”（参见下图）：
   ![](media/APIKey1.PNG)
4. 您可能需要重新验证您的账户。之后将显示API密钥（参见下图）：
   ![](media/APIKey2.png)

使用时，只需输入您的 `域名`、`电子邮件` 和 `API密钥`。如下图所示：
   ![](media/DetailEnter.png)

### xray-ui 配置ssl证书
```bash
xray-ui  选择22
输入证书路径跟密钥路径
手动配置证书
/usr/local/xray-ui/xray-ui  cert -webCert /root/cert/你的域名/fullchain.pem -webCertKey /root/cert/你的域名/privkey.pem
清理证书不输入任何东西直接回车
重启xray-ui
xray-ui  选择10

```
### xray-ui 配置mTLS
```bash
xray-ui  选择23
输入证书路径跟密钥路径CA路径
手动配置证书
/usr/local/xray-ui/xray-ui cert -webCert /root/cert/你的域名/fullchain.pem -webCertKey /root/cert/你的域名/privkey.pem -webCa /root/cert/ca.cer
清理证书不输入任何东西直接回车
重启xray-ui
xray-ui  选择10
```
</details>

### docker运行
<details>
  <summary>点击查看 docker运行</summary>

```bash
# juestnow/xray-ui:latest 最新版本 指定版本号docker pull juestnow/xray-ui:1.8.6
 docker run -d --net=host -v/etc/xray-ui:/etc/xray-ui  -v/root/cert:/root/cert --restart=unless-stopped --name xray-ui juestnow/xray-ui:latest
# 查看默认账号密码
docker exec -ti  启动的容器名 /app/xray-ui setting -show
docker exec -ti xray-ui  /app/xray-ui setting -show
# 设置账号密码
docker exec -ti  启动的容器名 /app/xray-ui setting -password abcd -username abacd 
docker exec -ti xray-ui  /app/xray-ui setting -password abcd -username abacd
# 设置path 
docker exec -ti  启动的容器名 /app/xray-ui setting --webBasePath aaaaddffdf
docker exec -ti xray-ui  /app/xray-ui setting --webBasePath aaaaddffdf
# 证书配置 
## TLS 配置
docker exec -ti xray-ui  /app/xray-ui  cert -webCert /root/cert/你的域名/fullchain.pem -webCertKey /root/cert/你的域名/privkey.pem
## mTLS 配置
docker exec -ti xray-ui  /app/xray-ui cert -webCert /root/cert/你的域名/fullchain.pem -webCertKey /root/cert/你的域名/privkey.pem -webCa /root/cert/ca.cer
# 第一次访问
当前面板http只支持127.0.0.1访问如果外面访问请用ssh转发或者nginx代理或者xray-ui 配置证书 选择22配置证书
ssh 转发 客户机操作 ssh  -f -N -L 127.0.0.1:22222(ssh代理端口未使用端口):127.0.0.1:54321(xray-ui 端口) root@8.8.8.8(xray-ui 服务器ip)
浏览器访问 http://127.0.0.1:22222(ssh代理端口未使用端口)/path(web访问路径)
```
</details>

### 第一次访问

<details>
  <summary>点击查看 手动安装</summary>

```bash
当前面板http只支持127.0.0.1访问如果外面访问请用ssh转发或者nginx代理或者xray-ui 配置证书 选择22配置证书
ssh 转发 客户机操作 ssh  -f -N -L 127.0.0.1:22222(ssh代理端口未使用端口):127.0.0.1:54321(xray-ui 端口) root@8.8.8.8(xray-ui 服务器ip)
例子：ssh  -f -N -L 127.0.0.1:22222:127.0.0.1:54321 root@8.8.8.8
浏览器访问 http://127.0.0.1:22222(ssh代理端口未使用端口)/path(web访问路径)
或者服务器执行 ssh -f -N -L 0.0.0.0:22222(ssh代理端口未使用端口):127.0.0.1:54321(xray-ui 端口) root@127.0.0.1 
例子：ssh -f -N -L 0.0.0.0:22222:127.0.0.1:54321 root@127.0.0.1
然后用你服务器地址+ssh转发端口访问

xshell 配置：https://netsarang.atlassian.net/wiki/spaces/ENSUP/pages/27295927/XDMCP+connection+through+SSH+tunneling
putty 配置：https://knowledge.exlibrisgroup.com/Voyager/Knowledge_Articles/Set_Up_SSH_Port_Forwarding_in_Putty
SecureCRT  配置：https://www.vandyke.com/support/tips/socksproxy.html
windows openssh 配置： https://www.cnblogs.com/managechina/p/18189889
```
</details>

2023.4.20 添加 配置文件下载本地，DB文件下载到本地，更新依赖到最新！

2023.4.17 添加uTLS REALITY x25519 使用go原生生成公钥私钥

2023.4.12 升级依赖模块 sockopt 可以在 REALITY TLS NONE 可用！增加REALITY分享连接shortId随机选择

2023.4.11 REALITY 配置 生成 x25519 shortIds等 ！

2023.4.7 添加 xray-ui x25519 生成REALITY公私钥 ！

[xray-ui 面板配置 reality](./reality.md)

2023.3.13 添加reality 支持 !

* [reality 配置参考](./media/reality.png)

2023.3.10 删除旧版XTLS配置以便支持xray1.8.0版本 旧trojan配置请关闭然后打开编辑从新保存即可正常，旧VLESS配置可能需要删除重新创建xray才能启动成功

2023.1.7 添加VLESS-TCP-XTLS-Vision 支持

2022.10.19 更新xray时不更新geoip.dat geosite.dat . geoip.dat geosite.dat 使用[Loyalsoldier](https://github.com/Loyalsoldier/geoip)提供版本单独更新

2022.10.17 更改trojan 可以关闭tls配置可以使用nginx 对外代理

-------------------------------------------------------------------------------------------------------------------------------------------------
### 手动安装

<details>
  <summary>点击查看 手动安装</summary>

```bash
# 下载 
wget  --no-check-certificate -O /usr/local/xray-ui-linux-amd64.tar.gz https://github.com/qist/xray-ui/releases/latest/download/xray-ui-linux-amd64.tar.gz

# 解压
    cd /usr/local/
    tar -xvf xray-ui-linux-amd64.tar.gz
    rm xray-ui-linux-amd64.tar.gz -f
    cd xray-ui
    chmod +x xray-ui bin/xray-linux-amd64
    cp -f xray-ui.service /etc/systemd/system/
    wget --no-check-certificate -O /usr/bin/xray-ui https://raw.githubusercontent.com/qist/xray-ui/main/xray-ui.sh
    chmod +x /usr/bin/xray-ui
    systemctl daemon-reload
    systemctl enable xray-ui
    systemctl start xray-ui
    # 设置账号密码：
    /usr/local/xray-ui/xray-ui setting -username admin -password admin123
    # 设置端口
   /usr/local/xray-ui/xray-ui setting -port  5432
```
</details>

### VPS直接运行一键脚本

```bash
bash <(curl -Ls  https://raw.githubusercontent.com/qist/xray-ui/main/install.sh)
```
#### 编译

<details>
  <summary>点击查看 编译</summary>

```bash
git clone https://github.com/qist/xray-ui.git

cd xray-ui
debian/ubuntu解决方案：sudo apt-get install libc6-dev
redhat/centos解决方案：yum install glibc-static.x86_64 -y 或者 sudo yum install glibc-static
CGO_ENABLED=1 go build -o xray-ui/xray-ui  -ldflags '-linkmode "external" -extldflags "-static"' main.go
# 交叉编译
在centos7中安装，yum install gcc-aarch64-linux-gnu
去https://releases.linaro.org/components/toolchain/binaries/ 找 latest-7
下载 aarch64-linux-gnu/sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu.tar.xz
自己找个目录, 解压 tar Jxvf sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu.tar.xz
build时，指定 sysroot 的位置。

用 CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC="aarch64-linux-gnu-gcc" CGO_CFLAGS="-g -O2 --sysroot=/..../sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu/" CGO_LDFLAGS="-g -O2 --sysroot=/..../sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu/" go build -v -ldflags "-w -s" -o xray-ui/xray-ui main.go 编译成功。
debian/ubuntu解决方案
apt install gcc-aarch64-linux-gnu
CGO_ENABLED=1 GOARCH=arm64 CC="aarch64-linux-gnu-gcc" go build -o xray-ui/xray-ui  -ldflags '-linkmode "external" -extldflags "-static"' main.go 
```
</details>

--------------------------------------------------------------------------------------------------------------------------------------------------
### nginx 代理设置

<details>
  <summary>点击查看 反向代理配置</summary>

```nginx
upstream xray-ui {
        least_conn;
        server 127.0.0.1:54321 max_fails=3 fail_timeout=30s;
        keepalive 1000;
}
server {
    listen 443;
    server_name xray.test.com;
    client_max_body_size 0;
    chunked_transfer_encoding on;
    client_body_buffer_size 202400k;
    client_body_in_single_buffer on;
    add_header Strict-Transport-Security "max-age=63072000; includeSubdomains; preload" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Frame-Options SAMEORIGIN always;
    add_header X-Content-Type-Options nosniff;
    add_header X-Frame-Options "DENY";
    add_header Alt-Svc 'h3=":443"; ma=86400, h3-29=":443"; ma=86400';
    ssl_certificate /apps/nginx/sslkey/test.com/fullchain.crt;
    ssl_certificate_key /apps/nginx/sslkey/test.com/private.key;
    ssl_buffer_size 4k;
    ssl_protocols TLSv1.3 TLSv1.2;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305;
    ssl_prefer_server_ciphers on;
    ssl_ecdh_curve X25519:P-256:P-384;
    client_header_timeout 24h;
    keepalive_timeout 24h;
    location / {
        proxy_redirect     off;
        proxy_set_header   Host $host;
        proxy_set_header   X-Real-IP   $remote_addr;
        proxy_set_header   X-Forwarded-Proto $scheme;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        proxy_ssl_session_reuse off;
        proxy_ssl_server_name on;
        proxy_buffering    off;
        proxy_connect_timeout      90;
        proxy_send_timeout         90;
        proxy_read_timeout         90;
        proxy_buffer_size          4k;
        proxy_buffers              4 32k;
        proxy_busy_buffers_size    64k;
        proxy_http_version 1.1;
        proxy_set_header Accept-Encoding "";
        proxy_pass http://xray-ui;
        #proxy_pass_request_headers on;
        proxy_set_header Connection "keep-alive";
        proxy_store off;
    }
 }

 后端https转发配置参考：

 upstream xray-ui {
        least_conn;
        server 127.0.0.1:54321 max_fails=3 fail_timeout=30s;
        keepalive 1000;
}
server {
    listen 443;
    server_name xray.test.com;
    client_max_body_size 0;
    chunked_transfer_encoding on;
    client_body_buffer_size 202400k;
    client_body_in_single_buffer on;
    add_header Strict-Transport-Security "max-age=63072000; includeSubdomains; preload" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Frame-Options SAMEORIGIN always;
    add_header X-Content-Type-Options nosniff;
    add_header X-Frame-Options "DENY";
    add_header Alt-Svc 'h3=":443"; ma=86400, h3-29=":443"; ma=86400';
    ssl_certificate /apps/nginx/sslkey/test.com/fullchain.crt;
    ssl_certificate_key /apps/nginx/sslkey/test.com/private.key;
    ssl_buffer_size 4k;
    ssl_protocols TLSv1.3 TLSv1.2;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305;
    ssl_prefer_server_ciphers on;
    ssl_ecdh_curve X25519:P-256:P-384;
    client_header_timeout 24h;
    keepalive_timeout 24h;
    location / {
        proxy_redirect     off;
        proxy_set_header   Host $host;
        proxy_set_header   X-Real-IP   $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        proxy_set_header   X-Forwarded-Proto $scheme;
        proxy_ssl_session_reuse off;
        proxy_ssl_server_name on;
        proxy_buffering    off;
        proxy_ssl_name xray.test.com; #证书域名
        # 关闭对后端服务器自签名证书的验证
        proxy_ssl_verify off;
        proxy_connect_timeout      90;
        proxy_send_timeout         90;
        proxy_read_timeout         90;
        proxy_buffer_size          4k;
        proxy_buffers              4 32k;
        proxy_busy_buffers_size    64k;
        proxy_http_version 1.1;
        proxy_set_header Accept-Encoding "";
        proxy_pass https://xray-ui;
        #proxy_pass_request_headers on;
        proxy_set_header Connection "keep-alive";
        proxy_store off;
    }
 }

后端mTLS 转发配置参考：
 upstream xray-ui {
        least_conn;
        server 127.0.0.1:54321 max_fails=3 fail_timeout=30s;
        keepalive 1000;
}
server {
    listen 443;
    server_name xray.test.com;
    client_max_body_size 0;
    chunked_transfer_encoding on;
    client_body_buffer_size 202400k;
    client_body_in_single_buffer on;
    add_header Strict-Transport-Security "max-age=63072000; includeSubdomains; preload" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Frame-Options SAMEORIGIN always;
    add_header X-Content-Type-Options nosniff;
    add_header X-Frame-Options "DENY";
    add_header Alt-Svc 'h3=":443"; ma=86400, h3-29=":443"; ma=86400';
    ssl_certificate /apps/nginx/sslkey/test.com/fullchain.crt;
    ssl_certificate_key /apps/nginx/sslkey/test.com/private.key;
    ssl_buffer_size 4k;
    ssl_protocols TLSv1.3 TLSv1.2;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305;
    ssl_prefer_server_ciphers on;
    ssl_ecdh_curve X25519:P-256:P-384;
    client_header_timeout 24h;
    keepalive_timeout 24h;
    # 添加客户端证书和私钥路径
    ssl_client_certificate /apps/nginx/sslkey/test.com/fullchain.crt;
    ssl_certificate_key /apps/nginx/sslkey/test.com/private.key;

    # 如果需要指定 CA 证书
    # ssl_trusted_certificate /apps/nginx/sslkey/test.com/ca.crt;

    # 强制 SSL/TLS
    proxy_ssl_certificate /apps/nginx/sslkey/test.com/fullchain.crt;
    proxy_ssl_certificate_key /apps/nginx/sslkey/test.com/private.key;
    proxy_ssl_trusted_certificate /apps/nginx/sslkey/test.com/ca.crt;

    # 确保启用 TLS 验证
    proxy_ssl_verify on;
    proxy_ssl_verify_depth 2; # 可根据需要调整
    location / {
        proxy_redirect     off;
        proxy_set_header   Host $host;
        proxy_set_header   X-Real-IP   $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        proxy_set_header   X-Forwarded-Proto $scheme;
        proxy_ssl_session_reuse off;
        proxy_ssl_server_name on;
        proxy_buffering    off;
        proxy_ssl_name xray.test.com; #证书域名
        # 关闭对后端服务器自签名证书的验证
        proxy_ssl_verify off;
        proxy_connect_timeout      90;
        proxy_send_timeout         90;
        proxy_read_timeout         90;
        proxy_buffer_size          4k;
        proxy_buffers              4 32k;
        proxy_busy_buffers_size    64k;
        proxy_http_version 1.1;
        proxy_set_header Accept-Encoding "";
        proxy_pass https://xray-ui;
        #proxy_pass_request_headers on;
        proxy_set_header Connection "keep-alive";
        proxy_store off;
    }
 }
 # vpn代理nginx 配置参考
https://github.com/qist/xray/tree/main/xray/nginx
```
</details>

--------------------------------------------------------------------------------------------------------------------------------------------------

### 关于TG通知（上游内容）

<details>
  <summary>点击查看 关于TG通知</summary>

使用说明:在面板后台设置机器人相关参数

Tg机器人Token

Tg机器人ChatId

#### Tg机器人周期运行时间，采用crontab语法参考语法：

30 * * * * * //每一分的第30s进行通知

@hourly //每小时通知

@daily //每天通知（凌晨零点整）

@every 8h //每8小时通知

@every 30s  //每30s通知一次

#### TG通知内容：

节点流量使用

面板登录提醒

节点到期提醒

流量预警提醒

#### TG机器人可输入内容：

/delete port将会删除对应端口的节点

/restart 将会重启xray服务，该命令不会重启xray-ui面板自身

/status 将会获取当前系统状态

/enable port将会开启对应端口的节点

/disable port将会关闭对应端口的节点

/version 0.1.1.1 xray升级到1.6.0版本

/help 获取帮助信息

</details>
