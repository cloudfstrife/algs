# 主版本
VERSION ?= $(shell git describe --tags --always --dirty)

# git 提交 Hash
COMMIT_HASH ?= $(shell git show -s --format=%H)

# build 时间
BUILD_TIME ?= $(shell date +%Y%m%d%H%M%S)

# go文件列表
GOFILES := $(shell find . ! -path "./vendor/*" -name "*.go")

BUILD_ENV := 

# 单元测试附加选项
TEST_OPTS := -v

# 基准测试附加选项
BENCHMARK_OPTS := -cpu 1,2,3,4,5,6,7,8 -benchmem 

# sonar 相关报告输出路径（包括：单元测试报告输出，单元测试覆盖率报告，golint 报告，golangci-lint 报告）
REPORT_FOLDER := sonar
TEST_REPORT := ${REPORT_FOLDER}/test.report 
COVER_REPORT := ${REPORT_FOLDER}/cover.report
GOLANGCI_LINT_REPORT := ${REPORT_FOLDER}/golangci-lint.xml 
GOLINT_REPORT := ${REPORT_FOLDER}/golint.report 

.PHONY: format test benchmark sonar clean

# 单元测试
test: 
	${BUILD_ENV} go test ${TEST_OPTS} ./...

# 格式化
format:
	@for f in ${GOFILES} ; do 											\
		gofmt -w $${f};													\
	done																\

# 基准测试
benchmark:
	go test -bench . -run ^$$ ${BENCHMARK_OPTS}  ./...

# sonar
sonar: 
	mkdir -p ${REPORT_FOLDER}
	go test -json ./... > ${TEST_REPORT}
	go test -coverprofile=${COVER_REPORT} ./... 
	golangci-lint run --out-format checkstyle  ./... > ${GOLANGCI_LINT_REPORT}
	golint ./... > ${GOLINT_REPORT}
	# sonar-scanner

# 清理
clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${COVER_REPORT}
	-rm -f ${GOLANGCI_LINT_REPORT}
	-rm -f ${GOLINT_REPORT}
	-go clean 
	-go clean -cache