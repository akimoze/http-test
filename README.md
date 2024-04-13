# http-test

HTTPのテストサーバー

## ダウンロード

### linux

#### amd64

```
wget https://
```

#### arm64

```
wget https://
```

### mac

#### arm64

```
wget https://
```

## 設定ファイル

`config.toml`をバイナリと同じディレクトリに配置してください。

```toml
address = ":3000"
use_tls = true

[cert]
cert_file = "./cert/example.com.pem"
cert_key = "./cert/example.com.key"
```

| key | type | description |
| -- | -- | -- |
| address | string | listenAddressを指定します |
| use_tls | bool | SSL/TLSを使用するかどうかを指定します。 |
| cert.cert_file | string | SSLの証明書ファイルのパスを指定します(SSL/TLS使用時は必須) |
| cert.cert_key | string | SSLの秘密鍵のパスを指定します(SSL/TLS使用時は必須) |
