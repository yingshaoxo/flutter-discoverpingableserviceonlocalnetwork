import 'dart:convert';
import 'dart:developer';

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
  String ipWeFound = "";

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
          child: Text(_platformVersion),
        ),
        floatingActionButton: FloatingActionButton(
          child: Text("get services"),
          onPressed: () async {
            _platformVersion = "tapped";
            setState(() {});
            if (ipWeFound != "") {
              String? result = await Discoverpingableserviceonlocalnetwork
                  .findServicesInAHost(ipWeFound, 5000, 5100);
              _platformVersion =
                  DateTime.now().toString() + (result ?? "Can't find anything");
              setState(() {});
            } else {
              String? ip =
                  await Discoverpingableserviceonlocalnetwork.getWIFIaddress();
              _platformVersion = "searching at " + (ip ?? "") + "...";
              setState(() {});
              if (ip != null) {
                String? services = await Discoverpingableserviceonlocalnetwork
                    .findServicesInANetwork(ip + "/24", 5000, 5010);
                if (services != null) {
                  if (services != "") {
                    List<dynamic> data = jsonDecode(services);
                    if (data.length > 0) {
                      _platformVersion = "we found: " + data[0];
                      ipWeFound = data[0].split(":")[0];
                      setState(() {});
                    } else {}
                  }
                }
              }
            }
          },
        ),
      ),
    );
  }
}
