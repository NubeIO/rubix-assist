import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:frontend/model/git_plugins.dart';
import 'package:frontend/model/host.dart';
import 'package:frontend/model/ws.dart';
import 'package:frontend/services/api_plugins.dart';
import 'package:frontend/services/ws.dart';
import 'package:multiselect/multiselect.dart';
import 'edit_screen.dart';


class DetailWidget extends StatefulWidget {
  DetailWidget(this.host);
  final Host host;
  @override
  _DetailWidgetState createState() => _DetailWidgetState();
}

Future<void> printInteger(String message) async {
  Fluttertoast.showToast(
      msg: message.toString(),
      toastLength: Toast.LENGTH_LONG,
      gravity: ToastGravity.CENTER,
      timeInSecForIosWeb: 8,
      backgroundColor: Colors.black,
      textColor: Colors.white,
      fontSize: 16.0
  );
}

class _DetailWidgetState extends State<DetailWidget> {
  final ApiPlugins apiGit = ApiPlugins();
  late Future <List<GitPlugins>> gitPluginList;

  List<String> selected = [];
  List<String> pluginNames = [];

  bool _system = true;
  int? id;

  @override
  void initState() {
    Eventbus.on().listen((event) {
      WsPayload user = WsPayload.fromJson(jsonDecode(event));
      printInteger(user.topic);
    });
    id = widget.host.id!;
    gitPluginList = apiGit.getPlugins(id.toString());
    gitPluginList.then((value) => {
      value.forEach((element) {
        pluginNames.add(element.name);
      })
    });

    super.initState();
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('DEVICE: ${widget.host.name.toString()}'),
        actions: [
          IconButton(
              onPressed: () {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) => EditScreen(widget.host)));
              },
              icon: Icon(Icons.edit)),
        ],
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: EdgeInsets.all(20),
          child: Column(
            children: [
              Row(children: const [
                Expanded(child: Divider(thickness: 1.5)),
                Text("INSTALL PLUGINS",
                    style: TextStyle(fontSize: 20, color: Colors.grey)),
                Expanded(child: Divider(thickness: 1.5)),
              ]),
              Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  const Text(
                    '     Fetch Plugins (tick and un-tick)',
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                  Checkbox(
                    value: _system,
                    onChanged: (value) {
                      setState(() {
                        _system = !_system;
                      });
                    },
                  ),
                ],
              ),
              SizedBox(
                height: 20,
              ),
              DropDownMultiSelect(
                onChanged: (List<String> x) {
                  setState(()  {
                    selected = x;
                    print(selected);
                  });
                },
                options: pluginNames,
                selectedValues: selected,
                whenEmpty: 'Select Something',
              ),
              SizedBox(
                height: 20,
              ),
              // Image.network('${widget.host.image}')
            ],
          ),
        ),
      ),
    );
  }
}
