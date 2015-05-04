PREFIX=/usr
GOBIN=${DESTDIR}${PREFIX}/bin
NAME=tbpdump
PACKAGE_TYPE=rpm
PACKAGE_BUILD_DIR=pkg
PACKAGE_DIR=dist

build:

	go get
	go build

install: build

	#install binary
	GOBIN=${GOBIN} go install -v

package:

	#
	# export PACKAGE_TYPE to vary package type (e.g. deb, tar, rpm)
	#

	@if [ -z "$(shell which fpm 2>/dev/null)" ]; then \
		echo "error:\nPackaging requires effing package manager (fpm) to run.\nsee https://github.com/jordansissel/fpm\n"; \
		exit 1; \
	fi

	#run make install against the packaging dir
	$(MAKE) install DESTDIR=${PACKAGE_BUILD_DIR}

	#clean
	mkdir -p dist && rm -f dist/*.${PACKAGE_TYPE}

	#build package
	fpm --rpm-os linux \
		-s dir \
		-p dist \
		-t ${PACKAGE_TYPE} \
		-n ${NAME} \
		--vendor warmans \
		--url https://github.com/warmans/tbpdump \
		--description "decode thrift binary protocol messages from tcpdump input" \
		-v $(shell ${PACKAGE_BUILD_DIR}${GOBIN}/${NAME} -V) \
		-C ${PACKAGE_BUILD_DIR} .
