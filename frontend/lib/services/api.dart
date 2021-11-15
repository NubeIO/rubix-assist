import 'dart:convert';

import 'package:frontend/model/product.dart';
import 'package:http/http.dart' as http;
import 'package:http/http.dart';


class ApiService {
  final String apiUrl = "http://0.0.0.0:8080/api/hosts";
  Future<List<Product>> getProduct() async {
    final res = await http.get(Uri.parse(apiUrl+"/"));
    if (res.statusCode == 200) {
      List jsRespone = json.decode(res.body);
      print(jsRespone);
      return jsRespone.map((e) => new Product.fromJson(e)).toList();
    } else {
      throw Exception('Failed to load product') ;
    }
  }

  Future<Product> getProductById(String id) async {
    final response = await http.get(Uri.parse('$apiUrl/$id'));

    if (response.statusCode == 200) {
      return Product.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load a case');
    }
  }

  Future<Product> createProduct(Product product) async {
    Map data = {
      'name': product.name
    };

    final  response = await post(
      Uri.parse(apiUrl+"/"),
      headers: <String, String>{
        'Content-Type':'application/json; charset=UTF-8',
        'Accept' : 'application/json'
      },
      body: jsonEncode(data),
    );
    if (response.statusCode == 201) {
      return Product.fromJson(json.decode(response.body));
    } else {
      throw Exception(response.statusCode);
    }
  }

  Future<Product> updateProduct(int id, Product product) async {
    Map data = {
      'name': product.name
    };

    final Response response = await patch(
      Uri.parse('$apiUrl/$id'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(data),
    );
    if (response.statusCode == 200) {
      return Product.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update a case');
    }
  }
  Future  deleteProduct(String id) async {
    final http.Response res = await delete(Uri.parse('$apiUrl/$id'),
      headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },);

    if (res.statusCode == 200) {
      return true;
    } else {
      throw "Failed to delete a case.";
    }
  }
}
