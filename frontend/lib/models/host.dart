// class Album {
//   final int id;
//   final String ip;
//
//   Album({
//
//     required this.id,
//     required this.ip,
//   });
//
//   factory Album.fromJson(Map<String, dynamic> json) {
//     return Album(
//       id: json['id'],
//       ip: json['ip'],
//     );
//   }
// }

class Hosts {
  Hosts({
    required this.id,
    required this.ip,
    required this.port,
    required this.username,
    required this.password,
  });

  int id;
  String ip;
  int port;
  String username;
  String password;

  factory Hosts.fromJson(Map<String, dynamic> json) => Hosts(
    id: json["id"],
    ip: json["ip"],
    port: json["port"],
    username: json["username"],
    password: json["password"],
  );

  Map<String, dynamic> toJson() => {
    "id": id,
    "ip": ip,
    "port": port,
    "username": username,
    "password": password,
  };
}
