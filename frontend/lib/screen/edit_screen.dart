import 'package:flutter/material.dart';
import 'package:frontend/model/product.dart';
import 'package:frontend/services/api.dart';


class EditScreen extends StatefulWidget {
  const EditScreen(this.product);
  final Product product;
  @override
  _EditScreenState createState() => _EditScreenState();
}

class _EditScreenState extends State<EditScreen> {
  final ApiService apiService = ApiService();
  final _formKey = GlobalKey<FormState>();
  final _productNameController = TextEditingController();
  int? id;
  @override
  void initState() {
    // TODO: implement initState
    _productNameController.text = widget.product.name!;
    id = widget.product.id!;
    super.initState();
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('${widget.product.name}'),
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
                Center(
                    child:  Container(
                      width: 250,
                      height: 50,
                      child: ElevatedButton(
                          onPressed: () {

                              upDateProduct();
                              setState(() {
                                apiService.getProduct();
                              });

                            // show();
                          Navigator.of(context).pop();
                          },
                          style: ElevatedButton.styleFrom(
                            shape: new RoundedRectangleBorder(
                              borderRadius: new BorderRadius.circular(30.0),
                            ),),
                          child: Text('Update')),
                    )
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
  void clear(){
    _productNameController.clear();
  }
  void upDateProduct(){
    if(_formKey.currentState!.validate()){
      _formKey.currentState!.save();
      apiService.updateProduct(id!,Product(name: _productNameController.text)
      );
    }
  }
}
