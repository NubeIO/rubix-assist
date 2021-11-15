// ignore_for_file: prefer_const_constructors, avoid_unnecessary_containers, sized_box_for_whitespace, avoid_print

import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

import 'api/funtion.dart';

class Homepage extends StatefulWidget {
  const Homepage({Key? key}) : super(key: key);

  @override
  _HomepageState createState() => _HomepageState();
}

class _HomepageState extends State<Homepage> {
  Createddata obj = Createddata();
  TextEditingController nameController = TextEditingController();
  TextEditingController emailController = TextEditingController();
  @override
  void initState() {
    super.initState();
    getdata();
    // // datadelet("1");
    // obj.datacreated(nameController.text, emailController.text);
  }

  List data = [];
  late String id;
  Future getdata() async {
    final responce =
        await http.get(Uri.parse('http://0.0.0.0:8080/api/hosts/'));
    if (responce.statusCode == 200) {
      setState(() {
        data = jsonDecode(responce.body);
      });
      print('Add data$data');
    } else {
      print('erro');
    }
  }

  Future datadelet(id) async {
    final responce = await http
        .delete(Uri.parse('http://0.0.0.0:8080/api/hosts/$id'));
    if (responce.statusCode == 200) {
      print('DELETE COMPLETE');
    } else {
      print('nOT dELET');
    }
  }

  Future update(id) async {
    final responce = await http
        .patch(Uri.parse('http://0.0.0.0:8080/api/hosts/$id'),
            body: jsonEncode({
              "name": nameController.text,
              "username": emailController.text,
            }),
            headers: {
          'Content-type': 'application/json; charset=UTF-8',
        });
    print(responce.statusCode);
    if (responce.statusCode == 200) {
      print('Data Update Sucassfully');
      nameController.clear();
      emailController.clear();
    } else {
      print('erro');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(20),
          child: Container(
            child: Column(
              children: [
                TextField(
                  controller: nameController,
                  decoration: InputDecoration(
                    hintText: 'Name',
                  ),
                ),
                TextField(
                  controller: emailController,
                  decoration: InputDecoration(
                    hintText: 'Number',
                  ),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    ElevatedButton(
                        onPressed: () {
                          setState(() {
                            obj.datacreated(
                              nameController.text,
                              emailController.text,
                            );
                            getdata();
                          });
                        },
                        child: Text('Submit')),
                    ElevatedButton(
                        onPressed: () {
                          setState(() {
                            update(id);
                          });
                        },
                        child: Text('Update')),
                  ],
                ),
                Expanded(
                  child: ListView.builder(
                      itemCount: data.length,
                      itemBuilder: (context, index) {
                        return Container(
                          child: ListTile(
                            title: Text(data[index]['name']),
                            subtitle: Text(data[index]['ip']),
                            trailing: Container(
                              width: 100,
                              child: Row(
                                children: [
                                  IconButton(
                                      onPressed: () {
                                        nameController.text =
                                            data[index]['name'];
                                        emailController.text =
                                            data[index]['ip'];
                                      },
                                      icon: Icon(Icons.edit)),
                                  IconButton(
                                      onPressed: () {
                                        setState(() {
                                          datadelet(data[index]['id']);
                                        });
                                      },
                                      icon: Icon(Icons.delete))
                                ],
                              ),
                            ),
                          ),
                        );
                      }),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
