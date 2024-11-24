tempFile = ".temp"


build:
	@touch ${tempFile}
	@/usr/bin/time -o ${tempFile} go build -o bin/bin -v
	@echo "Build complete: ./bin/bin\nTime spent:"
	@cat ${tempFile}
	@rm ${tempFile}

