module server

go 1.15

replace simpleredis.task/api => ../api
replace simpleredis.task/protocol => ../protocol
replace simpleredis.task/serverapi => ../serverapi
replace simpleredis.task/coder => ../coder

require simpleredis.task/serverapi v0.0.0-00010101000000-000000000000
