#!/bin/bash
#yum install nginx wget -y
# 断点续传
set -e
wget -c https://github.com/xishengcai/cloud/releases/download/v1.0.0/proxy-server -O /usr/local/bin/proxy-server
wget -c https://github.com/xishengcai/cloud/releases/download/v1.0.0/v2ctl  -O  /usr/local/bin/v2ctl
chmod +x /usr/local/bin/proxy-server
chmod +x /usr/local/bin/v2ctl

# 写入nginx 配置
echo 'server {
    listen       {{ .ExternalPort }} ssl http2;
    server_name  {{ .CommonName }};
    charset utf-8;

    # ssl配置
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
    ssl_ecdh_curve secp384r1;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    ssl_session_tickets off;

    ssl_certificate /opt/proxy-cert/server.crt;
    ssl_certificate_key /opt/proxy-cert/server.key;

    access_log  /var/log/nginx/proxy.access.log;
    error_log /var/log/nginx/proxy.error.log;

    root /usr/share/nginx/html;

    location / {
      proxy_redirect off;
      proxy_pass http://127.0.0.1:{{ .ProxyPort }};
      proxy_http_version 1.1;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
      proxy_set_header Host $host;
      # Show real IP in proxy access.log
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}' > /etc/nginx/conf.d/proxy-1234.conf

# 写入代理 配置
mkdir -p /etc/proxy
echo '{
  "inbounds": [{
    "port": {{ .ProxyPort }},
    "protocol": "vmess",
    "settings": {
      "clients": [
        {
          "id": "100f267d-0d53-689c-a773-afd0047fa6b0",
          "level": 1,
          "alterId": 64
        }
      ]
    },
    "streamSettings": {
      "network": "ws",
      "wsSettings": {
        "path": "/"
      }
    }
    }],
  "outbounds": [{
    "protocol": "freedom",
    "settings": {}
  }]
}' > /etc/proxy/vmess-tls-ws.json

echo '[Unit]
      Description=proxy-server
      After=network.target

      [Service]
      Type=simple
      User=root
      ExecStart=/usr/local/bin/proxy-server --config /etc/proxy/vmess-tls-ws.json

      [Install]
      WantedBy=multi-user.target' > /etc/systemd/system/proxy-server.service

sh /root/gencert --CN {{ .CommonName }} --dir /opt/proxy-cert
systemctl daemon-reload
systemctl start proxy-server
systemctl start nginx