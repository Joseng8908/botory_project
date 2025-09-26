# --- 1단계: 빌드 전용 환경 (Builder) ---
# Go 컴파일러가 포함된 공식 이미지를 기반으로 시작합니다.
FROM golang:1.25-alpine AS builder

# 작업 디렉터리를 설정합니다.
WORKDIR /app

# 의존성 캐싱을 위해 go.mod와 go.sum 파일을 먼저 복사합니다.
COPY go.mod go.sum ./

# 의존성을 다운로드합니다.
RUN go mod download

# 나머지 모든 소스 코드를 복사합니다.
COPY . .

# 애플리케이션을 빌드합니다.
# CGO_ENABLED=0: C 라이브러리 없이 순수 Go 코드로 빌드하여, 어떤 리눅스 환경에서도 실행 가능하게 만듭니다.
# -o ./bin/botory: 결과물을 bin/botory 라는 이름으로 저장합니다.
RUN CGO_ENABLED=0 go build -o ./bin/botory ./cmd/botory


# --- 2단계: 최종 실행 환경 (Final) ---
# 아주 가벼운 'alpine' 이미지를 기반으로 시작합니다. Go 컴파일러 등은 포함되지 않습니다.
FROM alpine:latest

# 작업 디렉터리를 설정합니다.
WORKDIR /app

# 빌더(builder) 환경에서 만들어진 실행 파일만 복사해옵니다.
COPY --from=builder /app/bin/botory .

# 챗봇 설정 파일들이 들어있는 configs 폴더를 복사합니다.
COPY configs/ ./configs/

# 컨테이너의 8080 포트를 외부에 노출시킬 것임을 명시합니다.
EXPOSE 8080

# 컨테이너가 시작될 때 실행할 기본 명령어를 설정합니다.
# ["./botory", "start"]는 터미널에서 ./botory start 를 실행하는 것과 같습니다.
CMD ["./botory", "start"]