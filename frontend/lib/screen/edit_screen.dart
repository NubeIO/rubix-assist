import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:frontend/model/host.dart';
import 'package:frontend/services/api.dart';

class EditScreen extends StatefulWidget {
  const EditScreen(this.host);
  final Host host;
  @override
  _EditScreenState createState() => _EditScreenState();
}

class _EditScreenState extends State<EditScreen> {
  final ApiService apiService = ApiService();
  final _formKey = GlobalKey<FormState>();
  final _hostNameController = TextEditingController();
  final _hostIpController = TextEditingController();
  int? id;
  int? port;
  @override
  void initState() {
    _hostNameController.text = widget.host.name!;
    _hostIpController.text = widget.host.ip!;
    port =  widget.host.port;
    id = widget.host.id!;
    super.initState();
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('${widget.host.name}'),
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
                    icon: Icon(Icons.font_download),
                    labelText: 'NAME *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _hostIpController,
                  decoration: InputDecoration(
                    icon: Icon(Icons.font_download),
                    labelText: 'IP *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  initialValue: port.toString(),
                  onSaved: (input) => port = int.parse(input!),
                  inputFormatters: [FilteringTextInputFormatter.digitsOnly],
                  decoration: InputDecoration(
                    icon: Icon(Icons.pin),
                    labelText: 'PORT *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
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
                                apiService.getHost();
                              });
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
    _hostNameController.clear();
  }
  void upDateProduct(){
    if(_formKey.currentState!.validate()){
      _formKey.currentState!.save();
      apiService.updateHost(id!,Host(name: _hostNameController.text, ip: _hostIpController.text, port: port)
      );
    }
  }
}
