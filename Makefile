rebuild:
	(cd packages/shrike; make schema-update)
	(cd packages/wren; make codegen)

start-server:
	(cd packages/shrike; make startdb) &
	(cd packages/shrike; make start) &
	(cd packages/wren; make start)
start-client:
	(cd packages/starling; make hard-start)

reset:
	(cd packages/shrike; make destroy)
	(cd packages/shrike; make initial-migration)