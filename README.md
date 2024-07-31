# Cloud Run Sample 1

## Goモジュールの初期化

```bash:
-- Goモジュールの初期化
% go mod init helloworld

-- Goのバージョン変更（go 1.22.5 <-- 1.20）
% vi go.mod
  --------------------------------------------------
  module helloworld

  go 1.20
  --------------------------------------------------
```

## Google Cloud Platform の設定
```
-- 次のページで「Cloud Run」「Cloud Build」を有効にする
https://console.cloud.google.com/cloud-build/settings/service-account?hl=ja&project=topodaisy
```


## リリース方法
```bash:
-- Google Cloud プロジェクトの設定
% gcloud init

-- ユーザーアカウントのログイン
% gcloud auth login

-- コンテナイメージのビルド
% gcloud builds submit --tag gcr.io/topodaisy/helloworld

-- Cloud Run にデプロイ
% gcloud run deploy helloworld --image gcr.io/topodaisy/helloworld --platform managed   
```


## 動作確認
```
-- ブラウザで次のURLにアクセス
https://helloworld-qxhavcoa2a-de.a.run.app
```


## ローカルでの実行

### 初期設定
```bash:
-- .envを作成
% cp .env.default .env

-- 自身のローカル環境に合わせて.envを設定
% code .env
```

### docker-compose実行
```bash:
% docker-compose -up -d
```

### 動作確認
```
-- ブラウザで次のURLにアクセス
http://localhost:8080/
```


## 初回リリース履歴

・Goモジュールの初期化
```bash:
% go mod init helloworld
  --------------------------------------------------
  go: creating new go.mod: module helloworld
  go: to add module requirements and sums:
          go mod tidy
  --------------------------------------------------
```

・Google Cloud プロジェクトの設定
```bash:
% gcloud init
  --------------------------------------------------
  Welcome! This command will take you through the configuration of gcloud.

  Settings from your current configuration [default] are:
  core:
    account: murayama@orb-japan.co.jp
    disable_usage_reporting: 'True'

  Pick configuration to use:
  [1] Re-initialize this configuration [default] with new settings 
  [2] Create a new configuration
  Please enter your numeric choice:  
  Please enter a value between 1 and 2:  1

  Your current configuration has been set to: [default]

  You can skip diagnostics next time by using the following flag:
    gcloud init --skip-diagnostics

  Network diagnostic detects and fixes local network connection issues.
  Checking network connection...done.                                                                                                                  
  Reachability Check passed.
  Network diagnostic passed (1/1 checks passed).

  Choose the account you would like to use to perform operations for this configuration:
  [1] murayama@orb-japan.co.jp
  [2] Log in with a new account
  Please enter your numeric choice:  1

  You are logged in as: [murayama@orb-japan.co.jp].

  Pick cloud project to use: 
  [1] big-query-test-257303
  [2] chatbot-demo-4a679
  [3] ec-app-aba1b
  [4] topodaisy
  [5] Enter a project ID
  [6] Create a new project
  Please enter numeric choice or text value (must exactly match list item):  4

  Your current project has been set to: [topodaisy].

  Your project default Compute Engine zone has been set to [us-central1-f].
  You can change it by running [gcloud config set compute/zone NAME].

  Your project default Compute Engine region has been set to [us-central1].
  You can change it by running [gcloud config set compute/region NAME].

  Created a default .boto configuration file at [/Users/murayama/.boto]. See this file and
  [https://cloud.google.com/storage/docs/gsutil/commands/config] for more
  information about configuring Google Cloud Storage.
  Your Google Cloud SDK is configured and ready to use!

  * Commands that require authentication will use murayama@orb-japan.co.jp by default
  * Commands will reference project `topodaisy` by default
  * Compute Engine commands will use region `us-central1` by default
  * Compute Engine commands will use zone `us-central1-f` by default

  Run `gcloud help config` to learn how to change individual settings

  This gcloud configuration is called [default]. You can create additional configurations if you work with multiple accounts and/or projects.
  Run `gcloud topic configurations` to learn more.

  Some things to try next:

  * Run `gcloud --help` to see the Cloud Platform services you can interact with. And run `gcloud help COMMAND` to get help on any gcloud command.
  * Run `gcloud topic --help` to learn about advanced features of the SDK like arg files and output formatting
  * Run `gcloud cheat-sheet` to see a roster of go-to `gcloud` commands.
  --------------------------------------------------
```

