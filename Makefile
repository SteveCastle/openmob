bootstrap:
	(cd packages/shrike; make bootstrap)
	(npm install -g lerna)
	(lerna bootstrap)
rebuild:
	(cd packages/shrike; make schema-update)
	(cd packages/wren; make codegen)
	(cd packages/starling; make gen; make prettier)
startdb:
	(cd packages/shrike; make startdb)
server:
	(cd packages/shrike; make startdb) &
	(cd packages/shrike; make start) &
	(cd packages/wren; make start)
mock:
	(cd packages/shrike; make load-data)
client:
	(cd packages/starling; make hard-start)

reset:
	(cd packages/shrike; make destroy)
	(cd packages/shrike; make initial-migration)
