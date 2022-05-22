const date = new Date();

const helloWorldC = `
    #include <iostream>

    int main()
    {
        std::cout << "Hello, world!\\n";
        return 0;
    }
    `;

const helloWorldJava = `
    class HelloWorldApp {
        public static void main(String[] args) {
            System.out.println("Hello World!"); // Prints the string to the console.
        }
    }
    `;

const helloWorldScala = `
    object HelloWorld extends App {
        println("Hello, World!")
    }
    `;

const helloWorldJS = `
    console.log("Hello World!");
    `;

const helloWorldCsharp = `
    using System;

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello, world!");
        }
    }
    `;

const helloWorldDart = `
    main() {
        print('Hello World!');
    }
    `;

export const feed = [
    {
        TopicID: 0, TopicName: "hello world",
        TopicDescription: "programs that greet the world", SnippetsSinceLastVisit: 23
    },
    {
        TopicID: 1, TopicName: "file I/O",
        TopicDescription: "efficient File ops", SnippetsSinceLastVisit: 12
    },
    {
        TopicID: 2, TopicName: "sorting",
        TopicDescription: "sorting techniques", SnippetsSinceLastVisit: 9
    }
];

export const feedData = () => {
    return feed;
}

export const helloWorldList = [
    { ID: 1, Author: "Michael Jackson", DateCreated: date.toLocaleDateString(), Code: helloWorldC, Language: "C", FileName: "Michale_Says_Hello.c" },
    { ID: 2, Author: "Chuck Norris", DateCreated: date.toLocaleDateString(), Code: helloWorldJava, Language: "java", FileName: "World_greets_Chuck.java" },
    { ID: 3, Author: "Bruce Lee", DateCreated: date.toLocaleDateString(), Code: helloWorldScala, Language: "scala", FileName: "Chuck_greets_Bruce.scala" },
    { ID: 4, Author: "Cher", DateCreated: date.toLocaleDateString(), Code: helloWorldJS, Language: "js", FileName: "Browser_Says_Hello.js" },
    { ID: 5, Author: "Sponge Bob", DateCreated: date.toLocaleDateString(), Code: helloWorldCsharp, Language: "csharp", FileName: "Bill_gates_greets_world.cs" },
    { ID: 6, Author: "John Travolta", DateCreated: date.toLocaleDateString(), Code: helloWorldDart, Language: "dart", FileName: "Travolta_greets_himself.dart" },
]

export const helloWorldData = () => {
    return helloWorldList;
}

export const getTopicNameByID = (id) => {
    return feed[id].TopicName;
}
