import 'dart:convert';

import 'package:frontend/model/host.dart';
import 'package:http/http.dart' as http;
import 'package:http/http.dart';


class ApiService {
  final String hostUrl = "http://0.0.0.0:8080/api/hosts";
  final String pluginUrl = "http://0.0.0.0:8080/api/plugins";

  Future<List<Host>> getHost() async {
    final res = await http.get(Uri.parse(hostUrl+"/"));
    if (res.statusCode == 200) {
      List jsRespone = json.decode(res.body);
      print(jsRespone);
      return jsRespone.map((e) => new Host.fromJson(e)).toList();
    } else {
      throw Exception('Failed to load host') ;
    }
  }

  Future<Host> getHostById(String id) async {
    final response = await http.get(Uri.parse('$hostUrl/$id'));

    if (response.statusCode == 200) {
      return Host.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load a case');
    }
  }

  Future<Host> createHost(Host host) async {
    Map data = {
      'name': host.name
    };
    final  response = await post(
      Uri.parse(hostUrl+"/"),
      headers: <String, String>{
        'Content-Type':'application/json; charset=UTF-8',
        'Accept' : 'application/json'
      },
      body: jsonEncode(data),
    );
    if (response.statusCode == 201) {
      return Host.fromJson(json.decode(response.body));
    } else {
      throw Exception(response.statusCode);
    }
  }

  Future<Host> updateHost(int id, Host host) async {
    Map data = {
      'name': host.name,
      'ip': host.ip,
      'port': host.port,
      'username': host.username,
      'password':host.password,
      'rubix_port':host.rubixPort,
      'rubix_username':host.rubixUsername,
      'rubix_password':host.rubixPassword,
    };
    print(data);
    final Response response = await patch(
      Uri.parse('$hostUrl/$id'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(data),
    );
    if (response.statusCode == 200) {
      return Host.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update a case');
    }
  }
  Future deleteHost(String id) async {
    final http.Response res = await delete(Uri.parse('$hostUrl/$id'),
      headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },);
    if (res.statusCode == 200) {
      return true;
    } else {
      throw "Failed to delete a case.";
    }
  }
}

