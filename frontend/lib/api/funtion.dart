import 'dart:convert';

import 'package:http/http.dart' as http;


class Createddata {

  Future datacreated(nametext, emailtext,) async {
    print(nametext);
    print(emailtext);
    final responce =
    await  http.post(
      Uri.parse('http://0.0.0.0:8080/api/hosts'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Access-Control-Allow-Origin':'true'
      },
      body: jsonEncode(<String, String>{
        "name":nametext,
        "ip": emailtext,
      }),
    );


    // print(1234);
    // final responce =
    //     await http.post(Uri.parse('http://0.0.0.0:8080/api/hosts'),
    //         body: jsonEncode({
    //           "name":nametext,
    //           "ip": emailtext,
    //         })
    //     });
    print(232334);
    print(responce.statusCode);
    if (responce.statusCode == 201) {
     
      print('Data Created Successfully');
    } else {
      print('erro');
    }
  }

}
