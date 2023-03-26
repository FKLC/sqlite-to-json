oss=(darwin freebsd linux windows)
archs=(amd64 arm64 386 arm ppc64le riscv64)

for os in ${oss[@]}
do
    for arch in ${archs[@]}
    do
        echo Starting build for ${os} ${arch}
        GOOS=${os} GOARCH=${arch} go build -ldflags "-w" -o build/stj_${os}_${arch}
        echo Ending build for ${os} ${arch}
    done
done

# Overwrite any possible go build failures
exit 0