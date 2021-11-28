import 'dart:convert';
import 'package:frontend/model/git_plugins.dart';
import 'package:http/http.dart';

class ApiPlugins {

  //Install plugins
  final String pluginUrl = "http://0.0.0.0:8080/api/plugins";
  final String gitUrl = "http://0.0.0.0:8080/api/git";

  Future updatePlugins(String id, List<String> plugins) async {
    Map data = {
      "from_path": "/home/aidan/Downloads",
      "to_path": "/home/pi/.tmp_uploads",
      "unzip_path": "/data/flow-framework/data/plugins",
      "unzip": true,
      "clear_dir": true,
      "zips": plugins
    };
    print(data);
    final Response response = await post(
      Uri.parse('$pluginUrl/full_install/$id'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(data),
    );
    print(response.statusCode);
    print('$pluginUrl/full_install/$id');
    print(response.body);
    if (response.statusCode == 200) {
      return response.body;
    } else {
      throw Exception('Failed to update a case');
    }
  }

  Future<List<GitPlugins>> getPlugins(String id) async {
    final response = await get(Uri.parse('$gitUrl/$id'));
    if (response.statusCode == 200) {
      print(json.decode(response.body));
      dynamic result = jsonDecode(response.body);
      return (result as List).map((e) => GitPlugins.fromJson(e)).toList();
    } else {
      throw Exception('Failed to load a case');
    }
  }

}

