<!--
https://github.com/othneildrew/Best-README-Template
https://gist.github.com/lukas-h/2a5d00690736b4c3a7ba
https://github.com/aleen42/badges
https://github.com/Naereen/badges
-->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h2 align="center">Asastarealty - Backend</h2> <br />
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This system is using for [Asastarealty][company-website] frontend and mobile system. right now we're using Monolith architecture, but it does not rule out the possibility we rewrite using Microservice architecture.

<!-- BUILD WITH -->
#### Built With

This section we explain what is backend is develop with 
* [Golang](https://golang.org)
* [PostgreSQL](https://www.postgresql.org/)
* [Fiber](https://docs.gofiber.io/)
* [Viper](https://github.com/spf13/viper)
* [Reflex](https://github.com/cespare/reflex)
* [Air](https://github.com/cosmtrek/air)
* [Traefik](https://traefik.io/)
* [Prometheus](https://prometheus.io/)
* [Grafana](https://grafana.com/)
* [Fluentd](https://www.fluentd.org/)
* [Elasticsearch](https://www.elastic.co/)
* [Redis](https://redis.io/)
* [NATS](https://nats.io/)
* [cosul](https://www.consul.io/)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

<!-- PREREQUISITES -->
#### Prerequisites

* First you need to install [Golang](https://golang.org/doc/install)

* After that check the installation and Golang version it must be above than 1.11 because we need the [Golang Modules](https://blog.golang.org/using-go-modules)
  ```sh
  > go version
  go version go1.16.3 darwin/amd64
   ```
* Check if your OS have [make](https://www.gnu.org/software/make/), to check it run this command (in my case is using OSX)
  ```
  ❯ make -v
  GNU Make 3.81
  Copyright (C) 2006  Free Software Foundation, Inc.
  This is free software; see the source for copying conditions.
  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
  PARTICULAR PURPOSE.

  This program built for i386-apple-darwin11.3.0
   ```
* Docker is required trought our development process, so highly need it
  ```
  ❯ docker version
  Client:
  Cloud integration: 1.0.14
  Version:           20.10.6
  API version:       1.41
  Go version:        go1.16.3
  Git commit:        370c289
  Built:             Fri Apr  9 22:46:57 2021
  OS/Arch:           darwin/amd64
  Context:           default
  Experimental:      true

  Server: Docker Engine - Community
  Engine:
    Version:          20.10.6
    API version:      1.41 (minimum version 1.12)
    Go version:       go1.13.15
    Git commit:       8728dd2
    Built:            Fri Apr  9 22:44:56 2021
    OS/Arch:          linux/amd64
    Experimental:     false
  containerd:
    Version:          1.4.4
    GitCommit:        05f951a3781f4f2c1911b05e61c160e9c30eaa8e
  runc:
    Version:          1.0.0-rc93
    GitCommit:        12644e614e25b05da6fd08a38ffa0cfe1903fdec
  docker-init:
    Version:          0.19.0
    GitCommit:        de40ad0
   ```
  As you can se the **_Engine_** > **_Version_** > **_20.10.6_** (on my OSX)


<!-- INSTALLATION -->
## Installation

#### Manual Installation
1. Clone the repo
   ```sh
   git clone https://github.com/jadahbakar/asastarealty-backend.git
   ```
2. Install Modules packages
   ```sh
   go mod update
   ```
3. Before to next step, firstable you must manual install [reflex](https://github.com/cespare/reflex)
   ```sh
   go get github.com/cespare/reflex
   ```
4. Check reflex installation
   ```sh
   relex --help
   ```
5. This is Optional if point 4 shows error, you must manualy add to **_SYSTEM PATH_** \n
   in my case i'm using _**zsh**_, so add it to _**.zshrc**_  
   ```sh
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```
   After that refresh the console with
   ```sh
   ❯ . ~/.zshrc
   ```
   Or 
   ```sh
   ❯ source ~/.zshrc
   ```
3. Run it
   ```sh
   make watch
   ```
#### Docker
1. Clone the repo
   ```sh
   git clone https://github.com/jadahbakar/asastarealty-backend.git
   ```
2. Run it (it will create docker images and container with autoreload)
   ```sh
   make devel-up
   ```
3. Stop 
   ```sh
   make devel-down
   ```


<!-- Project structure -->
## Project Structure

```
```
<!-- [Reflex](https://github.com/cespare/reflex)
* [Air](https://github.com/cosmtrek/air) -->

<!-- FEATURES -->
## Features
- [x] Create scaffolding for Monolith - Hexagonal Architechture
- [x] Integrating [Viper](https://github.com/spf13/viper) with go for reading environment variable
  * Define on `appConfig` with `mapstructure` 
- [x] Config REST API with [Fiber](https://docs.gofiber.io/), this include:
  * Middleware 
    * Prefork=**true**
    * Config Logger (/log/fiber.log)
    * Config CORS
    * Config Compress
    * Config CSRF
  * Server
    * Shutdown Normaly
    * Graceful Shutdown
- [x] Configuration for Build Golang Manualy **_go build_**
- [x] mix recepies Golang with Docker
- [x] Using [Traefik](https://traefik.io/) for __Reverse Proxy__
- [x] ~~Create Watcher with **__reflex__** for Development with Hot Reloading~~
- [x] Create Development with Docker and [Air](https://github.com/cosmtrek/air) for Development with Hot Reloading
- [x] Configure [Traefik](https://traefik.io/) desktop
- [ ] Create Monitoring 
  - [ ] Build and Configure [Prometheus](https://prometheus.io/)
  - [ ] Build and Configure [Grafana](https://grafana.com/)
  - [ ] Colaborate  [Traefik](https://traefik.io/) + [Prometheus](https://prometheus.io/) + [Grafana](https://grafana.com/)
  - [ ] Node Health Check
- [ ] Deploy on Local [kubernetes](https://kubernetes.io/id/) Container Orchestration
- [ ] Centralization logger multistaging (Application, Database) with [Fluentd](https://www.fluentd.org/) + [Elasticsearch](https://www.elastic.co/)
- [x] Create Database on PemKot Server (Persistent Storage)
  - [ ] Schema hst (History)
  - [ ] Schema log (Logging)
  - [x] Schema mst (Master)
- [x] Create Makefile
  - [x] Include external environment file (_**app.env**_)
  - [x] All In One service (_**up**_ - _**clean**_)
  - [x] Development (_**build**_ - _**run**_ - _**watch**_)
  - [x] Development with Docker (_**devel-up**_ - _**devel-stop**_ - _**devel-down**_)
  - [x] Build for Docker (_**go-buil**_)
  - [x] Docker (_**docker-build**_ - _**docker-push**_ - _**docker-remove-image**_ - _**docker-remove-container**_ - _**docker-prune**_)
  - [x] Docker-Compose (_**compose-up**_ - _**compose-down**_ - _**compose-clean**_)
- [ ] Deploy and Configure Redis for caching
- [ ] Create Authentication Service
  - [ ] Create Security Service
- [ ] Create Authorization Service
  - [ ] Create Access Control List for User
- [ ] Next Will be Microservice Architechture
- [ ] Next Will Using [NATS](https://nats.io/) for Message Broker (Microservice Architechture)
- [ ] Next Will use Centralizen Configuration (Microservice Architechture)
- [ ] Next Will use Service Registery with Health Check (Microservice Architechture)


<!-- LICENSE -->
## License

Distributed under the Proprietary License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact
Dedy Styawan - [dedy.styawan](https://twitter.com/dedystyawan) - dedy.styawan@gmail.com

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Build Docker Image](https://github.com/target/flottbot/blob/master/docker/Dockerfile)
* [Building APIs with Go — Part 2 -> the newest ](https://fernando-bandeira.medium.com/building-apis-with-go-part-2-4afc50fd0ff8)
* [hexagonal-architecture-in-go -> using it now](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
* [stygis-golang-hexagonal-architecture](https://idevoid.medium.com/stygis-golang-hexagonal-architecture-a2d89d01f84b)
* [how-i-structure-services-in-go](https://medium.com/@ott.kristian/how-i-structure-services-in-go-19147ad0e6bd)
* [Load config from file & environment variables in Golang with Viper](https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d)
* [Go Fiber](https://gofiber.io/)
* [Go Fiber 2](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j)
* [semver for versioning](https://github.com/Masterminds/semver/)
* [string to ascii](https://patorjk.com/software/taag/#p=display&h=2&v=3&f=Graceful&t=BUILD)
* [how-to-build-a-smaller-docker-image](https://medium.com/@gdiener/how-to-build-a-smaller-docker-image-76779e18d48a)
* [go-docker-dev-environment-with-go-modules-and-live-code-reloading](https://threedots.tech/post/go-docker-dev-environment-with-go-modules-and-live-code-reloading/)t
* [self-compiling-go-docker-container](https://www.getwrecked.com/self-compiling-go-docker-container/)
* [Today I Learned: Golang Live-Reload for Development Using Docker Compose + Air](https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076)
* [monitoring-traefik-with-grafana <-- Next Implementation](https://medium.com/platform-engineering/monitoring-traefik-with-grafana-1d037af5b952)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[company-website]:http://asastarealty.com/
[contributors-shield]: https://img.shields.io/github/contributors/jadahbakar/asastarealty-backend.svg?style=for-the-badge
[contributors-url]: https://github.com/jadahbakar/asastarealty-backend/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jadahbakar/asastarealty-backend.svg?style=for-the-badge
[forks-url]: https://github.com/jadahbakar/asastarealty-backend/network/members
[stars-shield]: https://img.shields.io/github/stars/jadahbakar/asastarealty-backend.svg?style=for-the-badge
[stars-url]: https://github.com/jadahbakar/asastarealty-backend/stargazers
[issues-shield]: https://img.shields.io/github/issues/jadahbakar/asastarealty-backend.svg?style=for-the-badge
[issues-url]: https://github.com/jadahbakar/asastarealty-backend/issues
[license-shield]: https://img.shields.io/github/license/jadahbakar/asastarealty-backend.svg?style=for-the-badge
[license-url]: https://github.com/jadahbakar/asastarealty-backend/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/dedystyawan
[product-screenshot]: images/screenshot.png
[logo-screenshot]: images/screenshot.png
[golang-image]: images/golang.png
[golang-install-url]: https://golang.org/doc/install

