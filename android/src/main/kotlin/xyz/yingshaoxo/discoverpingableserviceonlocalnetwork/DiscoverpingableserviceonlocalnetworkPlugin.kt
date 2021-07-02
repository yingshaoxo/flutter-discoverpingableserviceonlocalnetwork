package xyz.yingshaoxo.discoverpingableserviceonlocalnetwork

import GoFind.GoFind
import android.app.Activity
import android.content.Context
import android.content.Context.WIFI_SERVICE
import android.net.wifi.WifiManager
import android.os.Build
import androidx.annotation.NonNull
import io.flutter.embedding.engine.plugins.FlutterPlugin
import io.flutter.plugin.common.MethodCall
import io.flutter.plugin.common.MethodChannel
import io.flutter.plugin.common.MethodChannel.MethodCallHandler
import java.math.BigInteger
import java.net.InetAddress
import java.net.UnknownHostException
import java.nio.ByteOrder
import io.flutter.embedding.engine.plugins.activity.ActivityAware
import io.flutter.embedding.engine.plugins.activity.ActivityPluginBinding


/** DiscoverpingableserviceonlocalnetworkPlugin */
class DiscoverpingableserviceonlocalnetworkPlugin : FlutterPlugin, MethodCallHandler,
    ActivityAware {
    /// The MethodChannel that will the communication between Flutter and native Android
    ///
    /// This local reference serves to register the plugin with the Flutter Engine and unregister it
    /// when the Flutter Engine is detached from the Activity
    private lateinit var context: Context
    private lateinit var activity: Activity

    private lateinit var channel: MethodChannel

    override fun onDetachedFromActivity() {
        TODO("Not yet implemented")
    }

    override fun onReattachedToActivityForConfigChanges(binding: ActivityPluginBinding) {
        TODO("Not yet implemented")
    }

    override fun onAttachedToActivity(binding: ActivityPluginBinding) {
        activity = binding.activity;
    }

    override fun onDetachedFromActivityForConfigChanges() {
        TODO("Not yet implemented")
    }

    override fun onAttachedToEngine(@NonNull flutterPluginBinding: FlutterPlugin.FlutterPluginBinding) {
        channel = MethodChannel(
            flutterPluginBinding.binaryMessenger,
            "discoverpingableserviceonlocalnetwork"
        )
        channel.setMethodCallHandler(this)

        context = flutterPluginBinding.applicationContext
    }

    override fun onDetachedFromEngine(@NonNull binding: FlutterPlugin.FlutterPluginBinding) {
        channel.setMethodCallHandler(null)
    }

    override fun onMethodCall(call: MethodCall, result: MethodChannel.Result) {
        if (call.method == "getPlatformVersion") {
            result.success("Android ${Build.VERSION.RELEASE}")
        } else if (call.method == "find_services_in_a_host") {
            var host: String? = call.argument<String>("host");
            var startPort: Int? = call.argument<Int>("startPort");
            var endPort: Int? = call.argument<Int>("endPort");
            if (host!=null && startPort!=null && endPort!=null) {
                var found: String = GoFind.scanPorts(host, startPort.toLong(), endPort.toLong())
                result.success(found)
                return ;
            }
            result.success("")
        } else if (call.method == "find_all_services") {
            var network: String? = call.argument<String>("network");
            var startPort: Int? = call.argument<Int>("startPort");
            var endPort: Int? = call.argument<Int>("endPort");
            if (network!=null && startPort!=null && endPort!=null) {
                var found: String = GoFind.scanAllHosts(network, startPort.toLong(), endPort.toLong())
                result.success(found)
                return ;
            }
            result.success("")
        } else if (call.method == "get_wifi_address") {
            var wifi_ip_address = getWifiIpAddress(context)
            if (wifi_ip_address != null) {
                result.success(wifi_ip_address)
            } else {
                result.success("")
            }
        } else {
            result.notImplemented()
        }
    }
}

fun getWifiIpAddress(context: Context): String? {
    val wifiManager = context.getSystemService(WIFI_SERVICE) as WifiManager
    var ipAddress = wifiManager.connectionInfo.ipAddress

    if (ByteOrder.nativeOrder().equals(ByteOrder.LITTLE_ENDIAN)) {
        ipAddress = Integer.reverseBytes(ipAddress)
    }
    val ipByteArray: ByteArray = BigInteger.valueOf(ipAddress.toLong()).toByteArray()

    return try {
        InetAddress.getByAddress(ipByteArray).hostAddress
    } catch (ex: UnknownHostException) {
        null
    }
}