・ユーザーアカウントの許可
```bash:
% gcloud auth login
  --------------------------------------------------
  Your browser has been opened to visit:

      https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=32555940559.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A8085%2F&scope=openid+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcloud-platform+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fappengine.admin+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fsqlservice.login+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcompute+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Faccounts.reauth&state=j0Ij3IlUMdYfzS8zPgasRB8bET4O2M&access_type=offline&code_challenge=cFMAiuz1pGZMgR5ORRzAVhnRAC_tnZUJxBdihiE-xfs&code_challenge_method=S256


  You are now logged in as [murayama@orb-japan.co.jp].
  Your current project is [topodaisy].  You can change this setting by running:
    $ gcloud config set project PROJECT_ID
  --------------------------------------------------
```

・コンテナイメージのビルド1
```bash:
% gcloud builds submit --tag gcr.io/topodaisy/helloworld
  --------------------------------------------------
  Creating temporary archive of 2 file(s) totalling 567 bytes before compression.
  Uploading tarball of [.] to [gs://topodaisy_cloudbuild/source/1722205750.073117-05011cb4859a4d17af1264f588cde470.tgz]
  Created [https://cloudbuild.googleapis.com/v1/projects/topodaisy/locations/global/builds/7906c4ac-d155-4624-9435-ec6d96ed2220].
  Logs are available at [ https://console.cloud.google.com/cloud-build/builds/7906c4ac-d155-4624-9435-ec6d96ed2220?project=823120530617 ].
  Waiting for build to complete. Polling interval: 1 second(s).
  ---------------------------------------------------------------- REMOTE BUILD OUTPUT -----------------------------------------------------------------
  starting build "7906c4ac-d155-4624-9435-ec6d96ed2220"

  FETCHSOURCE
  Fetching storage object: gs://topodaisy_cloudbuild/source/1722205750.073117-05011cb4859a4d17af1264f588cde470.tgz#1722205751284021
  Copying gs://topodaisy_cloudbuild/source/1722205750.073117-05011cb4859a4d17af1264f588cde470.tgz#1722205751284021...
  / [1 files][  633.0 B/  633.0 B]                                                
  Operation completed over 1 objects/633.0 B.
  BUILD
  Already have image (with digest): gcr.io/cloud-builders/docker
  Sending build context to Docker daemon  3.072kB
  Step 1/9 : FROM golang:latest
  latest: Pulling from library/golang
  ca4e5d672725: Already exists
  30b93c12a9c9: Pulling fs layer
  10d643a5fa82: Pulling fs layer
  5e53f8e7a11a: Pulling fs layer
  32a2f51ff3dd: Pulling fs layer
  ddb1467b736f: Pulling fs layer
  4f4fb700ef54: Pulling fs layer
  32a2f51ff3dd: Waiting
  ddb1467b736f: Waiting
  4f4fb700ef54: Waiting
  30b93c12a9c9: Verifying Checksum
  30b93c12a9c9: Download complete
  10d643a5fa82: Verifying Checksum
  10d643a5fa82: Download complete
  5e53f8e7a11a: Verifying Checksum
  5e53f8e7a11a: Download complete
  ddb1467b736f: Verifying Checksum
  ddb1467b736f: Download complete
  4f4fb700ef54: Verifying Checksum
  4f4fb700ef54: Download complete
  32a2f51ff3dd: Verifying Checksum
  32a2f51ff3dd: Download complete
  30b93c12a9c9: Pull complete
  10d643a5fa82: Pull complete
  5e53f8e7a11a: Pull complete
  32a2f51ff3dd: Pull complete
  ddb1467b736f: Pull complete
  4f4fb700ef54: Pull complete
  Digest: sha256:86a3c48a61915a8c62c0e1d7594730399caa3feb73655dfe96c7bc17710e96cf
  Status: Downloaded newer image for golang:latest
  ---> a61b645b609b
  Step 2/9 : LABEL maintainer="Your Name <your.email@example.com>"
  ---> Running in 9af5eb5d5b4a
  Removing intermediate container 9af5eb5d5b4a
  ---> 83f9f0114cc5
  Step 3/9 : WORKDIR /app
  ---> Running in 2e3259e4a9f5
  Removing intermediate container 2e3259e4a9f5
  ---> c6ec7577413f
  Step 4/9 : COPY go.mod ./
  COPY failed: file not found in build context or excluded by .dockerignore: stat go.mod: file does not exist
  ERROR
  ERROR: build step 0 "gcr.io/cloud-builders/docker" failed: step exited with non-zero status: 1
  ------------------------------------------------------------------------------------------------------------------------------------------------------

  BUILD FAILURE: Build step failure: build step 0 "gcr.io/cloud-builders/docker" failed: step exited with non-zero status: 1
  ERROR: (gcloud.builds.submit) build 7906c4ac-d155-4624-9435-ec6d96ed2220 completed with status "FAILURE"
  --------------------------------------------------
```


