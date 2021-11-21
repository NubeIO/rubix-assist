import 'dart:async';
import 'package:flutter/material.dart';
import 'package:frontend/model/host.dart';
import 'package:frontend/services/api.dart';
import 'add_host.dart';
import 'detail_widget.dart';


class HomePage extends StatefulWidget {
  HomePage({Key? key}) : super(key: key);
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

  final ApiService apiService = ApiService();
  late Future<List<Host>> hostList;


  String? id = '';

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    setState(() {
      hostList = apiService.getHost();
    });
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          navigateToAddScreen(context);
        },
        child: Icon(Icons.add),
      ),
      appBar: AppBar(
        title: Text('Nube Devices'),
      ),
      body: Container(
        child: Center(
          child: RefreshIndicator(
            onRefresh: refresh,
            child: FutureBuilder<List<Host>>(
                future: hostList,
                builder: (context, snapshot) {
                  if (snapshot.hasData) {
                    List<Host>? host = snapshot.data;
                    return ListView.builder(
                      itemCount: host!.length,
                      itemBuilder: (BuildContext context, int index) {
                        return Card(
                          child: InkWell(
                            child: ListTile(
                              leading: Icon(Icons.auto_awesome_mosaic_outlined),
                              title: Text('DEVICE: ${host[index].name}'),
                              subtitle: Text(
                                'IP: ${host[index].ip}',
                                style: TextStyle(color: Colors.grey),
                              ),
                              trailing: Row(
                                mainAxisSize: MainAxisSize.min,
                                children: [
                                  IconButton(
                                      tooltip: 'Install Flow-Framework Plugins',
                                      onPressed: () {}, icon: Icon(Icons.downloading)),
                                  IconButton(
                                      tooltip: 'Update Bios',
                                      onPressed: () {},
                                      icon: Icon(Icons.restart_alt)),
                                  IconButton(
                                      tooltip: 'Delete Host',
                                      onPressed: () {
                                        apiService.deleteHost('${host[index].id}');
                                        setState(() {
                                          hostList = apiService.getHost();
                                        });
                                      },
                                      icon: Icon(Icons.delete)),
                                ],
                              ),
                            ),
                            onTap: () {
                              Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) =>
                                        DetailWidget(host[index]),
                                  ));
                            },
                          ),
                        );
                      },
                    );
                  } else if (snapshot.hasError) {
                    return Text('${snapshot.error}');
                  }
                  return CircularProgressIndicator();
                }),
          ),
        ),
      ),
    );
  }

  navigateToAddScreen(BuildContext context) async {
    final result = await Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => AddHost()),
    );
  }

  Future<void> refresh() {
    setState(() {
      hostList = apiService.getHost();
    });
    return Future.delayed(Duration(seconds: 2));
  }
}
