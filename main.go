package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/graphql-go/graphql"
    "github.com/graphql-go/handler"
)

// Definir el esquema GraphQL
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
        Name: "RootQuery",
        Fields: graphql.Fields{
            "hello": &graphql.Field{
                Type: graphql.String,
                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                    return "Hola, Mundo desde GraphQL en Go!", nil
                },
            },
        },
    }),
})

func main() {
    // Crear un manejador de GraphQL
    h := handler.New(&handler.Config{
        Schema:     &schema,
        GraphiQL:   true, // Esto habilita la interfaz GraphiQL en el navegador
        Playground: true, // Habilita el Playground
    })

    // Configurar el servidor HTTP
    http.Handle("/graphql", h)
    fmt.Println("Servidor escuchando en http://localhost:8080/graphql")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
