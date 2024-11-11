import 'package:flutter/material.dart';
import 'package:pr10/Pages/EditPage.dart';
import 'package:pr10/api_service.dart';
import 'package:pr10/model/person.dart';

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  late Future<Person> person;

  @override
  void initState() {
    super.initState();
    person = ApiService().getUserByID(1);
  }

  void _refreshData() {
    setState(() {
      person = ApiService().getUserByID(1);
    });
  }

  void navToEdit(BuildContext context) async {
    await Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => EditPage()),
    );
    _refreshData();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.amber[200],
        appBar: AppBar(
          title: const Text('Профиль'),
          backgroundColor: Colors.white70,
        ),
        body: FutureBuilder<Person>(
            future: person,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return CircularProgressIndicator();
              } else if (snapshot.hasError) {
                return Scaffold(
                    backgroundColor: Colors.amber[200],
                    appBar: AppBar(
                      backgroundColor: const Color.fromARGB(255, 255, 246, 218),
                    ),
                    body: Center(child: Text('Error: ${snapshot.error}')));
              } else if (!snapshot.hasData) {
                return Scaffold(
                    backgroundColor: Colors.amber[200],
                    appBar: AppBar(
                      backgroundColor: const Color.fromARGB(255, 255, 246, 218),
                    ),
                    body: Center(child: Text('No product found')));
              }

              final person = snapshot.data!;
              return Column(
                children: [
                  Padding(
                    padding: const EdgeInsets.only(top: 15.0, bottom: 15.0),
                    child: Center(
                      child: Container(
                        width: MediaQuery.of(context).size.width * 0.5,
                        height: MediaQuery.of(context).size.width * 0.5,
                        decoration: BoxDecoration(
                            color: Colors.amber[100],
                            shape: BoxShape.circle,
                            border: Border.all(color: Colors.grey, width: 2),
                            image: DecorationImage(
                              image: NetworkImage(person.image),
                              fit: BoxFit.cover,
                            )),
                      ),
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(10.0),
                    child: Container(
                      decoration: BoxDecoration(
                        color: const Color.fromARGB(255, 255, 246, 218),
                        borderRadius: BorderRadius.circular(15.0),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.all(15.0),
                        child: Column(
                          children: [
                            Padding(
                              padding: const EdgeInsets.only(
                                  top: 15.0, bottom: 35.0),
                              child: Text(
                                person.name,
                                style: const TextStyle(fontSize: 16.0),
                              ),
                            ),
                            Padding(
                                padding: const EdgeInsets.all(8.0),
                                child: Expanded(
                                    child: Align(
                                  alignment: Alignment.centerLeft,
                                  child: Text(
                                    'Телефон: ${person.phone}',
                                    style: const TextStyle(fontSize: 14.0),
                                  ),
                                ))),
                            Padding(
                                padding: const EdgeInsets.all(8.0),
                                child: Expanded(
                                    child: Align(
                                  alignment: Alignment.centerLeft,
                                  child: Text(
                                    'Почта: ${person.mail}',
                                    style: const TextStyle(fontSize: 14.0),
                                  ),
                                ))),
                            Padding(
                                padding: const EdgeInsets.all(8.0),
                                child: Expanded(
                                  child: Align(
                                    alignment: Alignment.centerRight,
                                    child: IconButton(
                                        onPressed: () => navToEdit(context),
                                        icon: const Icon(Icons.edit)),
                                  ),
                                ))
                          ],
                        ),
                      ),
                    ),
                  ),
                ],
              );
            }));
  }
}
