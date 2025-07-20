data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./database/loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://${getenv("DB_USER")}:${getenv("DB_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("DB_NAME")}?search_path=public"
  url = "postgres://${getenv("DB_USER")}:${getenv("DB_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("DB_NAME")}?search_path=public"
  migration {
    dir = "file://./database/dev-migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}