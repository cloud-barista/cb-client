# cb-client
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cloud-barista/cb-spider?label=go.mod)](https://github.com/cloud-barista/cb-client/blob/main/cb-ctl/go.mod)
[![GoDoc](https://godoc.org/github.com/cloud-barista/cb-client?status.svg)](https://pkg.go.dev/github.com/cloud-barista/cb-client@master)&nbsp;&nbsp;&nbsp;
[![Release Version](https://img.shields.io/github/v/release/cloud-barista/cb-spider)](https://github.com/cloud-barista/cb-client/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/cloud-barista/cb-client/blob/master/LICENSE)


cb-client is the Cloud-Barista CLI and client library.


### [개발 및 실행 환경]

- #### Ubuntu 18.04.5 LTS, Ubuntu 20.04.2 LTS 등

- #### Go 개발 환경 설치
  ```
  $ sudo apt update
  $ sudo apt install -y make gcc
  $ sudo snap install go --classic
  ```
  ` sudo apt update; sudo apt install -y make gcc; sudo snap install go --classic `

- #### 소스 설치 방법
  ```
  $ mkdir -p ~/go/src/github.com/cloud-barista
  $ cd ~/go/src/github.com/cloud-barista
  $ git clone https://github.com/cloud-barista/cb-client.git
  ```
  ` mkdir -p ~/go/src/github.com/cloud-barista; cd ~/go/src/github.com/cloud-barista; git clone https://github.com/cloud-barista/cb-client.git `

### [Cloud-Barista CLI cbctl]
- #### 원격지 터미널 환경에서 Cloud-Barista 운영

- #### cbctl 빌드 방법
  ```
  $ cd cb-client
  $ go build cbctl.go
  ```

- #### 대상 서버 설정 방법
  - 기본 설정 파일 이용한 설정 방법
    - 기본 설정 파일 위치: ./grpc_conf.yaml
    - 대상 서버 설정: server_addr 및 endpoint 정보 설정
    - 그외 설정: 설정 파일 참조
  - 참고: 대상 서버별로 설정 파일 유지 및 적용 가능
    - 실행 방법 참고

- #### cbctl 실행 방법
  - 기본 설정 파일 사용시
    - `$ ./cbctl connect-info list`
  - 특정 설정 파일 사용시: `-c config_file`
    - `$ ./cbctl -c 54.248.3.145.yaml connect-info list`

- #### 로그 설정 방법
  - 설정 파일 기본 위치: ./conf/log_conf.yaml
  - 환경 변수 통한 위치: $CBLOG_ROOT/conf/log_conf.yaml  
  - 로그 설정 참고: [설정 정보](https://github.com/cloud-barista/cb-log#%EC%84%A4%EC%A0%95-%EB%B0%A9%EB%B2%95)
  

***

### [Cloud-Barista Go API]
- #### Go API 활용 예시
  - 클라우드 인프라 연동 정보 통합 관리: [MCIF](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/1.mcif)
  - 멀티 클라우드 네임스페이스 관리: [MCNS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/2.mcns)
  - 멀티 클라우드 인프라 자원 관리: [MCIR](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/3.mcir)
  - 멀티 클라우드 인프라 서비스 관리: [MCIS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/4.mcis)
  - 멀티 클라우드 쿠베르네테스 서비스 관리: [MCKS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/5.mcks)
  - 멀티 클라우드 애플리케이션 운용 관리: [MCAS](https://github.com/cloud-barista/cb-client/tree/main/go-api-examples/6.mcas)
  - 멀티 클라우드 인프라 서비스 통합 모니터링: [MCMON](github.com/cloud-barista/cb-client/tree/main/go-api-examples/7.mcmon)

***

#### [참고: Cloud-Barista 사용자 인터페이스 종류]
-	WebTool: 웹도구 [cb-webtool](https://github.com/cloud-barista/cb-webtool) 참고
-	REST API: [REST 규격 인터페이스](https://github.com/cloud-barista/docs/blob/master/technical_docs/API/CB-User_REST-API.md) 참고
-	Go API: gRPC-based Go API
-	CLI: Cloud-Barista CLI (cbctl)

***
