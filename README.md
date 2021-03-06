# cb-client
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cloud-barista/cb-spider?label=go.mod)](https://github.com/cloud-barista/cb-client/blob/main/cb-ctl/go.mod)
[![GoDoc](https://godoc.org/github.com/cloud-barista/cb-client?status.svg)](https://pkg.go.dev/github.com/cloud-barista/cb-client@master)&nbsp;&nbsp;&nbsp;
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/cloud-barista/cb-client/blob/master/LICENSE)


cb-client is the Cloud-Barista CLI and Go Client Library.

#### [Client 구성]

- ##### cbctl: Cloud-Barista CLI, 터미널 환경에서 Cloud-Barista 운영 제공
- ##### Go API: Cloud-Barista Go API examples, Build시 Client Library 자동 다운로드

*** 

#### [실행 환경]

- ##### 공식환경
  - OS: Ubuntu 20.04
  - Build: Go 1.15

- ##### 시험환경
  - OS: Ubuntu 18.04, Ubuntu 20.04, Debian 10.6, Windows 10, macOS Catalina 10.15 등
  - Build: latest Go

- ##### Go 설치
  ```
  $ sudo apt update
  $ sudo apt install -y make gcc
  $ sudo snap install go --classic
  ```
  ` sudo apt update; sudo apt install -y make gcc; sudo snap install go --classic `

- ##### 소스 다운로드
  ```
  $ mkdir -p ~/go/src/github.com/cloud-barista
  $ cd ~/go/src/github.com/cloud-barista
  $ git clone https://github.com/cloud-barista/cb-client.git
  ```
  ` mkdir -p ~/go/src/github.com/cloud-barista; cd ~/go/src/github.com/cloud-barista; git clone https://github.com/cloud-barista/cb-client.git `

*** 

#### [cbctl 실행 방법]

- ##### cbctl 빌드
  ```
  $ cd cb-client/cb-ctl
  $ go build cbctl.go
  ```

- ##### 대상 서버 설정
  - 기본 설정 파일 이용한 설정 방법
    - 기본 설정 파일 위치: ./grpc_conf.yaml
    - 대상 서버 설정: server_addr 및 endpoint 정보 설정
  - 지정 설정 파일 이용한 설정 방법
    - `-c config_file` 옵션 활용

- ##### cbctl 실행
  - Help
    - `$ ./cbctl`
    - `$ ./cbctl driver`
  - 기본 설정 파일 사용 방법
    - `$ ./cbctl connect-info list`
  - 특정 설정 파일 사용 방법: `-c config_file`
    - `$ ./cbctl -c 123.456.789.10.yaml connect-info list`

- ##### 로그 설정 방법
  - 설정 파일 기본 위치: ./conf/log_conf.yaml
  - 환경 변수 통한 위치: $CBLOG_ROOT/conf/log_conf.yaml  
  - 로그 설정 참고: [설정 정보](https://github.com/cloud-barista/cb-log#%EC%84%A4%EC%A0%95-%EB%B0%A9%EB%B2%95)
  

***

#### [Go API 활용 방법]

- ##### Go API 활용 예시
  - 클라우드 인프라 연동 정보 통합 관리: [MCIF](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/1.mcif)
  - 멀티 클라우드 네임스페이스 관리: [MCNS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/2.mcns)
  - 멀티 클라우드 인프라 자원 관리: [MCIR](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/3.mcir)
  - 멀티 클라우드 인프라 서비스 관리: [MCIS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/4.mcis)
  - 멀티 클라우드 쿠베르네테스 서비스 관리: [MCKS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/5.mcks)
  - 멀티 클라우드 애플리케이션 운용 관리: [MCAS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/6.mcas)
  - 멀티 클라우드 인프라 서비스 통합 모니터링: [MCMON](github.com/cloud-barista/cb-client/tree/main/go-api-examples/7.mcmon)

***

#### [참고] Cloud-Barista 사용자 인터페이스 종류

-	##### WebTool: 웹도구 [cb-webtool](https://github.com/cloud-barista/cb-webtool) 참고
-	##### REST API: [REST 규격 인터페이스](https://github.com/cloud-barista/docs/blob/master/technical_docs/API/CB-User_REST-API.md) 참고
-	##### CLI: Cloud-Barista CLI (cbctl)
-	##### Go API: Cloud-Barista Go API(gRPC-based)

***
