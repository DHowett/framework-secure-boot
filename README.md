## Framework Laptop UEFI Secure Boot Certificates

**Source**: Extracted from a live machine (`FRANBMCP08`)
**Date**: 2021-10-21

### KEK (Key Exchange Key)

This certificate allows Framework to add additional certificates to the DB.

Current thumbprint (from `frame.work-LaptopKEK.der`): `a1c36bd73e143c77954be234f71337370136074f`

### DB (Database)

This certificate allows for the automatic trust of any Framework-signed EFI binaries (such as the
BIOS updater.)

Current thumbprint (from `frame.work-LaptopDB.der`): `732bcb5921f51141a8cd6ff213e4aad43cbb6adc`
