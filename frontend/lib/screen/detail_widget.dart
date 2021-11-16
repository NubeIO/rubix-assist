import 'package:flutter/material.dart';
import 'package:frontend/model/product.dart';
import 'package:frontend/services/api.dart';

import 'edit_screen.dart';

class DetailWidget extends StatefulWidget {
  DetailWidget(this.product);

  final Product product;

  @override
  _DetailWidgetState createState() => _DetailWidgetState();
}

class _DetailWidgetState extends State<DetailWidget> {
  final ApiService apiService = ApiService();

  List<String> _locations = ['1', '2', '2', 'D']; // Option 2
  String? _selectedLocation; // Option 2
  late Future <Product> productList;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('DEVICE: ${widget.product.name.toString()}'),
        actions: [
          IconButton(
              onPressed: () {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) => EditScreen(widget.product)));
              },
              icon: Icon(Icons.edit)),
        ],
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: EdgeInsets.all(20),
          child: Column(
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: <Widget>[
                  ElevatedButton(
                    child: Text("LogIn"),
                    onPressed: () { print(111);},
                  ),
                  SizedBox(width: 5),
                  DropdownButton(
                    hint: Text('Please choose a location'), // Not necessary for Option 1
                    value: _selectedLocation,
                    onChanged: (newValue) {
                      setState(() {
                        _selectedLocation = newValue.toString();
                        print(_selectedLocation);
                        // getProductById("1");
                        try {
                          productList = apiService.getProductById(_selectedLocation.toString());
                          productList.then((value) => {
                            print(value.name)
                          });
                        } catch (e) {
                          print(e);
                        }

                      });
                    },
                    items: _locations.map((location) {
                      return DropdownMenuItem(
                        child: new Text(location),
                        value: location,
                      );
                    }).toList(),
                  ),
                ],
              ),
              SizedBox(
                height: 20,
              ),
              Row(
                children: [
                  Text('Product Name2: '),
                  SizedBox(
                    width: 20,
                  ),
                  Text('${widget.product.name}')
                ],
              ),
              SizedBox(
                height: 20,
              ),

              // Image.network('${widget.product.image}')
            ],
          ),
        ),
      ),
    );
  }
}
