MAC=GOOS=darwin GOARCH=amd64
LIN64=GOOS=linux GOARCH=amd64
LIN32=GOOS=linux GOARCH=386

OUTFILE_NAME=alter-barman
OUTFILE_NAME_MAC=alter-barman-mac
OUTFILE_NAME_LIN64=alter-barman-lin64
OUTFILE_NAME_LIN32=alter-barman-lin32
OUTDIR=./bin

CC=go build

all: run

run: go run .

build-all: mac lin64 lin32

build:
	mkdir -p ${OUTDIR}
	GOARCH=${GOARCH} GOOS=${GOOS} ${CC} -o ${OUTDIR}/${OUTFILE_NAME} ./cmd/main.go

mac:
	$(MAC) $(CC) -o ${OUTDIR}/${OUTFILE_NAME_MAC} ./cmd/main.go

lin64:
	$(LIN64) $(CC) -o ${OUTDIR}/${OUTFILE_NAME_LIN64} ./cmd/main.go

lin32:
	$(LIN32) $(CC) -o ${OUTDIR}/${OUTFILE_NAME_LIN32} ./cmd/main.go

.PHONY: clean
clean:
	rm -rf ${OUTDIR}

rebuild: clean build-all