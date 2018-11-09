# Remote runtime configuration using Viper and Consul

* run docker-compose up --build
* insert under key config/app in consul the following YAML config:

```yaml
port: 80
```