・次のページで「Cloud Run」「Cloud Build」を有効にする
```
https://console.cloud.google.com/cloud-build/settings/service-account?hl=ja&project=topodaisy
```


・コンテナイメージのビルド2
```bash:
% gcloud builds submit --tag gcr.io/topodaisy/helloworld
  --------------------------------------------------
  Creating temporary archive of 4 file(s) totalling 8.8 KiB before compression.
  Uploading tarball of [.] to [gs://topodaisy_cloudbuild/source/1722207987.189063-a64a8f0a00404962b4e908704e94367b.tgz]
  Created [https://cloudbuild.googleapis.com/v1/projects/topodaisy/locations/global/builds/5a3118cd-47d1-4b81-bc33-bdf8fb605abd].
  Logs are available at [ https://console.cloud.google.com/cloud-build/builds/5a3118cd-47d1-4b81-bc33-bdf8fb605abd?project=823120530617 ].
  Waiting for build to complete. Polling interval: 1 second(s).
  ---------------------------------------------------------------- REMOTE BUILD OUTPUT -----------------------------------------------------------------
  starting build "5a3118cd-47d1-4b81-bc33-bdf8fb605abd"

  FETCHSOURCE
  Fetching storage object: gs://topodaisy_cloudbuild/source/1722207987.189063-a64a8f0a00404962b4e908704e94367b.tgz#1722207988030855
  Copying gs://topodaisy_cloudbuild/source/1722207987.189063-a64a8f0a00404962b4e908704e94367b.tgz#1722207988030855...
  / [1 files][  3.7 KiB/  3.7 KiB]                                                
  Operation completed over 1 objects/3.7 KiB.
  BUILD
  Already have image (with digest): gcr.io/cloud-builders/docker
  Sending build context to Docker daemon  13.31kB
  Step 1/8 : FROM golang:1.20 AS builder
  1.20: Pulling from library/golang
  6a299ae9cfd9: Pulling fs layer
  e08e8703b2fb: Pulling fs layer
  68e92d11b04e: Pulling fs layer
  4105062d1ee6: Pulling fs layer
  870a28135dd0: Pulling fs layer
  ffe80ff75448: Pulling fs layer
  4f4fb700ef54: Pulling fs layer
  4105062d1ee6: Waiting
  870a28135dd0: Waiting
  ffe80ff75448: Waiting
  4f4fb700ef54: Waiting
  e08e8703b2fb: Verifying Checksum
  e08e8703b2fb: Download complete
  6a299ae9cfd9: Download complete
  68e92d11b04e: Verifying Checksum
  68e92d11b04e: Download complete
  ffe80ff75448: Verifying Checksum
  ffe80ff75448: Download complete
  4f4fb700ef54: Verifying Checksum
  4f4fb700ef54: Download complete
  4105062d1ee6: Verifying Checksum
  4105062d1ee6: Download complete
  870a28135dd0: Verifying Checksum
  870a28135dd0: Download complete
  6a299ae9cfd9: Pull complete
  e08e8703b2fb: Pull complete
  68e92d11b04e: Pull complete
  4105062d1ee6: Pull complete
  870a28135dd0: Pull complete
  ffe80ff75448: Pull complete
  4f4fb700ef54: Pull complete
  Digest: sha256:8f9af7094d0cb27cc783c697ac5ba25efdc4da35f8526db21f7aebb0b0b4f18a
  Status: Downloaded newer image for golang:1.20
  ---> d5beeac3653f
  Step 2/8 : WORKDIR /app
  ---> Running in 4aec95f2cbc1
  Removing intermediate container 4aec95f2cbc1
  ---> bc181913bde2
  Step 3/8 : COPY . .
  ---> 48ad51c9157d
  Step 4/8 : RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server
  ---> Running in ffba7e70b763
  go: errors parsing go.mod:
  /app/go.mod:3: invalid go version '1.22.5': must match format 1.23
  The command '/bin/sh -c CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server' returned a non-zero code: 1
  ERROR
  ERROR: build step 0 "gcr.io/cloud-builders/docker" failed: step exited with non-zero status: 1

  ------------------------------------------------------------------------------------------------------------------------------------------------------

  BUILD FAILURE: Build step failure: build step 0 "gcr.io/cloud-builders/docker" failed: step exited with non-zero status: 1
  ERROR: (gcloud.builds.submit) build 5a3118cd-47d1-4b81-bc33-bdf8fb605abd completed with status "FAILURE"
  --------------------------------------------------
```


