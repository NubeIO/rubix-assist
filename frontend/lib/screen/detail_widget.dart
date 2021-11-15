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



  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Details'),
        actions: [
          IconButton(onPressed: (){
            Navigator.push(context, MaterialPageRoute(builder: (context)=>EditScreen(widget.product)));
          }, icon: Icon(Icons.edit)),

        ],
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: EdgeInsets.all(20),
          child: Column(
            children: [
              Row(
                children: [
                  Text('Product Name: '),
                  SizedBox(width: 20,),
                  Text('${widget.product.name}')
                ],
              ),
              SizedBox(height: 20,),
              // Image.network('${widget.product.image}')
            ],
          ),
        ),
      ),
    );
  }

}
