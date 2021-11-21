import 'package:flutter/material.dart';
import 'package:frontend/model/host.dart';
import 'package:frontend/services/api.dart';

class AddHost extends StatefulWidget {
  const AddHost({Key? key}) : super(key: key);

  @override
  _AddHostState createState() => _AddHostState();
}

class _AddHostState extends State<AddHost> {
  final ApiService apiService = ApiService();
  final _formKey = GlobalKey<FormState>();
  final _hostNameController = TextEditingController();


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Add New Device'),
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
                  controller: _hostNameController,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Please enter host name';
                    }
                    return null;
                  },
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(25)
                    ),
                    hintStyle: TextStyle(color: Colors.grey),
                    hintText: 'Input Host Name',
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
                           addHost();
                          setState(() {
                            Navigator.pop(context);
                            apiService.getHost();
                          });
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
  void addHost(){
    if(_formKey.currentState!.validate()){
      _formKey.currentState!.save();
      apiService.createHost(Host(name: _hostNameController.text)
      );
    }
  }
  void clear(){
  _hostNameController.clear();
  }
}