・バージョン変更（go 1.22.5 <-- 1.20）
```bash:
% vi go.mod
  --------------------------------------------------
  module helloworld

  go 1.20
  --------------------------------------------------
```


・コンテナイメージのビルド3
```bash:
% gcloud builds submit --tag gcr.io/topodaisy/helloworld
  --------------------------------------------------
  Creating temporary archive of 4 file(s) totalling 8.8 KiB before compression.
  Uploading tarball of [.] to [gs://topodaisy_cloudbuild/source/1722208422.543556-d0ce396ef4fc4adb822ca4b3ad39a7e4.tgz]
  Created [https://cloudbuild.googleapis.com/v1/projects/topodaisy/locations/global/builds/73150799-4b8e-4965-8f58-a4b8f1d10c9f].
  Logs are available at [ https://console.cloud.google.com/cloud-build/builds/73150799-4b8e-4965-8f58-a4b8f1d10c9f?project=823120530617 ].
  Waiting for build to complete. Polling interval: 1 second(s).
  ---------------------------------------------------------------- REMOTE BUILD OUTPUT -----------------------------------------------------------------
  starting build "73150799-4b8e-4965-8f58-a4b8f1d10c9f"

  FETCHSOURCE
  Fetching storage object: gs://topodaisy_cloudbuild/source/1722208422.543556-d0ce396ef4fc4adb822ca4b3ad39a7e4.tgz#1722208423682528
  Copying gs://topodaisy_cloudbuild/source/1722208422.543556-d0ce396ef4fc4adb822ca4b3ad39a7e4.tgz#1722208423682528...
  / [1 files][  3.7 KiB/  3.7 KiB]                                                
  Operation completed over 1 objects/3.7 KiB.
  BUILD
  Already have image (with digest): gcr.io/cloud-builders/docker
  Sending build context to Docker daemon  13.31kB
  Step 1/8 : FROM golang:1.20 AS builder
  1.20: Pulling from library/golang
  6a299ae9cfd9: Pulling fs layer
  e08e8703b2fb: Pulling fs layer
  68e92d11b04e: Pulling fs layer
  4105062d1ee6: Pulling fs layer
  870a28135dd0: Pulling fs layer
  ffe80ff75448: Pulling fs layer
  4f4fb700ef54: Pulling fs layer
  4105062d1ee6: Waiting
  870a28135dd0: Waiting
  ffe80ff75448: Waiting
  4f4fb700ef54: Waiting
  e08e8703b2fb: Verifying Checksum
  e08e8703b2fb: Download complete
  6a299ae9cfd9: Verifying Checksum
  6a299ae9cfd9: Download complete
  68e92d11b04e: Verifying Checksum
  68e92d11b04e: Download complete
  ffe80ff75448: Verifying Checksum
  ffe80ff75448: Download complete
  4f4fb700ef54: Verifying Checksum
  4f4fb700ef54: Download complete
  870a28135dd0: Verifying Checksum
  870a28135dd0: Download complete
  4105062d1ee6: Verifying Checksum
  4105062d1ee6: Download complete
  6a299ae9cfd9: Pull complete
  e08e8703b2fb: Pull complete
  68e92d11b04e: Pull complete
  4105062d1ee6: Pull complete
  870a28135dd0: Pull complete
  ffe80ff75448: Pull complete
  4f4fb700ef54: Pull complete
  Digest: sha256:8f9af7094d0cb27cc783c697ac5ba25efdc4da35f8526db21f7aebb0b0b4f18a
  Status: Downloaded newer image for golang:1.20
  ---> d5beeac3653f
  Step 2/8 : WORKDIR /app
  ---> Running in ad23ad1918c7
  Removing intermediate container ad23ad1918c7
  ---> 820bd4db8ba6
  Step 3/8 : COPY . .
  ---> 19b748c7577e
  Step 4/8 : RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server
  ---> Running in 6dcaa41abd7d
  internal/unsafeheader
  internal/goarch
  internal/cpu
  internal/abi
  internal/coverage/rtcov
  internal/bytealg
  internal/goexperiment
  internal/goos
  runtime/internal/atomic
  runtime/internal/math
  runtime/internal/sys
  runtime/internal/syscall
  internal/itoa
  math/bits
  runtime
  math
  unicode/utf8
  internal/race
  sync/atomic
  unicode
  container/list
  crypto/internal/alias
  crypto/subtle
  crypto/internal/boring/sig
  unicode/utf16
  vendor/golang.org/x/crypto/cryptobyte/asn1
  internal/nettrace
  vendor/golang.org/x/crypto/internal/alias
  internal/reflectlite
  sync
  internal/testlog
  internal/godebug
  math/rand
  internal/singleflight
  internal/intern
  errors
  sort
  strconv
  io
  internal/oserror
  syscall
  reflect
  internal/syscall/unix
  time
  internal/poll
  internal/safefilepath
  internal/syscall/execenv
  path
  internal/fmtsort
  io/fs
  bytes
  os
  strings
  bufio
  fmt
  encoding/binary
  hash
  hash/crc32
  log
  compress/flate
  context
  crypto
  crypto/cipher
  compress/gzip
  crypto/internal/boring
  crypto/internal/randutil
  math/big
  crypto/aes
  crypto/des
  crypto/internal/edwards25519/field
  crypto/internal/nistec/fiat
  crypto/rand
  embed
  crypto/internal/bigmod
  crypto/internal/nistec
  crypto/internal/boring/bbig
  crypto/sha512
  encoding/asn1
  crypto/ecdh
  crypto/elliptic
  vendor/golang.org/x/crypto/cryptobyte
  crypto/internal/edwards25519
  crypto/ecdsa
  crypto/ed25519
  crypto/hmac
  crypto/md5
  crypto/rc4
  crypto/rsa
  crypto/sha1
  crypto/sha256
  crypto/dsa
  encoding/hex
  encoding/base64
  crypto/x509/pkix
  vendor/golang.org/x/net/dns/dnsmessage
  encoding/pem
  net/netip
  net/url
  net
  path/filepath
  vendor/golang.org/x/crypto/chacha20
  vendor/golang.org/x/crypto/internal/poly1305
  io/ioutil
  vendor/golang.org/x/sys/cpu
  vendor/golang.org/x/crypto/chacha20poly1305
  vendor/golang.org/x/crypto/hkdf
  vendor/golang.org/x/text/transform
  vendor/golang.org/x/text/unicode/bidi
  vendor/golang.org/x/text/secure/bidirule
  vendor/golang.org/x/text/unicode/norm
  vendor/golang.org/x/net/idna
  crypto/x509
  net/textproto
  vendor/golang.org/x/net/http/httpguts
  vendor/golang.org/x/net/http/httpproxy
  vendor/golang.org/x/net/http2/hpack
  crypto/tls
  mime
  mime/quotedprintable
  mime/multipart
  net/http/internal
  net/http/internal/ascii
  net/http/httptrace
  net/http
  helloworld
  Removing intermediate container 6dcaa41abd7d
  ---> 6e0d126f756c
  Step 5/8 : FROM alpine:3
  3: Pulling from library/alpine
  Digest: sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5
  Status: Downloaded newer image for alpine:3
  ---> 324bc02ae123
  Step 6/8 : RUN apk add --no-cache ca-certificates
  ---> Running in ddec4abed5c4
  fetch https://dl-cdn.alpinelinux.org/alpine/v3.20/main/x86_64/APKINDEX.tar.gz
  fetch https://dl-cdn.alpinelinux.org/alpine/v3.20/community/x86_64/APKINDEX.tar.gz
  (1/1) Installing ca-certificates (20240705-r0)
  Executing busybox-1.36.1-r29.trigger
  Executing ca-certificates-20240705-r0.trigger
  OK: 8 MiB in 15 packages
  Removing intermediate container ddec4abed5c4
  ---> 37f04dbee273
  Step 7/8 : COPY --from=builder /app/server /server
  ---> b2b4f40fcca5
  Step 8/8 : CMD ["/server"]
  ---> Running in 7989caa2d3d6
  Removing intermediate container 7989caa2d3d6
  ---> 830008158d34
  Successfully built 830008158d34
  Successfully tagged gcr.io/topodaisy/helloworld:latest
  PUSH
  Pushing gcr.io/topodaisy/helloworld
  The push refers to repository [gcr.io/topodaisy/helloworld]
  3bc5fe6de2ba: Preparing
  9c6fea0baa58: Preparing
  78561cef0761: Preparing
  3bc5fe6de2ba: Pushed
  9c6fea0baa58: Pushed
  78561cef0761: Pushed
  latest: digest: sha256:8f670eb8c9439f7c53164d07fbeeaee6ffe40ac94f996e30be5a03def74c1162 size: 949
  DONE
  ------------------------------------------------------------------------------------------------------------------------------------------------------
  ID                                    CREATE_TIME                DURATION  SOURCE                                                                                   IMAGES                                 STATUS
  73150799-4b8e-4965-8f58-a4b8f1d10c9f  2024-07-28T23:13:44+00:00  1M4S      gs://topodaisy_cloudbuild/source/1722208422.543556-d0ce396ef4fc4adb822ca4b3ad39a7e4.tgz  gcr.io/topodaisy/helloworld (+1 more)  SUCCESS
  --------------------------------------------------
```


