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
  final _hostUserController = TextEditingController();
  final _hostPasswordController = TextEditingController();
  final _hostRubixUsernameController = TextEditingController();
  final _hostRubixPasswordController = TextEditingController();
  int? id;
  int? port;
  int? rubixPort;
  @override
  void initState() {
    _hostNameController.text = widget.host.name!;
    _hostIpController.text = widget.host.ip!;
    _hostUserController.text = widget.host.username!;
    _hostPasswordController.text = widget.host.password!;
    _hostRubixUsernameController.text = widget.host.rubixUsername!;
    _hostRubixPasswordController.text = widget.host.rubixPassword!;
    port =  widget.host.port;
    id = widget.host.id!;
    rubixPort = widget.host.rubixPort;
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
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _hostUserController,
                  decoration: InputDecoration(
                    icon: Icon(Icons.font_download),
                    labelText: 'USER *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _hostPasswordController,
                  obscureText: true,
                  decoration: InputDecoration(
                    icon: Icon(Icons.font_download),
                    labelText: 'PASSWORD *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  initialValue: rubixPort.toString(),
                  onSaved: (input) => rubixPort = int.parse(input!),
                  inputFormatters: [FilteringTextInputFormatter.digitsOnly],
                  decoration: InputDecoration(
                    icon: Icon(Icons.pin),
                    labelText: 'RUBIX PORT *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _hostRubixUsernameController,
                  decoration: InputDecoration(
                    icon: Icon(Icons.font_download),
                    labelText: 'RUBIX USERNAME *',
                    hintText: 'What do people call you?',
                    hintStyle: TextStyle(color: Colors.grey),
                    filled: true,
                  ),
                ),
                TextFormField(
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.name,
                  controller: _hostRubixPasswordController,
                  obscureText: true,
                  decoration: InputDecoration(
                    icon: Icon(Icons.font_download),
                    labelText: 'RUBIX PASSWORD *',
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
      apiService.updateHost(id!,Host(name: _hostNameController.text, ip: _hostIpController.text, port: port, username: _hostUserController.text, password: _hostPasswordController.text, rubixPort: rubixPort, rubixUsername: _hostRubixUsernameController.text, rubixPassword: _hostRubixPasswordController.text)
      );
    }
  }
}
