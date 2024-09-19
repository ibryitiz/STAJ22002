import 'package:flutter/material.dart';
import 'package:flutter_frontend/model/product.dart';
import 'package:flutter_frontend/service/api_service.dart';

class HomeScreenViewModel with ChangeNotifier {
  final ProductService _service = ProductService();
  List<Product> _products = [];

  List<Product> get products => _products;

  set products(List<Product> value) {
    _products = value;
    notifyListeners();
  }

  Future<List<Product>> getProductFromService() async {
    try {
      products = await _service.getProducts();
      return _products;
    } catch (e) {
      throw Exception(e.toString());
    }
  }
}
