# discoverpingableserviceonlocalnetwork

Help you use flutter to do port scan, so that you could find services on local area network.

https://pub.dev/packages/discoverpingableserviceonlocalnetwork

## Usage
```bash
  discoverpingableserviceonlocalnetwork:
    git:
      url: https://github.com/yingshaoxo/flutter-discoverpingableserviceonlocalnetwork.git
      ref: master
```

```dart
import 'package:discoverpingableserviceonlocalnetwork/discoverpingableserviceonlocalnetwork.dart';


await Discoverpingableserviceonlocalnetwork.getWIFIaddress();

await Discoverpingableserviceonlocalnetwork
    .findServicesInAHost("192.168.49.1", 0, 49151, 500);

await Discoverpingableserviceonlocalnetwork
    .findServicesInANetwork("192.168.49.1/24", 5000, 8000, 500);
```


<!-- ## Develop logs

### aar error
Direct local .aar file dependencies are not supported when building an AAR. The resulting AAR would be broken because the classes and Android resources from any local .aar file dependencies would not be packaged in the resulting AAR. Previous versions of the Android Gradle Plugin produce broken AARs in this case too (despite not throwing this error). The following direct local .aar file dependencies of the :discoverpingableserviceonlocalnetwork project caused this error: /home/yingshaoxo/work/discoverpingableserviceonlocalnetwork/android/libs/GoFind.aar

So we couldn't use golang to make a flutter plugin anymore?

Or couldn't make a aar module with another bunch of aar modules?

Solution: https://stackoverflow.com/a/57075972/8667243 

### how to use golang to make a flutter plugin
* https://youtu.be/668hBBd9wTU
* https://yingshaoxo.blogspot.com/2021/07/how-to-use-golang-to-make-flutter-plugin.html -->