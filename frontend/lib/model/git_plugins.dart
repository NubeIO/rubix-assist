
import 'dart:convert';

List<GitPlugins> gitPluginsFromJson(String str) => List<GitPlugins>.from(json.decode(str).map((x) => GitPlugins.fromJson(x)));

class GitPlugins {
  GitPlugins({
    required this.url,
    required this.name,
  });

  String url;
  String name;



  factory GitPlugins.fromJson(Map<String, dynamic> json) => GitPlugins(
    url: json["url"],
    name: json["name"],
  );
  // MedicalRecordsModel.fromJson(jsonDecode(response.body) as Map<String, dynamic>);


  Map<String, dynamic> toJson() => {
    "url": url,
    "name": name,
  };
}