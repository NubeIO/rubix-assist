
class WsPayload {
  WsPayload({
    required this.msg,
    required this.topic,
  });

  String msg;
  String topic;

  factory WsPayload.fromJson(Map<String, dynamic> json) => WsPayload(
    msg: json["msg"],
    topic: json["topic"],
  );

  Map<String, dynamic> toJson() => {
    "msg": msg,
    "topic": topic,
  };
}
