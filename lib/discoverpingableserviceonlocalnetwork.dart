import 'dart:async';

import 'package:flutter/services.dart';

class Discoverpingableserviceonlocalnetwork {
  static const MethodChannel _channel =
      const MethodChannel('discoverpingableserviceonlocalnetwork');

  static Future<String?> get platformVersion async {
    final String? version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }

  static Future<String?> findServicesInAHost(
      String host, int startPort, int endPort) async {
    final String? json = await _channel.invokeMethod(
        'find_services_in_a_host', <String, dynamic>{
      'host': host,
      'startPort': startPort,
      'endPort': endPort
    });

    if (json == "") {
      return null;
    }

    return json;
  }

// network = "192.168.1.1/24"
  static Future<String?> findServicesInANetwork(
      String network, int startPort, int endPort) async {
    final String? json = await _channel.invokeMethod(
        'find_all_services', <String, dynamic>{
      'network': network,
      'startPort': startPort,
      'endPort': endPort
    });

    if (json == "") {
      return null;
    }

    return json;
  }

  static Future<String?> getWIFIaddress() async {
    final String? json = await _channel.invokeMethod('get_wifi_address');

    if (json == "") {
      return null;
    }

    return json;
  }
}
