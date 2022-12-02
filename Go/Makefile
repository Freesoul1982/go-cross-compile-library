all: build_x86_linux build_arm_linux build_arm_android

build_x86_linux:
	mkdir -p x86/linux
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -o x86/linux/libGoSDK.so goSDK.go

build_arm_linux:
	mkdir -p arm/linux
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -buildmode=c-shared -o arm/linux/libGoSDK.so goSDK.go

build_arm_android:
	mkdir -p arm/android
	CC=/home/ndk/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang CGO_ENABLED=1 GOOS=android GOARCH=arm64 go build -buildmode=c-shared -o arm/android/libGoSDK.so goSDK.go