・Cloud Run にデプロイ
```bash:
% gcloud run deploy helloworld --image gcr.io/topodaisy/helloworld --platform managed   
  --------------------------------------------------
  Please specify a region:
  [1] africa-south1
  [2] asia-east1
  [3] asia-east2
  [4] asia-northeast1
  [5] asia-northeast2
  [6] asia-northeast3
  [7] asia-south1
  [8] asia-south2
  [9] asia-southeast1
  [10] asia-southeast2
  [11] australia-southeast1
  [12] australia-southeast2
  [13] europe-central2
  [14] europe-north1
  [15] europe-southwest1
  [16] europe-west1
  [17] europe-west10
  [18] europe-west12
  [19] europe-west2
  [20] europe-west3
  [21] europe-west4
  [22] europe-west6
  [23] europe-west8
  [24] europe-west9
  [25] me-central1
  [26] me-central2
  [27] me-west1
  [28] northamerica-northeast1
  [29] northamerica-northeast2
  [30] southamerica-east1
  [31] southamerica-west1
  [32] us-central1
  [33] us-east1
  [34] us-east4
  [35] us-east5
  [36] us-south1
  [37] us-west1
  [38] us-west2
  [39] us-west3
  [40] us-west4
  [41] cancel
  Please enter numeric choice or text value (must exactly match list item):  2

  To make this the default region, run `gcloud config set run/region asia-east1`.

  Allow unauthenticated invocations to [helloworld] (y/N)?  y

  Deploying container to Cloud Run service [helloworld] in project [topodaisy] region [asia-east1]
  ✓ Deploying new service... Done.                                                                                                                     
    ✓ Creating Revision...                                                                                                                             
    ✓ Routing traffic...                                                                                                                               
    ✓ Setting IAM Policy...                                                                                                                            
  Done.                                                                                                                                                
  Service [helloworld] revision [helloworld-00001-hvb] has been deployed and is serving 100 percent of traffic.
  Service URL: https://helloworld-qxhavcoa2a-de.a.run.app
  --------------------------------------------------
```

・確認  
```bash:
  -- ブラウザで次のURLにアクセス
  https://helloworld-qxhavcoa2a-de.a.run.app
```

