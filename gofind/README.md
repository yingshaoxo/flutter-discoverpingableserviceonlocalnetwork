```
export ANDROID_HOME=/Users/yingshaoxo/Library/Android/sdk
export ANDROID_NDK_HOME=/Users/yingshaoxo/Library/Android/sdk/ndk/25.1.8937393

go get golang.org/x/mobile/cmd/gomobile@latest
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
gomobile bind -target=android -androidapi=19

mv -f *.aar ../android/libs 
mv -f *.jar ../android/libs 
```
