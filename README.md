# discoverpingableserviceonlocalnetwork

Help you use flutter to do port scan, so that you could find services on local area network.

## Usage
```bash
  discoverpingableserviceonlocalnetwork:
    git:
      url: git://github.com/yingshaoxo/flutter-discoverpingableserviceonlocalnetwork
      ref: master
```

## error
```
Direct local .aar file dependencies are not supported when building an AAR. The resulting AAR would be broken because the classes and Android resources from any local .aar file dependencies would not be packaged in the resulting AAR. Previous versions of the Android Gradle Plugin produce broken AARs in this case too (despite not throwing this error). The following direct local .aar file dependencies of the :discoverpingableserviceonlocalnetwork project caused this error: /home/yingshaoxo/work/discoverpingableserviceonlocalnetwork/android/libs/GoFind.aar
```

So we couldn't use golang to make a flutter plugin anymore?

Or couldn't make a aar module with another bunch of aar modules?

#golang #flutter #androidStudio

**I feel sorry for the update of android studio, it kills the possibility of using golang as a code base for android and ios.**
