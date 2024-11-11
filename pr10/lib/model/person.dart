class Person {
  final int id;
  String image;
  String name;
  String phone;
  String mail;

  Person(
      {required this.id,
      required this.image,
      required this.name,
      required this.phone,
      required this.mail});

  factory Person.fromJson(Map<String, dynamic> json) {
    return Person(
      id: json['ID'],
      name: json['Name'],
      image: json['Image'],
      phone: json['Phone'],
      mail: json['Mail'],
    );
  }
}
