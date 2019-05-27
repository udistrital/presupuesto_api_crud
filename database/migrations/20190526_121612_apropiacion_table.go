package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ApropiacionTable_20190526_121612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ApropiacionTable_20190526_121612{}
	m.Created = "20190526_121612"

	migration.Register("ApropiacionTable_20190526_121612", m)
}

// Run the migrations
func (m *ApropiacionTable_20190526_121612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE apropiacion ADD CONSTRAINT fk_apropiacion_rubro FOREIGN KEY (rubro) REFERENCES rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE rubro_rubro ADD CONSTRAINT fk_rubro_hijo FOREIGN KEY (rubro_hijo) REFERENCES rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE rubro_rubro ADD CONSTRAINT fk_rubro_padre FOREIGN KEY (rubro_padre) REFERENCES rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE apropiacion ADD CONSTRAINT fk_apropiacion_estado_apropiacion FOREIGN KEY (estado) REFERENCES estado_apropiacion(id) ON UPDATE CASCADE ON DELETE RESTRICT;")
}

// Reverse the migrations
func (m *ApropiacionTable_20190526_121612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
