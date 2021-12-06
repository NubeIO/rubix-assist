import 'package:flutter/material.dart';
import 'package:frontend/services/ws.dart';
import 'screen/homepage.dart';



void main() async{
  enableWS();
  runApp(MyApp());

}


class MyApp extends StatelessWidget {

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'API',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: HomePage(),
    );
  }
}
