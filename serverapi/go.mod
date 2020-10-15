module simpleredis.task/serverapi

go 1.15

replace simpleredis.task/protocol => ../protocol

replace simpleredis.task/coder => ../coder

require (
	simpleredis.task/coder v0.0.0-00010101000000-000000000000
	simpleredis.task/protocol v0.0.0-00010101000000-000000000000
)
