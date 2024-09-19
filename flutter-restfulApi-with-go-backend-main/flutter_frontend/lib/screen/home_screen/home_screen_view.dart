import 'package:flutter/material.dart';
import 'package:flutter_frontend/screen/home_screen/home_screen_view_model.dart';
import 'package:provider/provider.dart';

class HomeScreenView extends StatelessWidget {
  const HomeScreenView({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => HomeScreenViewModel()..getProductFromService(),
      builder: (context, child) {
        return Scaffold(
          appBar: _buildAppbar,
          body: const _BuildBody(),
        );
      },
    );
  }

  get _buildAppbar => AppBar(
        backgroundColor: Colors.blueAccent,
        title: const Text(
          "Go ile API istek atma",
          style: TextStyle(color: Colors.white),
        ),
      );
}

class _BuildBody extends StatelessWidget {
  const _BuildBody();

  @override
  Widget build(BuildContext context) {
    return Consumer<HomeScreenViewModel>(
      builder: (context, viewModel, child) {
        if (viewModel.products.isEmpty) {
          return const Center(child: CircularProgressIndicator());
        } else {
          return ListView.builder(
            itemCount: viewModel.products.length,
            itemBuilder: (context, index) {
              var product = viewModel.products[index];
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: ListTile(
                  title: Text(product.name),
                  leading: Text(product.id.toString()),
                  subtitle: Text(product.description),
                  trailing: Text("${product.price.toString()} \$"),
                ),
              );
            },
          );
        }
      },
    );
  }
}
