# 生态系统
1. Web服务器：Nginx
2. 应用服务器
3. 进程监控器

# Nginx
1. Web服务器
2. 反向代理服务器

```bash
docker run --name nginxServer -d -p 80:80 nginx
docker run --name nginxServer -d -p 80:80 --mount type=bind,source=$(pwd)/nginx.conf,destination=/etc/nginx/nginx.conf:readonly nginx

docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d nginx
```

> Nginx 也是一个上游服务器。上游服务器负责将一个服务器的请求转发到另一个服务器。


# Supervisord
```bash
sudo apt-get install -y supervisor
```
