# ユーザーガイド
## 目次
- kkok サーバーの起動	
- kkok クライアントの使い方
- REST API

## kkokサーバーの起動
`kkok` コマンドにより kkok サーバーを起動します。

使用可能なフラグは以下の通りです。

```
-f 'Configuration file name'
	設定ファイルのパスを指定します。
-v
	コンパイルされているコマンドのバージョンをプリントし、プログラムを終了します。
-test
	設定ファイルのテストを行います。
```

kkok サーバーは、サーバーの設定を行うために起動時に指定したパスから [TOML 形式](https://github.com/toml-lang/toml)の設定ファイルを読み込みます（デフォルトでは `/etc/kkok.toml` です）。設定ファイルの例は [https://github.com/cybozu-go/kkok/blob/master/cmd/kkok/sample.toml](https://github.com/cybozu-go/kkok/cmd/kkok/sample.toml)をご覧ください。

## kkokサーバーの使い方
`kkokc` コマンドによりkkokサーバーにリクエストを送る kkok クライアントを起動します。

使用可能なフラグは以下の通りです。

```
-url 'URL'
	kkok サーバーのURLを指定します。デフォルトは `http://localhost:19898` です。
-token 'Authentication token'
	kkok サーバーの認証に使用するトークンを指定します。デフォルトでは環境変数 `TokenEnv` が使用されます。
```

## REST API
kkok サーバーにリクエストを送るための REST API です。詳しくは[こちら](https://github.com/cybozu-go/kkok/blob/master/docs/API.md)をご覧ください。
