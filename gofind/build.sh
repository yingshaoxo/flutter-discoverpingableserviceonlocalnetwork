export ANDROID_HOME=/home/yingshaoxo/Android/Sdk
export ANDROID_NDK_HOME=home/yingshaoxo/Android/Sdk/ndk/25.1.8937393

go get golang.org/x/mobile/cmd/gomobile@latest
go install golang.org/x/mobile/cmd/gomobile@latest
export PATH="$PATH:$HOME/go/bin"

gomobile init
gomobile bind -target=android -androidapi=19

mv -f *.aar ../android/libs 
mv -f *.jar ../android/libs 