# go-whatsapp

[![GoDoc](https://godoc.org/github.com/rodrigo-brito/go-whatsapp?status.svg)](https://godoc.org/github.com/StudioSol/balancer)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodrigo-brito/go-whatsapp)](https://goreportcard.com/report/github.com/rodrigo-brito/go-whatsapp)

Simple chat with Go, GraphQL and React

![image](https://user-images.githubusercontent.com/7620947/84447162-aa55cc00-ac1d-11ea-85d4-3c3bd45aa654.png)

## Project setup

- **Fron-end**: Front-end based on React + Typescript, mantained in `website` folder.
  - `yarn start` - Development mode
  - `yarn build` - Build production version
  - `yarn type` - Update types from GraphQL schema

- **Back-end**: Back-end based in Go with persistence in Firestore (Googe Firebase)
  - Setup a new firebase project and copy project credentials to `credentials.json`, in project root.
  - `make run` - Start watcher for develpment mode
  - `make gqlgen` - Update and generate a fresh version of GraphQL code

## Licenses

- Released under [MIT License](LICENSE)
- Layout by [Zeno Rocha](https://github.com/zenorocha). Released with MIT license in [Codepen](https://codepen.io/zenorocha/pen/eZxYOK)
