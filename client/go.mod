module client

go 1.15

replace simpleredis.task/clientapi => ../clientapi

replace simpleredis.task/protocol => ../protocol
replace simpleredis.task/coder => ../coder

require simpleredis.task/clientapi v0.0.0-00010101000000-000000000000
