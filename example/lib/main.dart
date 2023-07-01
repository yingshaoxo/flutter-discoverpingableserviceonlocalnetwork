import 'package:flutter/material.dart';
import 'dart:async';

import 'package:flutter/services.dart';
import 'package:discoverpingableserviceonlocalnetwork/discoverpingableserviceonlocalnetwork.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  String _platformVersion = 'Unknown';
  String wifi_address = "";

  @override
  void initState() {
    super.initState();
    initPlatformState();
  }

  // Platform messages are asynchronous, so we initialize in an async method.
  Future<void> initPlatformState() async {
    String platformVersion;
    // Platform messages may fail, so we use a try/catch PlatformException.
    // We also handle the message potentially returning null.
    try {
      platformVersion =
          await Discoverpingableserviceonlocalnetwork.platformVersion ??
              'Unknown platform version';
    } on PlatformException {
      platformVersion = 'Failed to get platform version.';
    }

    try {
      // wifi_address =
      //     await Discoverpingableserviceonlocalnetwork.getWIFIaddress() ?? "";
      wifi_address = "192.168.49.1/24";
    } on PlatformException {
      wifi_address = 'Failed to get wifi address.';
    }

    // If the widget was removed from the tree while the asynchronous platform
    // message was in flight, we want to discard the reply rather than calling
    // setState to update our non-existent appearance.
    if (!mounted) return;

    setState(() {
      _platformVersion = platformVersion;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Plugin example app'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(_platformVersion),
            ],
          ),
        ),
        floatingActionButton: FloatingActionButton(
          child: Text("Find"),
          onPressed: () async {
            _platformVersion = "searching...";
            setState(() {});
            // List<String>? hosts = await Discoverpingableserviceonlocalnetwork
            //     .findServicesInAHost(wifi_address, 0, 49151, 500);
            List<String>? hosts = await Discoverpingableserviceonlocalnetwork
                .findServicesInANetwork(wifi_address, 0, 5020, 3000);
            if (hosts != null) {
              _platformVersion =
                  "This is what I found: \n\n" + hosts.toString();
              setState(() {});
            }
          },
        ),
      ),
    );
  }
}
