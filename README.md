# GoSimpleServer

## Iniciar Handler 
```cgo
handler := &server.Handler{}
```

## Criar Entrada get
```cgo
handler.Get("/get", func(rw *server.ResponseManager) {
    rw.WriteString("hello world")
})
```

## Criar Entrada Post
```cgo
handler.Post("/post", func(rw *server.ResponseManager) {
    rw.WriteString("hello world")
})
```

## Iniciar Servidor Http
```cgo
server.Start("0.0.0.0", 8081, "Servidor rodando em http://{host}:{port}\n")
```