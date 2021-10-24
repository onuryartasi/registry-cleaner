
<h3 align="center">Registry Cleaner</h3>

  <p align="center">
    Easy delete manifests and images from your own registry
    <br />
    <a href="https://github.com/onuryartasi/registry-cleaner"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/onuryartasi/registry-cleaner">View Demo</a>
    ·
    <a href="https://github.com/onuryartasi/registry-cleaner/issues">Report Bug</a>
    ·
    <a href="https://github.com/onuryartasi/registry-cleaner/issues">Request Feature</a>
  </p>
</div>

![GitHub](https://img.shields.io/github/license/onuryartasi/registry-cleaner)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/onuryartasi/registry-cleaner)
![GitHub issues](https://img.shields.io/github/issues/onuryartasi/registry-cleaner)



## About The Project

This project's aim is to delete unnecessary images and manifests from your own docker registry. Azure container registry (aka ACR) has purge command, Gitlab registry has a cleanup policy, etc. Registry-cleaner has a configuration for custom rules like GitLab cleanup policy.


## Getting Started
### Prerequisite

 [Go](https://golang.org/doc/install) 

### Installation

1. Get a [release](https://github.com/onuryartasi/registry-cleaner/releases) compatible your system from github.

Or


2. Clone the repo
   ```sh
   git clone https://github.com/onuryartasi/regitry-cleaner.git
   ```
3. Install go modules.
   ```sh
    go mod download
   ```
4. Build source code.
    ```sh
     go build -o registry-cleaner main.go 
    ```
5. Configure your rules `config.yaml`
   ```yaml
    regex:
    enable: true
    pattern:
      - group/image-name
   ```

## Usage

Before usage this tool enable delete function from registry.(https://docs.docker.com/registry/configuration/#delete). You can use env variable `REGISTRY_STORAGE_DELETE_ENABLED=true` for deletable registry from api.


See parameters.
```
  -config-file string
        Config file path. (default "config.yaml")
  -dry-run
        Print deletable images, don't remove.
  -host string
        Registry host (default "localhost")
  -password string
        Registry password
  -username string
        Registry username
```
Before use to `registry-cleaner` configure your own rules for according to your needs.
There is a example config file.
```yaml
#Regex pattern, you can run own regex rules for images.
regex: 
  # If you enable false can't apply this pattern on your images.
  enable: true 
  pattern:
    # Delete images to match those pattern. Multi pattern accaptable.
    - g1/hello-world
    - g2*

# Until-date pattern, when give a date delete all images before this date.
until-date: 
  enable: true
  #### Example date layout
  ###  02.01.2006 15:04:05
  ###  02.01.2006 15:04
  ###  02.01.2006
  date: "31.05.2021"

# N pattern, just latest n images  will be stored. other images will be delete.
n: 
  enable: false
  size: 10 # Number of to image to keeping.
```

```bash
registry-cleaner -host http://localhost:5000 -config-file="config.yaml"

registry-cleaner -host http://localhost:5000 -username admin -password changeme -config-file="config.yaml"
```


After all of them, manually run a garbage collector in your registry if you want to immediately delete images.
```sh
registry garbage-collect /etc/docker/registry/config.yml
```


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazinFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE.md` for more information.


[contributors-shield]: https://img.shields.io/github/contributors/onuryartasi/registry-cleaner.svg?style=for-the-badge
[contributors-url]: https://github.com/onuryartasi/registry-cleaner/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/onuryartasi/registry-cleaner.svg?style=for-the-badge
[forks-url]: https://github.com/onuryartasi/registry-cleaner/network/members
[stars-shield]: https://img.shields.io/github/stars/onuryartasi/registry-cleaner.svg?style=for-the-badge
[stars-url]: https://github.com/onuryartasi/registry-cleaner/stargazers
[issues-shield]: https://img.shields.io/github/issues/onuryartasi/registry-cleaner.svg?style=for-the-badge
[issues-url]: https://github.com/onuryartasi/registry-cleaner/issues
[license-shield]: https://img.shields.io/github/license/onuryartasi/registry-cleaner.svg?style=for-the-badge
[license-url]: https://github.com/onuryartasi/regsitry-cleaner/blob/master/LICENSE.md
[linkedin-url]: https://linkedin.com/in/onuryartasi
