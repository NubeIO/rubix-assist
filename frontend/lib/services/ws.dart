
import 'package:event_bus/event_bus.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

EventBus Eventbus = EventBus();
Future<void> enableWS() async {

  final channel = WebSocketChannel.connect(
    Uri.parse('ws://0.0.0.0:8080/ws'),
  );

  /// Listen for all incoming data
  channel.stream.listen(
        (data) {
      print(data.toString());
      Eventbus.fire(data.toString());
      print(22222);
    },
    onError: (error) => print(error),
  );

  channel.sink.add('Hello! aidan');
}