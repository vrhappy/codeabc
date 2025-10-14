openeuler上安装rust环境，由于用不了rustup，过程有些不同，可以使用conda建个虚拟环境，在其上操作。
1.conda create -c conda-forge -n rustenv rust
参考https://cn.linux-terminal.com/?p=4641
2.安装后conda activate rustenv激活虚拟环境，
执行cargo install --locked cargo-pgrx --version 0.15.0时报错“  = note: rust-lld: error: unable to find library -lssl
          rust-lld: error: unable to find library -lcrypto
          collect2: error: ld returned 1 exit status”
参考https://github.com/rust-cross/cargo-zigbuild/discussions/137，加上RUSTFLAGS='-L /usr/lib64'后安装成功，
RUSTFLAGS='-L /usr/lib64' cargo install --locked cargo-pgrx --version 0.15.0
至此，安装pgrx成功。
接下来，要编译pg_search。
