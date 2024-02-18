table "atlas_schema_revisions" {
  schema = schema.public
  column "version" {
    null = false
    type = character_varying
  }
  column "description" {
    null = false
    type = character_varying
  }
  column "type" {
    null    = false
    type    = bigint
    default = 2
  }
  column "applied" {
    null    = false
    type    = bigint
    default = 0
  }
  column "total" {
    null    = false
    type    = bigint
    default = 0
  }
  column "executed_at" {
    null = false
    type = timestamptz
  }
  column "execution_time" {
    null = false
    type = bigint
  }
  column "error" {
    null = true
    type = text
  }
  column "error_stmt" {
    null = true
    type = text
  }
  column "hash" {
    null = false
    type = character_varying
  }
  column "partial_hashes" {
    null = true
    type = jsonb
  }
  column "operator_version" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.version]
  }
}
table "cities" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = true
    type = text
  }
  column "name_kana" {
    null = true
    type = text
  }
  column "prefecture_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_cities_prefecture" {
    columns     = [column.prefecture_id]
    ref_columns = [table.prefectures.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_cities_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "prefectures" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "code" {
    null = true
    type = text
  }
  column "name" {
    null = true
    type = text
  }
  column "elias" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_prefectures_deleted_at" {
    columns = [column.deleted_at]
  }
}
schema "public" {
  comment = "standard public schema"
}
