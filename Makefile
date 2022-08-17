BUILD=go build
LDFLAGS=-X main.buildType=debug
DATE=$(shell date '+%s')
GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
GOARM=
ifeq "$(strip $(GOOS))" "darwin"
NPROC=$(shell sysctl -n hw.ncpu)
else
NPROC=$(shell nproc)
endif
ifndef OUTSUFFIX
OUTSUFFIX=${GOOS}-${GOARCH}
endif
ifndef OUT
OUT=$(shell pwd)/build/$(shell basename $(PWD))_${OUTSUFFIX}
endif

# Text coloring & styling
BOLD=\033[1m
UNDERLINE=\033[4m
HEADER=${BOLD}${UNDERLINE}

GREEN=\033[38;5;118m
RED=\033[38;5;196m
GREY=\033[38;5;250m

RESET=\033[m

# Targets
all:
	make -i aix ppc64 release build
	make -i android 386 release build
	make -i android amd64 release build
	make -i android arm release build
	make -i android arm64 release build
	make -i darwin amd64 release build
	make -i darwin arm64 release build
	make -i dragonfly amd64 release build
	make -i freebsd 386 release build
	make -i freebsd amd64 release build
	make -i freebsd arm release build
	make -i freebsd arm64 release build
	make -i illumos amd64 release build
	make -i ios amd64 release build
	make -i js wasm release build
	make -i linux 386 release build
	make -i linux amd64 release build
	make -i linux arm release build
	make -i linux arm64 release build
	make -i linux mips release build
	make -i linux mips64 release build
	make -i linux mips64le release build
	make -i linux mipsle release build
	make -i linux ppc64 release build
	make -i linux ppc64le release build
	make -i linux riscv64 release build
	make -i linux s390x release build
	make -i netbsd 386 release build
	make -i netbsd amd64 release build
	make -i netbsd arm release build
	make -i netbsd arm64 release build
	make -i openbsd 386 release build
	make -i openbsd amd64 release build
	make -i openbsd arm release build
	make -i openbsd arm64 release build
	make -i openbsd mips64 release build
	make -i plan9 386 release build
	make -i plan9 amd64 release build
	make -i plan9 arm release build
	make -i solaris amd64 release build
	make -i 386 windows release build
	make -i amd64 windows release build
	make -i arm windows release build
	make -i arm64 windows release build

l: lint
lint:
	@printf "${GREEN}${HEADER}Linting${RESET}\n"
	go vet ./...

d: deps
deps: dependencies
dependencies:
	@printf "${GREEN}${HEADER}Downloading dependencies${RESET}\n"
	go get ./...

f: format
format:
	@printf "${GREEN}${HEADER}Formatting${RESET}\n"
	go fmt ./...

t: test
test:
	@printf "${GREEN}${HEADER}Starting test suite${RESET}\n"
	go test -parallel ${NPROC} ./...

release:
	$(eval LDFLAGS=-w -s -X main.buildType=release)
	@:

run:
	go run main.go

b: build 
build: clean
	$(eval LDFLAGS=${LDFLAGS} -X main.buildVersion=${DATE})
	@printf "${GREEN}${HEADER}Compiling for ${GOARCH}-${GOOS} to '${OUT}'${RESET}\n"
	mkdir -p $(shell dirname ${OUT})
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} ${BUILD} -p ${NPROC} -ldflags="${LDFLAGS}" -o ${OUT}
	chmod +x ${OUT}
	upx -9 ${OUT}

clean:
	@printf "${GREEN}${HEADER}Cleaning previous build${RESET}\n"
	rm -rf ${OUT}

# OS presets
aix:
	$(eval GOOS=aix)
	@:
android:
	$(eval GOOS=android)
	@:
darwin:
	$(eval GOOS=darwin)
	@:
dragonfly:
	$(eval GOOS=dragonfly)
	@:
freebsd:
	$(eval GOOS=freebsd)
	@:
illumos:
	$(eval GOOS=illumos)
	@:
ios:
	$(eval GOOS=ios)
	@:
linux:
	$(eval GOOS=linux)
	@:
js:
	$(eval GOOS=js)
	@:
nacl:
	$(eval GOOS=nacl)
	@:
netbsd:
	$(eval GOOS=netbsd)
	@:
openbsd:
	$(eval GOOS=openbsd)
	@:
plan9:
	$(eval GOOS=plan9)
	@:
solaris:
	$(eval GOOS=solaris)
	@:
windows:
	$(eval GOOS=windows)
	$(eval OUT=${OUT}.exe)
	@:
# Architectures
ppc64:
	$(eval GOARCH=ppc64)
	@:
ppc64le:
	$(eval GOARCH=ppc64le)
	@:
mips:
	$(eval GOARCH=mips)
	@:
mipsle:
	$(eval GOARCH=mipsle)
	@:
mips64:
	$(eval GOARCH=mips64)
	@:
mips64le:
	$(eval GOARCH=mips64le)
	@:
386:
	$(eval GOARCH=386)
	@:
amd64:
	$(eval GOARCH=amd64)
	@:
amd64p32:
	$(eval GOARCH=amd64p32)
	@:
arm:
	$(eval GOARCH=arm)
	@:
7:
	$(eval GOARM=7)
	@:
6:
	$(eval GOARM=6)
	@:
5:
	$(eval GOARM=5)
	@:
arm64:
	$(eval GOARCH=arm64)
	@:
s390x:
	$(eval GOARCH=s390x)
	@:
wasm:
	$(eval GOARCH=wasm)
	@:
riscv64:
	$(eval GOARCH=riscv64)
	@: