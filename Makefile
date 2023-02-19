GOARCH=arm64
GOOS=darwin
OUTFILE_NAME=alter-barman
OUTDIR=./bin

CC=go build

build:
	mkdir -p ${OUTDIR}
	GOARCH=${GOARCH} GOOS=${GOOS} ${CC} -o ${OUTDIR}/${OUTFILE_NAME} main.go

.PHONY: clean
clean:
	rm -rf ${OUTDIR}
