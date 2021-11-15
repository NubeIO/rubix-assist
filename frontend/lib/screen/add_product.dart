import 'package:flutter/material.dart';
import 'package:frontend/model/product.dart';
import 'package:frontend/services/api.dart';

class AddProduct extends StatefulWidget {
  const AddProduct({Key? key}) : super(key: key);

  @override
  _AddProductState createState() => _AddProductState();
}

class _AddProductState extends State<AddProduct> {
  final ApiService apiService = ApiService();
  final _formKey = GlobalKey<FormState>();
  final _productNameController = TextEditingController();


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Add New Product'),
      ),
      body: SingleChildScrollView(
        child: Form(
          key: _formKey,
          child: Container(
            padding: EdgeInsets.all(25),
            child: Column(
              children: [
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _productNameController,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Please enter product name';
                    }
                    return null;
                  },
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25)
                    ),
                    hintStyle: TextStyle(color: Colors.grey),
                    hintText: 'Input Product Name',
                    filled: true,

                  ),
                ),
                SizedBox(height: 20,),
                InkWell(
                  child: Container(
                    width: 300,
                    height: 200,
                    margin: EdgeInsets.all(15),
                    decoration: BoxDecoration(
                      border: Border.all(
                        width: 1,
                      ),
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                  onTap: () {
                    _showPicker(context);
                  },
                ),

                SizedBox(height: 20,),
                Center(
                   child:  Container(
                     width: 250,
                     height: 50,
                     child: ElevatedButton(
                         onPressed: () {
                           addProduct();
                          setState(() {
                            Navigator.pop(context);
                            apiService.getProduct();
                          });

                           // show();
                           clear();
                         },
                          style: ElevatedButton.styleFrom(
                            shape: new RoundedRectangleBorder(
                            borderRadius: new BorderRadius.circular(30.0),
                          ),),
                          child: Text('Add')),
                   )
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
  void addProduct(){
    if(_formKey.currentState!.validate()){
      _formKey.currentState!.save();
      apiService.createProduct(Product(name: _productNameController.text)
      );
    }
  }
  void clear(){
  _productNameController.clear();
  }

  void _showPicker(context) {
    showModalBottomSheet(
        context: context,
        builder: (BuildContext bc) {
          return SafeArea(
            child: Container(
              child: new Wrap(
                children: <Widget>[
                  new ListTile(
                      leading: new Icon(Icons.photo_library),
                      title: new Text('Photo Library'),
                      onTap: () {
                        Navigator.of(context).pop();
                      }),
                  new ListTile(
                    leading: new Icon(Icons.photo_camera),
                    title: new Text('Camera'),
                    onTap: () {
                      Navigator.of(context).pop();
                    },
                  ),
                ],
              ),
            ),
          );
        }
    );
  }
  void show(){
    final snackBar = SnackBar(
      content: const Text('Add Success'),
      backgroundColor: (Colors.black12),
      action: SnackBarAction(
        label: 'dismiss',
        onPressed: () {
        },
      ),
    );
    ScaffoldMessenger.of(context).showSnackBar(snackBar);
  }
}
