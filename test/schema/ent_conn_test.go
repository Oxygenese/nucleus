package schema_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/enttest"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/migrate"
)

func Test_Ent_Conn(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
}
