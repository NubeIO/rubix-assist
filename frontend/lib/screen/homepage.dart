import 'dart:async';

import 'package:flutter/material.dart';
import 'package:frontend/model/product.dart';
import 'package:frontend/services/api.dart';

import 'add_product.dart';
import 'detail_widget.dart';

class HomePage extends StatefulWidget {
  HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final ApiService apiService = ApiService();
  late Future<List<Product>> productList;
  late Future<Product> _product;
  String? id = '';

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    setState(() {
      productList = apiService.getProduct();
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
            child: FutureBuilder<List<Product>>(
                future: productList,
                builder: (context, snapshot) {
                  if (snapshot.hasData) {
                    List<Product>? product = snapshot.data;
                    return ListView.builder(
                      itemCount: product!.length,
                      itemBuilder: (BuildContext context, int index) {
                        return Card(
                          child: InkWell(
                            child: ListTile(
                              leading: Icon(Icons.auto_awesome_mosaic_outlined),
                              title: Text('NAME ${product[index].name}'),
                              subtitle: Text(
                                '${product[index].name}',
                                style: TextStyle(color: Colors.grey),
                              ),
                              trailing: Row(
                                mainAxisSize: MainAxisSize.min,
                                children: [
                                  IconButton(
                                      onPressed: () {}, icon: Icon(Icons.edit)),
                                  IconButton(
                                      onPressed: () {
                                        apiService.deleteProduct('${product[index].id}');
                                           setState(() {
                                             productList = apiService.getProduct();
                                           });
                                      },
                                      icon: Icon(Icons.delete)),
                                  IconButton(
                                      onPressed: () {},
                                      icon: Icon(Icons.add_box)),
                                ],
                              ),
                              // trailing: IconButton(
                              //   icon: Icon(Icons.delete),
                              //   onPressed: (){
                              //    apiService.deleteProduct('${product[index].id}');
                              //    setState(() {
                              //      productList = apiService.getProduct();
                              //    });
                              //   },
                              // ),
                            ),
                            onTap: () {
                              Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) =>
                                        DetailWidget(product[index]),
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
      MaterialPageRoute(builder: (context) => AddProduct()),
    );
  }

  Future<void> refresh() {
    setState(() {
      productList = apiService.getProduct();
    });
    return Future.delayed(Duration(seconds: 2));
  }
}
