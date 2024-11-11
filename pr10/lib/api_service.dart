import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:pr10/model/items.dart';
import 'package:pr10/model/person.dart';

class ApiService {
  final Dio _dio = Dio();

  final ip = '192.168.1.121';

  Future<List<Items>> getProducts() async {
    try {
      final response = await _dio.get('http://${ip}:8080/products');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<Items> getProductsByID(int index) async {
    final link = 'http://${ip}:8080/products/${index.toString()}';
    try {
      final response = await _dio.get(link);
      if (response.statusCode == 200) {
        Items product = Items.fromJson(response.data);
        print(product);
        return product;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> addProduct(Items item) async {
    final link = 'http://${ip}:8080/products';
    try {
      final response = await _dio.post(link, data: {
        'ID': 0,
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> updateProductStatus(Items item) async {
    final link = 'http://${ip}:8080/products/${item.id}';
    try {
      final response = await _dio.put(link, data: {
        'ID': item.id,
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> deleteProduct(int index) async {
    final link = 'http://${ip}:8080/products/${index}';
    try {
      final response = await _dio.delete(link);
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<List<Items>> getShopCartProducts(user_id) async {
    try {
      final response = await _dio.get('http://${ip}:8080/cart/$user_id');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> addProductShopCart(Items item, user_id) async {
    final link = 'http://${ip}:8080/cart/$user_id';
    try {
      final response = await _dio.post(link, data: {
        'ID': item.id,
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> updateProductShopCart(Items item, user_id) async {
    final link = 'http://${ip}:8080/cart/$user_id';
    try {
      final response = await _dio.put(link, data: {
        'ID': item.id,
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> deleteProductShopCart(int user_id, int product_id) async {
    final link = 'http://${ip}:8080/cart/$user_id/$product_id';
    try {
      final response = await _dio.delete(link);
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<List<Items>> getFavoriteProducts(user_id) async {
    try {
      final response = await _dio.get('http://${ip}:8080/favorites/$user_id');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> addProductFavorite(Items item, user_id) async {
    final link = 'http://${ip}:8080/favorites/$user_id';
    try {
      final response = await _dio.post(link, data: {
        'ID': item.id,
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> deleteProductFavorite(int user_id, int product_id) async {
    final link = 'http://${ip}:8080/favorites/$user_id/$product_id';
    try {
      final response = await _dio.delete(link);
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<Person> getUserByID(int index) async {
    final link = 'http://${ip}:8080/users/$index';
    try {
      final response = await _dio.get(link);
      if (response.statusCode == 200) {
        Person product = Person.fromJson(response.data);
        return product;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> updateUser(Person item) async {
    final link = 'http://${ip}:8080/users/$item.id';
    try {
      final response = await _dio.put(link, data: {
        'ID': item.id,
        'Name': item.name,
        'Image': item.image,
        'Phone': item.phone,
        'Mail': item.mail,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }
}
