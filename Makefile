# 编译目标文件的目录
BUILD_DIR := /Users/matin/projects/test/Docker/MinIO/service/api/bin

# 源代码文件
SOURCE_FILES := $(wildcard *.go)

# 生成的可执行文件名
EXECUTABLE := minio-api-bin

# 默认目标
all: build

# 编译可执行文件
build: $(BUILD_DIR)/$(EXECUTABLE)

$(BUILD_DIR)/$(EXECUTABLE): $(SOURCE_FILES)
	GOOS=linux GOARCH=amd64 go build -o $@
