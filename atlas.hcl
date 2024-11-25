data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/adult",
    "--dialect", "mysql", // | postgres | sqlite | sqlserver
  ]
}

env "gorm" {
  url = "postgres://userAdmin:supersecret@localhost:5454/postgres_challenge?sslmode=disable"
  src = "postgres://userAdmin:supersecret@localhost:5454/postgres_challenge?sslmode=disable"
  dev = "docker://postgres/15"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
