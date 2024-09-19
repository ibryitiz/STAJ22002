import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_frontend/constant/constant.dart';
import 'package:flutter_frontend/model/product.dart';

class ProductService {
  final Dio dio = Dio();

  ProductService();

  Future<List<Product>> getProducts() async {
    late List<Product> products;

    try {
      final res = await dio.get(productsURL);

      products = res.data["products"]
          .map<Product>(
            (item) => Product.fromJson(item),
          )
          .toList();
    } on DioException catch (e) {
      products = [];
      debugPrint(e.toString());
    }

    return products;
  }
}
