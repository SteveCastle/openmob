codegen:
	(cd shrike; make codgen)
	(cd wren; make codegen)

start:
	(cd packages/shrike; make startdb) &
	(cd packages/shrike; make start) &
	(cd packages/wren; make start)