import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:discoverpingableserviceonlocalnetwork/discoverpingableserviceonlocalnetwork.dart';

void main() {
  const MethodChannel channel = MethodChannel('discoverpingableserviceonlocalnetwork');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await Discoverpingableserviceonlocalnetwork.platformVersion, '42');
  });
}
