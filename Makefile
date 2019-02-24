codegen:
	(cd shrike; make codgen)
	(cd wren; make codegen)

start:
	(cd packages/shrike; make startdb) &
	(cd packages/shrike; make start) &
	(cd packages/wren; make start)

reset:
	(cd packages/shrike; make destroy)
	(cd packages/shrike; make initial-migration)