class Items {
  final int id;
  final String name;
  final String image;
  final double cost;
  final String describtion;
  bool favorite;
  bool shopcart;
  int count;

  Items({
    required this.id,
    required this.name,
    required this.image,
    required this.cost,
    required this.describtion,
    required this.favorite,
    required this.shopcart,
    required this.count,
  });

  factory Items.fromJson(Map<String, dynamic> json) {
    return Items(
        id: json['ID'],
        name: json['Name'],
        image: json['Image'],
        cost: json['Cost'].toDouble(),
        describtion: json['Describtion'],
        favorite: json['Favorite'],
        shopcart: json['ShopCart'],
        count: json['Count']);
  }
}
