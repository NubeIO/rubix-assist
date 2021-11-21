class Host{
 final int? id;
 final String? name;
 final String? ip;
 final int? port;
 final String? username;
 final String? password;
 final int? rubixPort;
 final String? rubixUsername;
 final String? rubixPassword;

  Host({ this.id,this.name,this.ip,this.port,this.username,this.password,this.rubixPort,this.rubixUsername,this.rubixPassword});
  factory Host.fromJson(Map<String,dynamic> json){
    return Host(
      id: json['id'] ,
      name: json['name'] ,
      ip: json['ip'] ,
      port: json['port'] ,
      username: json['username'] ,
      password: json['password'] ,
      rubixPort: json['rubix_port'] ,
      rubixUsername: json['rubix_username'] ,
      rubixPassword: json['rubix_password'] ,
    );
  }
  @override
  String toString() {
    // TODO: implement toString
    return 'Host{id: $id, name: $name}';
  }
}
